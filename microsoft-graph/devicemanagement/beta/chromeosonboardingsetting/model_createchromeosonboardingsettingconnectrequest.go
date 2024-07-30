package chromeosonboardingsetting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateChromeOSOnboardingSettingConnectRequest struct {
	OwnerAccessToken       *string `json:"ownerAccessToken,omitempty"`
	OwnerUserPrincipalName *string `json:"ownerUserPrincipalName,omitempty"`
}
