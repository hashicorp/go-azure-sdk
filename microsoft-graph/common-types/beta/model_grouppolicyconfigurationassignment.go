package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyConfigurationAssignment struct {
	Id                   *string                                 `json:"id,omitempty"`
	LastModifiedDateTime *string                                 `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	Target               *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
