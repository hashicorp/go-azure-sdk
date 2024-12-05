package inferencepool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferencePool struct {
	Description            *string                     `json:"description,omitempty"`
	Properties             *[]StringStringKeyValuePair `json:"properties,omitempty"`
	ProvisioningState      *PoolProvisioningState      `json:"provisioningState,omitempty"`
	ScaleUnitConfiguration *ScaleUnitConfiguration     `json:"scaleUnitConfiguration,omitempty"`
}
