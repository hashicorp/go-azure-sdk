package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnitOfMeasure struct {
	Code                      *string `json:"code,omitempty"`
	DisplayName               *string `json:"displayName,omitempty"`
	Id                        *string `json:"id,omitempty"`
	InternationalStandardCode *string `json:"internationalStandardCode,omitempty"`
	LastModifiedDateTime      *string `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
}
