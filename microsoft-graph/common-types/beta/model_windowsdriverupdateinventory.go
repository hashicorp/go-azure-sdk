package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateInventory struct {
	ApplicableDeviceCount *int64                                      `json:"applicableDeviceCount,omitempty"`
	ApprovalStatus        *WindowsDriverUpdateInventoryApprovalStatus `json:"approvalStatus,omitempty"`
	Category              *WindowsDriverUpdateInventoryCategory       `json:"category,omitempty"`
	DeployDateTime        *string                                     `json:"deployDateTime,omitempty"`
	DriverClass           *string                                     `json:"driverClass,omitempty"`
	Id                    *string                                     `json:"id,omitempty"`
	Manufacturer          *string                                     `json:"manufacturer,omitempty"`
	Name                  *string                                     `json:"name,omitempty"`
	ODataType             *string                                     `json:"@odata.type,omitempty"`
	ReleaseDateTime       *string                                     `json:"releaseDateTime,omitempty"`
	Version               *string                                     `json:"version,omitempty"`
}
