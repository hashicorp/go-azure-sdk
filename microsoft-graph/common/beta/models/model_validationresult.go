package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationResult struct {
	Message          *string `json:"message,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	RuleName         *string `json:"ruleName,omitempty"`
	ValidationPassed *bool   `json:"validationPassed,omitempty"`
}
