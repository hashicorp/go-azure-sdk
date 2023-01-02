package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlexibleServerEditionCapability struct {
	Name                     *string                     `json:"name,omitempty"`
	Status                   *string                     `json:"status,omitempty"`
	SupportedServerVersions  *[]ServerVersionCapability  `json:"supportedServerVersions,omitempty"`
	SupportedStorageEditions *[]StorageEditionCapability `json:"supportedStorageEditions,omitempty"`
}
