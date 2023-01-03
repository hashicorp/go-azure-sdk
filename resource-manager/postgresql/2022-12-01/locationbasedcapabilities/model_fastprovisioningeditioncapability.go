package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FastProvisioningEditionCapability struct {
	SupportedServerVersions *string `json:"supportedServerVersions,omitempty"`
	SupportedSku            *string `json:"supportedSku,omitempty"`
	SupportedStorageGb      *int64  `json:"supportedStorageGb,omitempty"`
}
