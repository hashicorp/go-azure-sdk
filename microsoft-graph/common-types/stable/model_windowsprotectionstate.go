package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsProtectionState struct {
	AntiMalwareVersion             *string                              `json:"antiMalwareVersion,omitempty"`
	DetectedMalwareState           *[]WindowsDeviceMalwareState         `json:"detectedMalwareState,omitempty"`
	DeviceState                    *WindowsProtectionStateDeviceState   `json:"deviceState,omitempty"`
	EngineVersion                  *string                              `json:"engineVersion,omitempty"`
	FullScanOverdue                *bool                                `json:"fullScanOverdue,omitempty"`
	FullScanRequired               *bool                                `json:"fullScanRequired,omitempty"`
	Id                             *string                              `json:"id,omitempty"`
	IsVirtualMachine               *bool                                `json:"isVirtualMachine,omitempty"`
	LastFullScanDateTime           *string                              `json:"lastFullScanDateTime,omitempty"`
	LastFullScanSignatureVersion   *string                              `json:"lastFullScanSignatureVersion,omitempty"`
	LastQuickScanDateTime          *string                              `json:"lastQuickScanDateTime,omitempty"`
	LastQuickScanSignatureVersion  *string                              `json:"lastQuickScanSignatureVersion,omitempty"`
	LastReportedDateTime           *string                              `json:"lastReportedDateTime,omitempty"`
	MalwareProtectionEnabled       *bool                                `json:"malwareProtectionEnabled,omitempty"`
	NetworkInspectionSystemEnabled *bool                                `json:"networkInspectionSystemEnabled,omitempty"`
	ODataType                      *string                              `json:"@odata.type,omitempty"`
	ProductStatus                  *WindowsProtectionStateProductStatus `json:"productStatus,omitempty"`
	QuickScanOverdue               *bool                                `json:"quickScanOverdue,omitempty"`
	RealTimeProtectionEnabled      *bool                                `json:"realTimeProtectionEnabled,omitempty"`
	RebootRequired                 *bool                                `json:"rebootRequired,omitempty"`
	SignatureUpdateOverdue         *bool                                `json:"signatureUpdateOverdue,omitempty"`
	SignatureVersion               *string                              `json:"signatureVersion,omitempty"`
	TamperProtectionEnabled        *bool                                `json:"tamperProtectionEnabled,omitempty"`
}
