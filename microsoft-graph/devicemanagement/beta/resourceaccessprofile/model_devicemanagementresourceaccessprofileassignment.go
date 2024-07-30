package resourceaccessprofile

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileAssignment struct {
	Id        *string                                                `json:"id,omitempty"`
	Intent    *DeviceManagementResourceAccessProfileAssignmentIntent `json:"intent,omitempty"`
	ODataType *string                                                `json:"@odata.type,omitempty"`
	SourceId  *string                                                `json:"sourceId,omitempty"`
	Target    *DeviceAndAppManagementAssignmentTarget                `json:"target,omitempty"`
}
