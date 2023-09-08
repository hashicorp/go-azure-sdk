package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswerChoice struct {
	ActualValue  *string                        `json:"actualValue,omitempty"`
	DisplayValue *AccessPackageLocalizedContent `json:"displayValue,omitempty"`
	ODataType    *string                        `json:"@odata.type,omitempty"`
}
