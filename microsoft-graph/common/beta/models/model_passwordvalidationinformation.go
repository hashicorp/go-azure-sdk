package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PasswordValidationInformation struct {
	IsValid           *bool               `json:"isValid,omitempty"`
	ODataType         *string             `json:"@odata.type,omitempty"`
	ValidationResults *[]ValidationResult `json:"validationResults,omitempty"`
}
