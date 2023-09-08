package userjoinedgroup

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdJoinedGroupsUserOwnedObjectRequest struct {
	Type   *string `json:"type,omitempty"`
	UserId *string `json:"userId,omitempty"`
}
