package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessScheduleRequest struct {
	Action            *PrivilegedAccessScheduleRequestAction `json:"action,omitempty"`
	ApprovalId        *string                                `json:"approvalId,omitempty"`
	CompletedDateTime *string                                `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet                           `json:"createdBy,omitempty"`
	CreatedDateTime   *string                                `json:"createdDateTime,omitempty"`
	CustomData        *string                                `json:"customData,omitempty"`
	Id                *string                                `json:"id,omitempty"`
	IsValidationOnly  *bool                                  `json:"isValidationOnly,omitempty"`
	Justification     *string                                `json:"justification,omitempty"`
	ODataType         *string                                `json:"@odata.type,omitempty"`
	ScheduleInfo      *RequestSchedule                       `json:"scheduleInfo,omitempty"`
	Status            *string                                `json:"status,omitempty"`
	TicketInfo        *TicketInfo                            `json:"ticketInfo,omitempty"`
}
