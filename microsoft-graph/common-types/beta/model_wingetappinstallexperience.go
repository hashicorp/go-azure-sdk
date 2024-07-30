package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WinGetAppInstallExperience struct {
	ODataType    *string                                 `json:"@odata.type,omitempty"`
	RunAsAccount *WinGetAppInstallExperienceRunAsAccount `json:"runAsAccount,omitempty"`
}
