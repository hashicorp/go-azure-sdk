package userjoinedteampermissiongrant

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdJoinedTeamByIdPermissionGrantsByIdsRequest struct {
	Ids   *[]string `json:"ids,omitempty"`
	Types *[]string `json:"types,omitempty"`
}
