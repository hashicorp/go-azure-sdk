package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesFilter struct {
	Mode      *DevicesFilterMode `json:"mode,omitempty"`
	ODataType *string            `json:"@odata.type,omitempty"`
	Rule      *string            `json:"rule,omitempty"`
}
