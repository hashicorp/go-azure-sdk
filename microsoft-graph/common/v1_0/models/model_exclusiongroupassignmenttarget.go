package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExclusionGroupAssignmentTarget struct {
	GroupId   *string `json:"groupId,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
