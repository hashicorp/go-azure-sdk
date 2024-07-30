package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryDataSourceContainer struct {
	CreatedDateTime      *string                                  `json:"createdDateTime,omitempty"`
	DisplayName          *string                                  `json:"displayName,omitempty"`
	HoldStatus           *EdiscoveryDataSourceContainerHoldStatus `json:"holdStatus,omitempty"`
	Id                   *string                                  `json:"id,omitempty"`
	LastIndexOperation   *EdiscoveryCaseIndexOperation            `json:"lastIndexOperation,omitempty"`
	LastModifiedDateTime *string                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                  `json:"@odata.type,omitempty"`
	ReleasedDateTime     *string                                  `json:"releasedDateTime,omitempty"`
	Status               *EdiscoveryDataSourceContainerStatus     `json:"status,omitempty"`
}
