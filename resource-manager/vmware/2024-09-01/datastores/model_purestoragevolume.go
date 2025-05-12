package datastores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PureStorageVolume struct {
	SizeGb        int64  `json:"sizeGb"`
	StoragePoolId string `json:"storagePoolId"`
}
