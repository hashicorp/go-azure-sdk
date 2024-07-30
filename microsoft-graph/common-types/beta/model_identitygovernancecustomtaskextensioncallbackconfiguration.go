package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceCustomTaskExtensionCallbackConfiguration struct {
	AuthorizedApps  *[]Application `json:"authorizedApps,omitempty"`
	ODataType       *string        `json:"@odata.type,omitempty"`
	TimeoutDuration *string        `json:"timeoutDuration,omitempty"`
}
