package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAttributeValuesItem struct {
	IsDefault *bool   `json:"isDefault,omitempty"`
	Name      *string `json:"name,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Value     *string `json:"value,omitempty"`
}
