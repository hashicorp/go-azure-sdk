package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryCase struct {
	ClosedBy                *IdentitySet                                `json:"closedBy,omitempty"`
	ClosedDateTime          *string                                     `json:"closedDateTime,omitempty"`
	CreatedDateTime         *string                                     `json:"createdDateTime,omitempty"`
	Custodians              *[]SecurityEdiscoveryCustodian              `json:"custodians,omitempty"`
	Description             *string                                     `json:"description,omitempty"`
	DisplayName             *string                                     `json:"displayName,omitempty"`
	ExternalId              *string                                     `json:"externalId,omitempty"`
	Id                      *string                                     `json:"id,omitempty"`
	LastModifiedBy          *IdentitySet                                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    *string                                     `json:"lastModifiedDateTime,omitempty"`
	NoncustodialDataSources *[]SecurityEdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`
	ODataType               *string                                     `json:"@odata.type,omitempty"`
	Operations              *[]SecurityCaseOperation                    `json:"operations,omitempty"`
	ReviewSets              *[]SecurityEdiscoveryReviewSet              `json:"reviewSets,omitempty"`
	Searches                *[]SecurityEdiscoverySearch                 `json:"searches,omitempty"`
	Settings                *SecurityEdiscoveryCaseSettings             `json:"settings,omitempty"`
	Status                  *SecurityEdiscoveryCaseStatus               `json:"status,omitempty"`
	Tags                    *[]SecurityEdiscoveryReviewTag              `json:"tags,omitempty"`
}
