package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedApproval struct {
	ApprovalDuration *string                          `json:"approvalDuration,omitempty"`
	ApprovalState    *PrivilegedApprovalApprovalState `json:"approvalState,omitempty"`
	ApprovalType     *string                          `json:"approvalType,omitempty"`
	ApproverReason   *string                          `json:"approverReason,omitempty"`
	EndDateTime      *string                          `json:"endDateTime,omitempty"`
	Id               *string                          `json:"id,omitempty"`
	ODataType        *string                          `json:"@odata.type,omitempty"`
	Request          *PrivilegedRoleAssignmentRequest `json:"request,omitempty"`
	RequestorReason  *string                          `json:"requestorReason,omitempty"`
	RoleId           *string                          `json:"roleId,omitempty"`
	RoleInfo         *PrivilegedRole                  `json:"roleInfo,omitempty"`
	StartDateTime    *string                          `json:"startDateTime,omitempty"`
	UserId           *string                          `json:"userId,omitempty"`
}
