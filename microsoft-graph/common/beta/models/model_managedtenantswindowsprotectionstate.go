package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWindowsProtectionState struct {
	AntiMalwareVersion             *string `json:"antiMalwareVersion,omitempty"`
	AttentionRequired              *bool   `json:"attentionRequired,omitempty"`
	DeviceDeleted                  *bool   `json:"deviceDeleted,omitempty"`
	DevicePropertyRefreshDateTime  *string `json:"devicePropertyRefreshDateTime,omitempty"`
	EngineVersion                  *string `json:"engineVersion,omitempty"`
	FullScanOverdue                *bool   `json:"fullScanOverdue,omitempty"`
	FullScanRequired               *bool   `json:"fullScanRequired,omitempty"`
	Id                             *string `json:"id,omitempty"`
	LastFullScanDateTime           *string `json:"lastFullScanDateTime,omitempty"`
	LastFullScanSignatureVersion   *string `json:"lastFullScanSignatureVersion,omitempty"`
	LastQuickScanDateTime          *string `json:"lastQuickScanDateTime,omitempty"`
	LastQuickScanSignatureVersion  *string `json:"lastQuickScanSignatureVersion,omitempty"`
	LastRefreshedDateTime          *string `json:"lastRefreshedDateTime,omitempty"`
	LastReportedDateTime           *string `json:"lastReportedDateTime,omitempty"`
	MalwareProtectionEnabled       *bool   `json:"malwareProtectionEnabled,omitempty"`
	ManagedDeviceHealthState       *string `json:"managedDeviceHealthState,omitempty"`
	ManagedDeviceId                *string `json:"managedDeviceId,omitempty"`
	ManagedDeviceName              *string `json:"managedDeviceName,omitempty"`
	NetworkInspectionSystemEnabled *bool   `json:"networkInspectionSystemEnabled,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	QuickScanOverdue               *bool   `json:"quickScanOverdue,omitempty"`
	RealTimeProtectionEnabled      *bool   `json:"realTimeProtectionEnabled,omitempty"`
	RebootRequired                 *bool   `json:"rebootRequired,omitempty"`
	SignatureUpdateOverdue         *bool   `json:"signatureUpdateOverdue,omitempty"`
	SignatureVersion               *string `json:"signatureVersion,omitempty"`
	TenantDisplayName              *string `json:"tenantDisplayName,omitempty"`
	TenantId                       *string `json:"tenantId,omitempty"`
}
