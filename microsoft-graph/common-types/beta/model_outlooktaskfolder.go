package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskFolder struct {
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	IsDefaultFolder               *bool                                `json:"isDefaultFolder,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	Name                          *string                              `json:"name,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	ParentGroupKey                *string                              `json:"parentGroupKey,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	Tasks                         *[]OutlookTask                       `json:"tasks,omitempty"`
}
