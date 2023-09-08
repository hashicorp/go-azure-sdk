package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsProtectionStateOperationPredicate struct {
	AntiMalwareVersion             *string
	EngineVersion                  *string
	FullScanOverdue                *bool
	FullScanRequired               *bool
	Id                             *string
	IsVirtualMachine               *bool
	LastFullScanDateTime           *string
	LastFullScanSignatureVersion   *string
	LastQuickScanDateTime          *string
	LastQuickScanSignatureVersion  *string
	LastReportedDateTime           *string
	MalwareProtectionEnabled       *bool
	NetworkInspectionSystemEnabled *bool
	ODataType                      *string
	QuickScanOverdue               *bool
	RealTimeProtectionEnabled      *bool
	RebootRequired                 *bool
	SignatureUpdateOverdue         *bool
	SignatureVersion               *string
	TamperProtectionEnabled        *bool
}

func (p WindowsProtectionStateOperationPredicate) Matches(input WindowsProtectionState) bool {

	if p.AntiMalwareVersion != nil && (input.AntiMalwareVersion == nil || *p.AntiMalwareVersion != *input.AntiMalwareVersion) {
		return false
	}

	if p.EngineVersion != nil && (input.EngineVersion == nil || *p.EngineVersion != *input.EngineVersion) {
		return false
	}

	if p.FullScanOverdue != nil && (input.FullScanOverdue == nil || *p.FullScanOverdue != *input.FullScanOverdue) {
		return false
	}

	if p.FullScanRequired != nil && (input.FullScanRequired == nil || *p.FullScanRequired != *input.FullScanRequired) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsVirtualMachine != nil && (input.IsVirtualMachine == nil || *p.IsVirtualMachine != *input.IsVirtualMachine) {
		return false
	}

	if p.LastFullScanDateTime != nil && (input.LastFullScanDateTime == nil || *p.LastFullScanDateTime != *input.LastFullScanDateTime) {
		return false
	}

	if p.LastFullScanSignatureVersion != nil && (input.LastFullScanSignatureVersion == nil || *p.LastFullScanSignatureVersion != *input.LastFullScanSignatureVersion) {
		return false
	}

	if p.LastQuickScanDateTime != nil && (input.LastQuickScanDateTime == nil || *p.LastQuickScanDateTime != *input.LastQuickScanDateTime) {
		return false
	}

	if p.LastQuickScanSignatureVersion != nil && (input.LastQuickScanSignatureVersion == nil || *p.LastQuickScanSignatureVersion != *input.LastQuickScanSignatureVersion) {
		return false
	}

	if p.LastReportedDateTime != nil && (input.LastReportedDateTime == nil || *p.LastReportedDateTime != *input.LastReportedDateTime) {
		return false
	}

	if p.MalwareProtectionEnabled != nil && (input.MalwareProtectionEnabled == nil || *p.MalwareProtectionEnabled != *input.MalwareProtectionEnabled) {
		return false
	}

	if p.NetworkInspectionSystemEnabled != nil && (input.NetworkInspectionSystemEnabled == nil || *p.NetworkInspectionSystemEnabled != *input.NetworkInspectionSystemEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QuickScanOverdue != nil && (input.QuickScanOverdue == nil || *p.QuickScanOverdue != *input.QuickScanOverdue) {
		return false
	}

	if p.RealTimeProtectionEnabled != nil && (input.RealTimeProtectionEnabled == nil || *p.RealTimeProtectionEnabled != *input.RealTimeProtectionEnabled) {
		return false
	}

	if p.RebootRequired != nil && (input.RebootRequired == nil || *p.RebootRequired != *input.RebootRequired) {
		return false
	}

	if p.SignatureUpdateOverdue != nil && (input.SignatureUpdateOverdue == nil || *p.SignatureUpdateOverdue != *input.SignatureUpdateOverdue) {
		return false
	}

	if p.SignatureVersion != nil && (input.SignatureVersion == nil || *p.SignatureVersion != *input.SignatureVersion) {
		return false
	}

	if p.TamperProtectionEnabled != nil && (input.TamperProtectionEnabled == nil || *p.TamperProtectionEnabled != *input.TamperProtectionEnabled) {
		return false
	}

	return true
}
