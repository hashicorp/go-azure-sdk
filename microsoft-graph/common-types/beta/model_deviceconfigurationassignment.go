package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignment struct {
	Id        *string                                 `json:"id,omitempty"`
	Intent    *DeviceConfigurationAssignmentIntent    `json:"intent,omitempty"`
	ODataType *string                                 `json:"@odata.type,omitempty"`
	Source    *DeviceConfigurationAssignmentSource    `json:"source,omitempty"`
	SourceId  *string                                 `json:"sourceId,omitempty"`
	Target    *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
