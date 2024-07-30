package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderApplicationControlSupplementalPolicyDeploymentSummary struct {
	DeployedDeviceCount *int64  `json:"deployedDeviceCount,omitempty"`
	FailedDeviceCount   *int64  `json:"failedDeviceCount,omitempty"`
	Id                  *string `json:"id,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
}
