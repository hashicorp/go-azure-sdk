package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationServiceProperties struct {
	Instances         *[]ConfigurationServiceInstance        `json:"instances,omitempty"`
	ProvisioningState *ConfigurationServiceProvisioningState `json:"provisioningState,omitempty"`
	ResourceRequests  *ConfigurationServiceResourceRequests  `json:"resourceRequests"`
	Settings          *ConfigurationServiceSettings          `json:"settings"`
}
