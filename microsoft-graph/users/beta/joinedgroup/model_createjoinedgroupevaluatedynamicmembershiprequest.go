package joinedgroup

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedGroupEvaluateDynamicMembershipRequest struct {
	MemberId       *string `json:"memberId,omitempty"`
	MembershipRule *string `json:"membershipRule,omitempty"`
}
