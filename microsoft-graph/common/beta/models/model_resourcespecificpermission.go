package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSpecificPermission struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Id          *string `json:"id,omitempty"`
	IsEnabled   *bool   `json:"isEnabled,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Value       *string `json:"value,omitempty"`
}
