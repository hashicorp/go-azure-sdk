package devicemanagement

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAssignmentFiltersStatusDetailRequest struct {
	AssignmentFilterIds *[]string             `json:"assignmentFilterIds,omitempty"`
	ManagedDeviceId     nullable.Type[string] `json:"managedDeviceId,omitempty"`
	PayloadId           nullable.Type[string] `json:"payloadId,omitempty"`
	Skip                nullable.Type[int64]  `json:"skip,omitempty"`
	Top                 nullable.Type[int64]  `json:"top,omitempty"`
	UserId              nullable.Type[string] `json:"userId,omitempty"`
}
