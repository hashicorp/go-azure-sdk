package embeddedsimactivationcodepool

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCodePoolAssignment struct {
	Id        *string                                 `json:"id,omitempty"`
	ODataType *string                                 `json:"@odata.type,omitempty"`
	Target    *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
