package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionProperty struct {
	AppDisplayName         *string   `json:"appDisplayName,omitempty"`
	DataType               *string   `json:"dataType,omitempty"`
	DeletedDateTime        *string   `json:"deletedDateTime,omitempty"`
	Id                     *string   `json:"id,omitempty"`
	IsMultiValued          *bool     `json:"isMultiValued,omitempty"`
	IsSyncedFromOnPremises *bool     `json:"isSyncedFromOnPremises,omitempty"`
	Name                   *string   `json:"name,omitempty"`
	ODataType              *string   `json:"@odata.type,omitempty"`
	TargetObjects          *[]string `json:"targetObjects,omitempty"`
}
