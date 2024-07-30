package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilter struct {
	AssignmentFilterType *PayloadByFilterAssignmentFilterType `json:"assignmentFilterType,omitempty"`
	GroupId              *string                              `json:"groupId,omitempty"`
	ODataType            *string                              `json:"@odata.type,omitempty"`
	PayloadId            *string                              `json:"payloadId,omitempty"`
	PayloadType          *PayloadByFilterPayloadType          `json:"payloadType,omitempty"`
}
