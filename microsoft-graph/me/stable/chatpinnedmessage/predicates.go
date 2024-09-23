package chatpinnedmessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type PinnedChatMessageInfoOperationPredicate struct {
}

func (p PinnedChatMessageInfoOperationPredicate) Matches(input stable.PinnedChatMessageInfo) bool {

	return true
}
