package purestoragepolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PureStoragePolicyProperties struct {
	ProvisioningState       *PureStoragePolicyProvisioningState `json:"provisioningState,omitempty"`
	StoragePolicyDefinition string                              `json:"storagePolicyDefinition"`
	StoragePoolId           string                              `json:"storagePoolId"`
}
