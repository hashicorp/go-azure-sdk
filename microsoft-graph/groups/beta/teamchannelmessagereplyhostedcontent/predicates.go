package teamchannelmessagereplyhostedcontent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ChatMessageHostedContentOperationPredicate struct {
}

func (p ChatMessageHostedContentOperationPredicate) Matches(input beta.ChatMessageHostedContent) bool {

	return true
}
