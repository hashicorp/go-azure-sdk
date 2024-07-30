package embeddedsimactivationcodepool

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementEmbeddedSIMActivationCodePoolsRequest struct {
	Assignments *[]EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
}
