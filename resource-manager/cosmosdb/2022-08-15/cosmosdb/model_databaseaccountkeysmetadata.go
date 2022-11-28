package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseAccountKeysMetadata struct {
	PrimaryMasterKey           *AccountKeyMetadata `json:"primaryMasterKey"`
	PrimaryReadonlyMasterKey   *AccountKeyMetadata `json:"primaryReadonlyMasterKey"`
	SecondaryMasterKey         *AccountKeyMetadata `json:"secondaryMasterKey"`
	SecondaryReadonlyMasterKey *AccountKeyMetadata `json:"secondaryReadonlyMasterKey"`
}
