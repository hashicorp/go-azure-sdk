package usermanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdManagedDeviceExecuteActionRequest struct {
	ActionName             *CreateUserByIdManagedDeviceExecuteActionRequestActionName `json:"actionName,omitempty"`
	CarrierUrl             *string                                                    `json:"carrierUrl,omitempty"`
	DeprovisionReason      *string                                                    `json:"deprovisionReason,omitempty"`
	DeviceIds              *[]string                                                  `json:"deviceIds,omitempty"`
	DeviceName             *string                                                    `json:"deviceName,omitempty"`
	KeepEnrollmentData     *bool                                                      `json:"keepEnrollmentData,omitempty"`
	KeepUserData           *bool                                                      `json:"keepUserData,omitempty"`
	NotificationBody       *string                                                    `json:"notificationBody,omitempty"`
	NotificationTitle      *string                                                    `json:"notificationTitle,omitempty"`
	OrganizationalUnitPath *string                                                    `json:"organizationalUnitPath,omitempty"`
	PersistEsimDataPlan    *bool                                                      `json:"persistEsimDataPlan,omitempty"`
}
