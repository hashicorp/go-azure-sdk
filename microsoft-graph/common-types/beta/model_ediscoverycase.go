package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCase struct {
	ClosedBy                *IdentitySet                        `json:"closedBy,omitempty"`
	ClosedDateTime          *string                             `json:"closedDateTime,omitempty"`
	CreatedDateTime         *string                             `json:"createdDateTime,omitempty"`
	Custodians              *[]EdiscoveryCustodian              `json:"custodians,omitempty"`
	Description             *string                             `json:"description,omitempty"`
	DisplayName             *string                             `json:"displayName,omitempty"`
	ExternalId              *string                             `json:"externalId,omitempty"`
	Id                      *string                             `json:"id,omitempty"`
	LastModifiedBy          *IdentitySet                        `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    *string                             `json:"lastModifiedDateTime,omitempty"`
	LegalHolds              *[]EdiscoveryLegalHold              `json:"legalHolds,omitempty"`
	NoncustodialDataSources *[]EdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`
	ODataType               *string                             `json:"@odata.type,omitempty"`
	Operations              *[]EdiscoveryCaseOperation          `json:"operations,omitempty"`
	ReviewSets              *[]EdiscoveryReviewSet              `json:"reviewSets,omitempty"`
	Settings                *EdiscoveryCaseSettings             `json:"settings,omitempty"`
	SourceCollections       *[]EdiscoverySourceCollection       `json:"sourceCollections,omitempty"`
	Status                  *EdiscoveryCaseStatus               `json:"status,omitempty"`
	Tags                    *[]EdiscoveryTag                    `json:"tags,omitempty"`
}
