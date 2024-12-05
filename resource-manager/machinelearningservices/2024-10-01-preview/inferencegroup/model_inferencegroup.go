package inferencegroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferenceGroup struct {
	Description              *string                        `json:"description,omitempty"`
	EnvironmentConfiguration *GroupEnvironmentConfiguration `json:"environmentConfiguration,omitempty"`
	ModelConfiguration       *GroupModelConfiguration       `json:"modelConfiguration,omitempty"`
	NodeSkuType              *string                        `json:"nodeSkuType,omitempty"`
	Properties               *[]StringStringKeyValuePair    `json:"properties,omitempty"`
	ProvisioningState        *PoolProvisioningState         `json:"provisioningState,omitempty"`
	ScaleUnitSize            *int64                         `json:"scaleUnitSize,omitempty"`
}
