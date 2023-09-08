package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordSingleSignOnField struct {
	CustomizedLabel *string `json:"customizedLabel,omitempty"`
	DefaultLabel    *string `json:"defaultLabel,omitempty"`
	FieldId         *string `json:"fieldId,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	Type            *string `json:"type,omitempty"`
}
