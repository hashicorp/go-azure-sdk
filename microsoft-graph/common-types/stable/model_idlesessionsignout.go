package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdleSessionSignOut struct {
	IsEnabled             *bool   `json:"isEnabled,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	SignOutAfterInSeconds *int64  `json:"signOutAfterInSeconds,omitempty"`
	WarnAfterInSeconds    *int64  `json:"warnAfterInSeconds,omitempty"`
}
