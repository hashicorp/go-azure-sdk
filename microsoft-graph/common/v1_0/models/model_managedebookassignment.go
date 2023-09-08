package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedEBookAssignment struct {
	Id            *string                                 `json:"id,omitempty"`
	InstallIntent *ManagedEBookAssignmentInstallIntent    `json:"installIntent,omitempty"`
	ODataType     *string                                 `json:"@odata.type,omitempty"`
	Target        *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
