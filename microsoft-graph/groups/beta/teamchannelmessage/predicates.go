package teamchannelmessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type ActionResultPartOperationPredicate struct {
}

func (p ActionResultPartOperationPredicate) Matches(input beta.ActionResultPart) bool {

	return true
}

type ChatMessageOperationPredicate struct {
}

func (p ChatMessageOperationPredicate) Matches(input beta.ChatMessage) bool {

	return true
}
