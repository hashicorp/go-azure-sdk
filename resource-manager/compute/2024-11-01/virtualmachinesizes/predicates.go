package virtualmachinesizes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineSizeOperationPredicate struct {
	MaxDataDiskCount     *int64
	MemoryInMB           *int64
	Name                 *string
	NumberOfCores        *int64
	OsDiskSizeInMB       *int64
	ResourceDiskSizeInMB *int64
}

func (p VirtualMachineSizeOperationPredicate) Matches(input VirtualMachineSize) bool {

	if p.MaxDataDiskCount != nil && (input.MaxDataDiskCount == nil || *p.MaxDataDiskCount != *input.MaxDataDiskCount) {
		return false
	}

	if p.MemoryInMB != nil && (input.MemoryInMB == nil || *p.MemoryInMB != *input.MemoryInMB) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.NumberOfCores != nil && (input.NumberOfCores == nil || *p.NumberOfCores != *input.NumberOfCores) {
		return false
	}

	if p.OsDiskSizeInMB != nil && (input.OsDiskSizeInMB == nil || *p.OsDiskSizeInMB != *input.OsDiskSizeInMB) {
		return false
	}

	if p.ResourceDiskSizeInMB != nil && (input.ResourceDiskSizeInMB == nil || *p.ResourceDiskSizeInMB != *input.ResourceDiskSizeInMB) {
		return false
	}

	return true
}
