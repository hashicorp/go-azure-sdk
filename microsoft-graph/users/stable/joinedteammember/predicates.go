package joinedteammember

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type ActionResultPartOperationPredicate struct {
}

func (p ActionResultPartOperationPredicate) Matches(input stable.ActionResultPart) bool {

	return true
}

type ConversationMemberOperationPredicate struct {
}

func (p ConversationMemberOperationPredicate) Matches(input stable.ConversationMember) bool {

	return true
}
