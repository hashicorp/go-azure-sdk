package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionGrantPolicy struct {
	DeletedDateTime *string                        `json:"deletedDateTime,omitempty"`
	Description     *string                        `json:"description,omitempty"`
	DisplayName     *string                        `json:"displayName,omitempty"`
	Excludes        *[]PermissionGrantConditionSet `json:"excludes,omitempty"`
	Id              *string                        `json:"id,omitempty"`
	Includes        *[]PermissionGrantConditionSet `json:"includes,omitempty"`
	ODataType       *string                        `json:"@odata.type,omitempty"`
}
