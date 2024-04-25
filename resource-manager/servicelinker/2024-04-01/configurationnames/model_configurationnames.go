package configurationnames

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationNames struct {
	AuthType       *AuthType            `json:"authType,omitempty"`
	ClientType     *ClientType          `json:"clientType,omitempty"`
	DaprProperties *DaprProperties      `json:"daprProperties,omitempty"`
	Names          *[]ConfigurationName `json:"names,omitempty"`
	SecretType     *SecretSourceType    `json:"secretType,omitempty"`
	TargetService  *string              `json:"targetService,omitempty"`
}
