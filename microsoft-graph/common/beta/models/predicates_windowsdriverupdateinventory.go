package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateInventoryOperationPredicate struct {
	ApplicableDeviceCount *int64
	DeployDateTime        *string
	DriverClass           *string
	Id                    *string
	Manufacturer          *string
	Name                  *string
	ODataType             *string
	ReleaseDateTime       *string
	Version               *string
}

func (p WindowsDriverUpdateInventoryOperationPredicate) Matches(input WindowsDriverUpdateInventory) bool {

	if p.ApplicableDeviceCount != nil && (input.ApplicableDeviceCount == nil || *p.ApplicableDeviceCount != *input.ApplicableDeviceCount) {
		return false
	}

	if p.DeployDateTime != nil && (input.DeployDateTime == nil || *p.DeployDateTime != *input.DeployDateTime) {
		return false
	}

	if p.DriverClass != nil && (input.DriverClass == nil || *p.DriverClass != *input.DriverClass) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Manufacturer != nil && (input.Manufacturer == nil || *p.Manufacturer != *input.Manufacturer) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReleaseDateTime != nil && (input.ReleaseDateTime == nil || *p.ReleaseDateTime != *input.ReleaseDateTime) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
