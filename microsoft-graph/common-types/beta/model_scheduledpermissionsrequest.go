package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledPermissionsRequest struct {
	CreatedDateTime      *string                                  `json:"createdDateTime,omitempty"`
	Id                   *string                                  `json:"id,omitempty"`
	Justification        *string                                  `json:"justification,omitempty"`
	Notes                *string                                  `json:"notes,omitempty"`
	ODataType            *string                                  `json:"@odata.type,omitempty"`
	RequestedPermissions *PermissionsDefinition                   `json:"requestedPermissions,omitempty"`
	ScheduleInfo         *RequestSchedule                         `json:"scheduleInfo,omitempty"`
	StatusDetail         *ScheduledPermissionsRequestStatusDetail `json:"statusDetail,omitempty"`
	TicketInfo           *TicketInfo                              `json:"ticketInfo,omitempty"`
}
