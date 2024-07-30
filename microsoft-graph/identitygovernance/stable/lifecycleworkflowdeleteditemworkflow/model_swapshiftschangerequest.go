package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwapShiftsChangeRequest struct {
	AssignedTo              *SwapShiftsChangeRequestAssignedTo `json:"assignedTo,omitempty"`
	CreatedDateTime         *string                            `json:"createdDateTime,omitempty"`
	Id                      *string                            `json:"id,omitempty"`
	LastModifiedBy          *IdentitySet                       `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime    *string                            `json:"lastModifiedDateTime,omitempty"`
	ManagerActionDateTime   *string                            `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage    *string                            `json:"managerActionMessage,omitempty"`
	ManagerUserId           *string                            `json:"managerUserId,omitempty"`
	ODataType               *string                            `json:"@odata.type,omitempty"`
	RecipientActionDateTime *string                            `json:"recipientActionDateTime,omitempty"`
	RecipientActionMessage  *string                            `json:"recipientActionMessage,omitempty"`
	RecipientShiftId        *string                            `json:"recipientShiftId,omitempty"`
	RecipientUserId         *string                            `json:"recipientUserId,omitempty"`
	SenderDateTime          *string                            `json:"senderDateTime,omitempty"`
	SenderMessage           *string                            `json:"senderMessage,omitempty"`
	SenderShiftId           *string                            `json:"senderShiftId,omitempty"`
	SenderUserId            *string                            `json:"senderUserId,omitempty"`
	State                   *SwapShiftsChangeRequestState      `json:"state,omitempty"`
}
