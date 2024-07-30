package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffRequest struct {
	AssignedTo            *TimeOffRequestAssignedTo `json:"assignedTo,omitempty"`
	CreatedDateTime       *string                   `json:"createdDateTime,omitempty"`
	EndDateTime           *string                   `json:"endDateTime,omitempty"`
	Id                    *string                   `json:"id,omitempty"`
	LastModifiedBy        *IdentitySet              `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime  *string                   `json:"lastModifiedDateTime,omitempty"`
	ManagerActionDateTime *string                   `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage  *string                   `json:"managerActionMessage,omitempty"`
	ManagerUserId         *string                   `json:"managerUserId,omitempty"`
	ODataType             *string                   `json:"@odata.type,omitempty"`
	SenderDateTime        *string                   `json:"senderDateTime,omitempty"`
	SenderMessage         *string                   `json:"senderMessage,omitempty"`
	SenderUserId          *string                   `json:"senderUserId,omitempty"`
	StartDateTime         *string                   `json:"startDateTime,omitempty"`
	State                 *TimeOffRequestState      `json:"state,omitempty"`
	TimeOffReasonId       *string                   `json:"timeOffReasonId,omitempty"`
}
