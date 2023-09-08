package mechatpermissiongrant

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMeChatByIdPermissionGrantsByIdsRequest struct {
	Ids   *[]string `json:"ids,omitempty"`
	Types *[]string `json:"types,omitempty"`
}
