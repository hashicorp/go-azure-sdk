package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SchemaExtension struct {
	Description *string                    `json:"description,omitempty"`
	Id          *string                    `json:"id,omitempty"`
	ODataType   *string                    `json:"@odata.type,omitempty"`
	Owner       *string                    `json:"owner,omitempty"`
	Properties  *[]ExtensionSchemaProperty `json:"properties,omitempty"`
	Status      *string                    `json:"status,omitempty"`
	TargetTypes *[]string                  `json:"targetTypes,omitempty"`
}
