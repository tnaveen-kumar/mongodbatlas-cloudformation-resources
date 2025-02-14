// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package profile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	DefaultProfile    = "default"
	profileNamePrefix = "cfn/atlas/profile"
)

type Profile struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	BaseURL    string `json:"BaseUrl,omitempty"`
}

func NewProfile(req *handler.Request, profileName *string) (*Profile, error) {
	if profileName == nil || *profileName == "" {
		profileName = aws.String(DefaultProfile)
	}

	secretsManagerClient := secretsmanager.New(req.Session)
	resp, err := secretsManagerClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(fmt.Sprintf("%s/%s", profileNamePrefix, *profileName))})
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	err = json.Unmarshal([]byte(*resp.SecretString), &profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *Profile) NewBaseURL() string {
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		return baseURL
	}

	return p.BaseURL
}

func (p *Profile) NewPublicKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY"); k != "" {
		return k
	}

	return p.PublicKey
}

func (p *Profile) NewPrivateKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY"); k != "" {
		return k
	}

	return p.PrivateKey
}

func (p *Profile) AreKeysAvailable() bool {
	return p.NewPublicKey() == "" || p.PrivateKey == ""
}
