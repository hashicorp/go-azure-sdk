package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessRoot struct {
	Group     *PrivilegedAccessGroup `json:"group,omitempty"`
	Id        *string                `json:"id,omitempty"`
	ODataType *string                `json:"@odata.type,omitempty"`
}
