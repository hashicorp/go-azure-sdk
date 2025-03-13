package appattachpackageinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppAttachPackageProperties struct {
	FailHealthCheckOnStagingFailure *FailHealthCheckOnStagingFailure `json:"failHealthCheckOnStagingFailure,omitempty"`
	HostPoolReferences              *[]string                        `json:"hostPoolReferences,omitempty"`
	Image                           *AppAttachPackageInfoProperties  `json:"image,omitempty"`
	KeyVaultURL                     *string                          `json:"keyVaultURL,omitempty"`
	ProvisioningState               *ProvisioningState               `json:"provisioningState,omitempty"`
}
