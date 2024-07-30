package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAudit struct {
	Action                       *RemoteActionAuditAction      `json:"action,omitempty"`
	ActionState                  *RemoteActionAuditActionState `json:"actionState,omitempty"`
	DeviceDisplayName            *string                       `json:"deviceDisplayName,omitempty"`
	DeviceIMEI                   *string                       `json:"deviceIMEI,omitempty"`
	DeviceOwnerUserPrincipalName *string                       `json:"deviceOwnerUserPrincipalName,omitempty"`
	Id                           *string                       `json:"id,omitempty"`
	InitiatedByUserPrincipalName *string                       `json:"initiatedByUserPrincipalName,omitempty"`
	ManagedDeviceId              *string                       `json:"managedDeviceId,omitempty"`
	ODataType                    *string                       `json:"@odata.type,omitempty"`
	RequestDateTime              *string                       `json:"requestDateTime,omitempty"`
	UserName                     *string                       `json:"userName,omitempty"`
}
