package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRequest struct {
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`
	CatalogId             *string                `json:"catalogId,omitempty"`
	ExecuteImmediately    *bool                  `json:"executeImmediately,omitempty"`
	ExpirationDateTime    *string                `json:"expirationDateTime,omitempty"`
	Id                    *string                `json:"id,omitempty"`
	IsValidationOnly      *bool                  `json:"isValidationOnly,omitempty"`
	Justification         *string                `json:"justification,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	RequestState          *string                `json:"requestState,omitempty"`
	RequestStatus         *string                `json:"requestStatus,omitempty"`
	RequestType           *string                `json:"requestType,omitempty"`
	Requestor             *AccessPackageSubject  `json:"requestor,omitempty"`
}
