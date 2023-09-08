package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionPageViewConfiguration struct {
	Description *string                                                `json:"description,omitempty"`
	Inputs      *[]AuthenticationAttributeCollectionInputConfiguration `json:"inputs,omitempty"`
	ODataType   *string                                                `json:"@odata.type,omitempty"`
	Title       *string                                                `json:"title,omitempty"`
}
