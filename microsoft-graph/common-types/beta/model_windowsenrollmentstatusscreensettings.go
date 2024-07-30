package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsEnrollmentStatusScreenSettings struct {
	AllowDeviceUseBeforeProfileAndAppInstallComplete *bool   `json:"allowDeviceUseBeforeProfileAndAppInstallComplete,omitempty"`
	AllowDeviceUseOnInstallFailure                   *bool   `json:"allowDeviceUseOnInstallFailure,omitempty"`
	AllowLogCollectionOnInstallFailure               *bool   `json:"allowLogCollectionOnInstallFailure,omitempty"`
	BlockDeviceSetupRetryByUser                      *bool   `json:"blockDeviceSetupRetryByUser,omitempty"`
	CustomErrorMessage                               *string `json:"customErrorMessage,omitempty"`
	HideInstallationProgress                         *bool   `json:"hideInstallationProgress,omitempty"`
	InstallProgressTimeoutInMinutes                  *int64  `json:"installProgressTimeoutInMinutes,omitempty"`
	ODataType                                        *string `json:"@odata.type,omitempty"`
}
