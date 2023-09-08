package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChoiceColumn struct {
	AllowTextEntry *bool     `json:"allowTextEntry,omitempty"`
	Choices        *[]string `json:"choices,omitempty"`
	DisplayAs      *string   `json:"displayAs,omitempty"`
	ODataType      *string   `json:"@odata.type,omitempty"`
}
