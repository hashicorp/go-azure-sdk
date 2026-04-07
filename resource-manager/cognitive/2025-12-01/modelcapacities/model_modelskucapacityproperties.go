package modelcapacities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelSkuCapacityProperties struct {
	AvailableCapacity         *float64         `json:"availableCapacity,omitempty"`
	AvailableFinetuneCapacity *float64         `json:"availableFinetuneCapacity,omitempty"`
	Model                     *DeploymentModel `json:"model,omitempty"`
	SkuName                   *string          `json:"skuName,omitempty"`
}
