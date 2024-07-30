package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryNoncustodialDataSource struct {
	CreatedDateTime      *string                                             `json:"createdDateTime,omitempty"`
	DataSource           *SecurityDataSource                                 `json:"dataSource,omitempty"`
	DisplayName          *string                                             `json:"displayName,omitempty"`
	HoldStatus           *SecurityEdiscoveryNoncustodialDataSourceHoldStatus `json:"holdStatus,omitempty"`
	Id                   *string                                             `json:"id,omitempty"`
	LastIndexOperation   *SecurityEdiscoveryIndexOperation                   `json:"lastIndexOperation,omitempty"`
	LastModifiedDateTime *string                                             `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                             `json:"@odata.type,omitempty"`
	ReleasedDateTime     *string                                             `json:"releasedDateTime,omitempty"`
	Status               *SecurityEdiscoveryNoncustodialDataSourceStatus     `json:"status,omitempty"`
}
