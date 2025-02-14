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

// Code generated by 'cfn generate', changes will be undone by the next invocation. DO NOT EDIT.
// Updates to this type are made my editing the schema file and executing the 'generate' command.
package resource

// Model is autogenerated from the json schema
type Model struct {
	ConnectionStrings            *ServerlessInstanceConnectionStrings `json:",omitempty"`
	ContinuousBackupEnabled      *bool                                `json:",omitempty"`
	CreateDate                   *string                              `json:",omitempty"`
	Id                           *string                              `json:",omitempty"`
	IncludeCount                 *bool                                `json:",omitempty"`
	ItemsPerPage                 *int                                 `json:",omitempty"`
	Links                        []Link                               `json:",omitempty"`
	MongoDBVersion               *string                              `json:",omitempty"`
	Name                         *string                              `json:",omitempty"`
	PageNum                      *int                                 `json:",omitempty"`
	ProjectID                    *string                              `json:",omitempty"`
	ProviderSettings             *ServerlessInstanceProviderSettings  `json:",omitempty"`
	StateName                    *string                              `json:",omitempty"`
	TerminationProtectionEnabled *bool                                `json:",omitempty"`
	TotalCount                   *float64                             `json:",omitempty"`
	Profile                      *string                              `json:",omitempty"`
}

// ServerlessInstanceConnectionStrings is autogenerated from the json schema
type ServerlessInstanceConnectionStrings struct {
	PrivateEndpoint []ServerlessInstancePrivateEndpoint `json:",omitempty"`
	StandardSrv     *string                             `json:",omitempty"`
}

// ServerlessInstancePrivateEndpoint is autogenerated from the json schema
type ServerlessInstancePrivateEndpoint struct {
	Endpoints           []ServerlessInstancePrivateEndpointEndpoint `json:",omitempty"`
	SrvConnectionString *string                                     `json:",omitempty"`
	Type                *string                                     `json:",omitempty"`
}

// ServerlessInstancePrivateEndpointEndpoint is autogenerated from the json schema
type ServerlessInstancePrivateEndpointEndpoint struct {
	EndpointId   *string `json:",omitempty"`
	ProviderName *string `json:",omitempty"`
	Region       *string `json:",omitempty"`
}

// Link is autogenerated from the json schema
type Link struct {
	Href *string `json:",omitempty"`
	Rel  *string `json:",omitempty"`
}

// ServerlessInstanceProviderSettings is autogenerated from the json schema
type ServerlessInstanceProviderSettings struct {
	ProviderName *string `json:",omitempty"`
	RegionName   *string `json:",omitempty"`
}
