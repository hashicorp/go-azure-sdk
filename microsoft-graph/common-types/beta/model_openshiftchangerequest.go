package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftChangeRequest struct {
	AssignedTo            *OpenShiftChangeRequestAssignedTo `json:"assignedTo,omitempty"`
	CreatedBy             *IdentitySet                      `json:"createdBy,omitempty"`
	CreatedDateTime       *string                           `json:"createdDateTime,omitempty"`
	Id                    *string                           `json:"id,omitempty"`
	LastModifiedBy        *IdentitySet                      `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime  *string                           `json:"lastModifiedDateTime,omitempty"`
	ManagerActionDateTime *string                           `json:"managerActionDateTime,omitempty"`
	ManagerActionMessage  *string                           `json:"managerActionMessage,omitempty"`
	ManagerUserId         *string                           `json:"managerUserId,omitempty"`
	ODataType             *string                           `json:"@odata.type,omitempty"`
	OpenShiftId           *string                           `json:"openShiftId,omitempty"`
	SenderDateTime        *string                           `json:"senderDateTime,omitempty"`
	SenderMessage         *string                           `json:"senderMessage,omitempty"`
	SenderUserId          *string                           `json:"senderUserId,omitempty"`
	State                 *OpenShiftChangeRequestState      `json:"state,omitempty"`
}
