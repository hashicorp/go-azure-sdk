package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearch struct {
	AddToReviewSetOperation         *SecurityEdiscoveryAddToReviewSetOperation  `json:"addToReviewSetOperation,omitempty"`
	AdditionalSources               *[]SecurityDataSource                       `json:"additionalSources,omitempty"`
	ContentQuery                    *string                                     `json:"contentQuery,omitempty"`
	CreatedBy                       *IdentitySet                                `json:"createdBy,omitempty"`
	CreatedDateTime                 *string                                     `json:"createdDateTime,omitempty"`
	CustodianSources                *[]SecurityDataSource                       `json:"custodianSources,omitempty"`
	DataSourceScopes                *SecurityEdiscoverySearchDataSourceScopes   `json:"dataSourceScopes,omitempty"`
	Description                     *string                                     `json:"description,omitempty"`
	DisplayName                     *string                                     `json:"displayName,omitempty"`
	Id                              *string                                     `json:"id,omitempty"`
	LastEstimateStatisticsOperation *SecurityEdiscoveryEstimateOperation        `json:"lastEstimateStatisticsOperation,omitempty"`
	LastModifiedBy                  *IdentitySet                                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime            *string                                     `json:"lastModifiedDateTime,omitempty"`
	NoncustodialSources             *[]SecurityEdiscoveryNoncustodialDataSource `json:"noncustodialSources,omitempty"`
	ODataType                       *string                                     `json:"@odata.type,omitempty"`
}
