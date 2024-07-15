package dotnetcomponents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DotNetComponentProperties struct {
	ComponentType     *DotNetComponentType                    `json:"componentType,omitempty"`
	Configurations    *[]DotNetComponentConfigurationProperty `json:"configurations,omitempty"`
	ProvisioningState *DotNetComponentProvisioningState       `json:"provisioningState,omitempty"`
	ServiceBinds      *[]DotNetComponentServiceBind           `json:"serviceBinds,omitempty"`
}
