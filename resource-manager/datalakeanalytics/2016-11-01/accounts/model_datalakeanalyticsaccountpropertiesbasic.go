package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeAnalyticsAccountPropertiesBasic struct {
	AccountId         *string                         `json:"accountId,omitempty"`
	CreationTime      *string                         `json:"creationTime,omitempty"`
	Endpoint          *string                         `json:"endpoint,omitempty"`
	LastModifiedTime  *string                         `json:"lastModifiedTime,omitempty"`
	ProvisioningState *DataLakeAnalyticsAccountStatus `json:"provisioningState,omitempty"`
	State             *DataLakeAnalyticsAccountState  `json:"state,omitempty"`
}
