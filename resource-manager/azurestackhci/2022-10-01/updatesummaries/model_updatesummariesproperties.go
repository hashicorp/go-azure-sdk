package updatesummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateSummariesProperties struct {
	CurrentVersion    *string                         `json:"currentVersion,omitempty"`
	HardwareModel     *string                         `json:"hardwareModel,omitempty"`
	OemFamily         *string                         `json:"oemFamily,omitempty"`
	PackageVersions   *[]PackageVersionInfo           `json:"packageVersions,omitempty"`
	ProvisioningState *ProvisioningState              `json:"provisioningState,omitempty"`
	State             *UpdateSummariesPropertiesState `json:"state,omitempty"`
}
