package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedTrainingInfo struct {
	AssignedUserCount  *int64  `json:"assignedUserCount,omitempty"`
	CompletedUserCount *int64  `json:"completedUserCount,omitempty"`
	DisplayName        *string `json:"displayName,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
}
