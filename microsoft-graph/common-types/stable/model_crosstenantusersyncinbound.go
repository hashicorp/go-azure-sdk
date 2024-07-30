package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantUserSyncInbound struct {
	IsSyncAllowed *bool   `json:"isSyncAllowed,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
