package serverlessruntimes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformaticaServerlessRuntimeProperties struct {
	AdvancedCustomProperties               *[]AdvancedCustomProperties             `json:"advancedCustomProperties,omitempty"`
	ApplicationType                        *ApplicationType                        `json:"applicationType,omitempty"`
	ComputeUnits                           *string                                 `json:"computeUnits,omitempty"`
	Description                            *string                                 `json:"description,omitempty"`
	ExecutionTimeout                       *string                                 `json:"executionTimeout,omitempty"`
	Platform                               *PlatformType                           `json:"platform,omitempty"`
	ProvisioningState                      *ProvisioningState                      `json:"provisioningState,omitempty"`
	ServerlessAccountLocation              string                                  `json:"serverlessAccountLocation"`
	ServerlessRuntimeConfig                *ServerlessRuntimeConfigProperties      `json:"serverlessRuntimeConfig,omitempty"`
	ServerlessRuntimeNetworkProfile        *ServerlessRuntimeNetworkProfile        `json:"serverlessRuntimeNetworkProfile,omitempty"`
	ServerlessRuntimeTags                  *[]ServerlessRuntimeTag                 `json:"serverlessRuntimeTags,omitempty"`
	ServerlessRuntimeUserContextProperties *ServerlessRuntimeUserContextProperties `json:"serverlessRuntimeUserContextProperties,omitempty"`
	SupplementaryFileLocation              *string                                 `json:"supplementaryFileLocation,omitempty"`
}
