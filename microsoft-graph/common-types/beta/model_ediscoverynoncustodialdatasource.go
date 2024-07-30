package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryNoncustodialDataSource struct {
	ApplyHoldToSource    *bool                                       `json:"applyHoldToSource,omitempty"`
	CreatedDateTime      *string                                     `json:"createdDateTime,omitempty"`
	DataSource           *EdiscoveryDataSource                       `json:"dataSource,omitempty"`
	DisplayName          *string                                     `json:"displayName,omitempty"`
	HoldStatus           *EdiscoveryNoncustodialDataSourceHoldStatus `json:"holdStatus,omitempty"`
	Id                   *string                                     `json:"id,omitempty"`
	LastIndexOperation   *EdiscoveryCaseIndexOperation               `json:"lastIndexOperation,omitempty"`
	LastModifiedDateTime *string                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                     `json:"@odata.type,omitempty"`
	ReleasedDateTime     *string                                     `json:"releasedDateTime,omitempty"`
	Status               *EdiscoveryNoncustodialDataSourceStatus     `json:"status,omitempty"`
}
