package databaseextensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseExtensionsProperties struct {
	OperationMode  OperationMode  `json:"operationMode"`
	StorageKey     string         `json:"storageKey"`
	StorageKeyType StorageKeyType `json:"storageKeyType"`
	StorageUri     string         `json:"storageUri"`
}
