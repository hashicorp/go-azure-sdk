package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDriverUpdateProfile struct {
	ApprovalType             *WindowsDriverUpdateProfileApprovalType        `json:"approvalType,omitempty"`
	Assignments              *[]WindowsDriverUpdateProfileAssignment        `json:"assignments,omitempty"`
	CreatedDateTime          *string                                        `json:"createdDateTime,omitempty"`
	DeploymentDeferralInDays *int64                                         `json:"deploymentDeferralInDays,omitempty"`
	Description              *string                                        `json:"description,omitempty"`
	DeviceReporting          *int64                                         `json:"deviceReporting,omitempty"`
	DisplayName              *string                                        `json:"displayName,omitempty"`
	DriverInventories        *[]WindowsDriverUpdateInventory                `json:"driverInventories,omitempty"`
	Id                       *string                                        `json:"id,omitempty"`
	InventorySyncStatus      *WindowsDriverUpdateProfileInventorySyncStatus `json:"inventorySyncStatus,omitempty"`
	LastModifiedDateTime     *string                                        `json:"lastModifiedDateTime,omitempty"`
	NewUpdates               *int64                                         `json:"newUpdates,omitempty"`
	ODataType                *string                                        `json:"@odata.type,omitempty"`
	RoleScopeTagIds          *[]string                                      `json:"roleScopeTagIds,omitempty"`
}
