package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LicenseUnitsDetail struct {
	Enabled   *int64  `json:"enabled,omitempty"`
	LockedOut *int64  `json:"lockedOut,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Suspended *int64  `json:"suspended,omitempty"`
	Warning   *int64  `json:"warning,omitempty"`
}
