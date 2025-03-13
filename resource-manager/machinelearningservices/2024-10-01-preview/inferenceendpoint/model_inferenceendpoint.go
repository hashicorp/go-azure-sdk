package inferenceendpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceEndpoint struct {
	AuthMode             AuthMode                    `json:"authMode"`
	Description          *string                     `json:"description,omitempty"`
	EndpointUri          *string                     `json:"endpointUri,omitempty"`
	GroupName            string                      `json:"groupName"`
	Properties           *[]StringStringKeyValuePair `json:"properties,omitempty"`
	ProvisioningState    *PoolProvisioningState      `json:"provisioningState,omitempty"`
	RequestConfiguration *RequestConfiguration       `json:"requestConfiguration,omitempty"`
}
