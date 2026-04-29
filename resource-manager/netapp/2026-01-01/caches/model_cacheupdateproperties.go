package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheUpdateProperties struct {
	CifsChangeNotifications           *CifsChangeNotifyState       `json:"cifsChangeNotifications,omitempty"`
	ExportPolicy                      *CachePropertiesExportPolicy `json:"exportPolicy,omitempty"`
	KeyVaultPrivateEndpointResourceId *string                      `json:"keyVaultPrivateEndpointResourceId,omitempty"`
	ProtocolTypes                     *[]ProtocolTypes             `json:"protocolTypes,omitempty"`
	Size                              *int64                       `json:"size,omitempty"`
	SmbSettings                       *SmbSettings                 `json:"smbSettings,omitempty"`
	ThroughputMibps                   *float64                     `json:"throughputMibps,omitempty"`
	WriteBack                         *EnableWriteBackState        `json:"writeBack,omitempty"`
}
