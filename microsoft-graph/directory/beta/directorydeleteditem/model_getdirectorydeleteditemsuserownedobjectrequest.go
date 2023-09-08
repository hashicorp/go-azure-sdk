package directorydeleteditem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDirectoryDeletedItemsUserOwnedObjectRequest struct {
	Type   *string `json:"type,omitempty"`
	UserId *string `json:"userId,omitempty"`
}
