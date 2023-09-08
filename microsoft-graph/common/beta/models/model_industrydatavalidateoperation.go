package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataValidateOperation struct {
	CreatedDateTime    *string                              `json:"createdDateTime,omitempty"`
	Errors             *[]PublicError                       `json:"errors,omitempty"`
	Id                 *string                              `json:"id,omitempty"`
	LastActionDateTime *string                              `json:"lastActionDateTime,omitempty"`
	ODataType          *string                              `json:"@odata.type,omitempty"`
	ResourceLocation   *string                              `json:"resourceLocation,omitempty"`
	Status             *IndustryDataValidateOperationStatus `json:"status,omitempty"`
	StatusDetail       *string                              `json:"statusDetail,omitempty"`
	Warnings           *[]PublicError                       `json:"warnings,omitempty"`
}
