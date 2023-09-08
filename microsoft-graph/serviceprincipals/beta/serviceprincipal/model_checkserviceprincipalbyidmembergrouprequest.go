package serviceprincipal

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckServicePrincipalByIdMemberGroupRequest struct {
	GroupIds *[]string `json:"groupIds,omitempty"`
}
