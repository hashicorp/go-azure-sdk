package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswer struct {
	AnsweredQuestion *AccessPackageQuestion `json:"answeredQuestion,omitempty"`
	DisplayValue     *string                `json:"displayValue,omitempty"`
	ODataType        *string                `json:"@odata.type,omitempty"`
}
