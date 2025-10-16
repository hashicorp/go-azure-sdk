package serverlessendpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessEndpoint struct {
	AuthMode                  ServerlessInferenceEndpointAuthMode `json:"authMode"`
	ContentSafety             *ContentSafety                      `json:"contentSafety,omitempty"`
	EndpointState             *ServerlessEndpointState            `json:"endpointState,omitempty"`
	InferenceEndpoint         *ServerlessInferenceEndpoint        `json:"inferenceEndpoint,omitempty"`
	MarketplaceSubscriptionId *string                             `json:"marketplaceSubscriptionId,omitempty"`
	ModelSettings             *ModelSettings                      `json:"modelSettings,omitempty"`
	ProvisioningState         *EndpointProvisioningState          `json:"provisioningState,omitempty"`
}
