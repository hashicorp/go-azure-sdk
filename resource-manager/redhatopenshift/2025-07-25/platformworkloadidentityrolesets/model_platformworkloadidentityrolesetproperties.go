package platformworkloadidentityrolesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformWorkloadIdentityRoleSetProperties struct {
	OpenShiftVersion              *string                         `json:"openShiftVersion,omitempty"`
	PlatformWorkloadIdentityRoles *[]PlatformWorkloadIdentityRole `json:"platformWorkloadIdentityRoles,omitempty"`
}
