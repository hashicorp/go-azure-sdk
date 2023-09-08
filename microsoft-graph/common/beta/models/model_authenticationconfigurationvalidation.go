package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationConfigurationValidation struct {
	Errors    *[]GenericError `json:"errors,omitempty"`
	ODataType *string         `json:"@odata.type,omitempty"`
	Warnings  *[]GenericError `json:"warnings,omitempty"`
}
