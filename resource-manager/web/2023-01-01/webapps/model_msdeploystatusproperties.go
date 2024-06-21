package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployStatusProperties struct {
	Complete          *bool                      `json:"complete,omitempty"`
	Deployer          *string                    `json:"deployer,omitempty"`
	EndTime           *string                    `json:"endTime,omitempty"`
	ProvisioningState *MSDeployProvisioningState `json:"provisioningState,omitempty"`
	StartTime         *string                    `json:"startTime,omitempty"`
}
