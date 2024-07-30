package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsRequestChange struct {
	ActiveOccurrenceStatus *PermissionsRequestChangeActiveOccurrenceStatus `json:"activeOccurrenceStatus,omitempty"`
	Id                     *string                                         `json:"id,omitempty"`
	ModificationDateTime   *string                                         `json:"modificationDateTime,omitempty"`
	ODataType              *string                                         `json:"@odata.type,omitempty"`
	PermissionsRequestId   *string                                         `json:"permissionsRequestId,omitempty"`
	StatusDetail           *PermissionsRequestChangeStatusDetail           `json:"statusDetail,omitempty"`
	TicketId               *string                                         `json:"ticketId,omitempty"`
}
