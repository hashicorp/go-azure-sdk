package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataLakeStoreAccountPropertiesBasic struct {
	AccountId         *string                     `json:"accountId,omitempty"`
	CreationTime      *string                     `json:"creationTime,omitempty"`
	Endpoint          *string                     `json:"endpoint,omitempty"`
	LastModifiedTime  *string                     `json:"lastModifiedTime,omitempty"`
	ProvisioningState *DataLakeStoreAccountStatus `json:"provisioningState,omitempty"`
	State             *DataLakeStoreAccountState  `json:"state,omitempty"`
}
