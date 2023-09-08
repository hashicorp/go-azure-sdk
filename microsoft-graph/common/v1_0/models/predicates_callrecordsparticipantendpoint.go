package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsParticipantEndpointOperationPredicate struct {
	CpuCoresCount          *int64
	CpuName                *string
	CpuProcessorSpeedInMhz *int64
	Name                   *string
	ODataType              *string
}

func (p CallRecordsParticipantEndpointOperationPredicate) Matches(input CallRecordsParticipantEndpoint) bool {

	if p.CpuCoresCount != nil && (input.CpuCoresCount == nil || *p.CpuCoresCount != *input.CpuCoresCount) {
		return false
	}

	if p.CpuName != nil && (input.CpuName == nil || *p.CpuName != *input.CpuName) {
		return false
	}

	if p.CpuProcessorSpeedInMhz != nil && (input.CpuProcessorSpeedInMhz == nil || *p.CpuProcessorSpeedInMhz != *input.CpuProcessorSpeedInMhz) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
