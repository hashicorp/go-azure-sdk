package chromeosonboardingsetting

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateChromeOSOnboardingSettingConnectRequest struct {
	OwnerAccessToken       nullable.Type[string] `json:"ownerAccessToken,omitempty"`
	OwnerUserPrincipalName nullable.Type[string] `json:"ownerUserPrincipalName,omitempty"`
}
