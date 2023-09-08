package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAutonomousSystem struct {
	Name         *string `json:"name,omitempty"`
	Number       *int64  `json:"number,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Value        *string `json:"value,omitempty"`
}
