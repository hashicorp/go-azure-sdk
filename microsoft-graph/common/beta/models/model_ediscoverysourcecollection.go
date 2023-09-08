package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySourceCollection struct {
	AddToReviewSetOperation         *EdiscoveryAddToReviewSetOperation          `json:"addToReviewSetOperation,omitempty"`
	AdditionalSources               *[]EdiscoveryDataSource                     `json:"additionalSources,omitempty"`
	ContentQuery                    *string                                     `json:"contentQuery,omitempty"`
	CreatedBy                       *IdentitySet                                `json:"createdBy,omitempty"`
	CreatedDateTime                 *string                                     `json:"createdDateTime,omitempty"`
	CustodianSources                *[]EdiscoveryDataSource                     `json:"custodianSources,omitempty"`
	DataSourceScopes                *EdiscoverySourceCollectionDataSourceScopes `json:"dataSourceScopes,omitempty"`
	Description                     *string                                     `json:"description,omitempty"`
	DisplayName                     *string                                     `json:"displayName,omitempty"`
	Id                              *string                                     `json:"id,omitempty"`
	LastEstimateStatisticsOperation *EdiscoveryEstimateStatisticsOperation      `json:"lastEstimateStatisticsOperation,omitempty"`
	LastModifiedBy                  *IdentitySet                                `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime            *string                                     `json:"lastModifiedDateTime,omitempty"`
	NoncustodialSources             *[]EdiscoveryNoncustodialDataSource         `json:"noncustodialSources,omitempty"`
	ODataType                       *string                                     `json:"@odata.type,omitempty"`
}
