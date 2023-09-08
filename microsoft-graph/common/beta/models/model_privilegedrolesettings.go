package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleSettings struct {
	ApprovalOnElevation           *bool     `json:"approvalOnElevation,omitempty"`
	ApproverIds                   *[]string `json:"approverIds,omitempty"`
	ElevationDuration             *string   `json:"elevationDuration,omitempty"`
	Id                            *string   `json:"id,omitempty"`
	IsMfaOnElevationConfigurable  *bool     `json:"isMfaOnElevationConfigurable,omitempty"`
	LastGlobalAdmin               *bool     `json:"lastGlobalAdmin,omitempty"`
	MaxElavationDuration          *string   `json:"maxElavationDuration,omitempty"`
	MfaOnElevation                *bool     `json:"mfaOnElevation,omitempty"`
	MinElevationDuration          *string   `json:"minElevationDuration,omitempty"`
	NotificationToUserOnElevation *bool     `json:"notificationToUserOnElevation,omitempty"`
	ODataType                     *string   `json:"@odata.type,omitempty"`
	TicketingInfoOnElevation      *bool     `json:"ticketingInfoOnElevation,omitempty"`
}
