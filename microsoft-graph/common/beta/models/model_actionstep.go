package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionStep struct {
	ActionUrl  *ActionUrl `json:"actionUrl,omitempty"`
	ODataType  *string    `json:"@odata.type,omitempty"`
	StepNumber *int64     `json:"stepNumber,omitempty"`
	Text       *string    `json:"text,omitempty"`
}
