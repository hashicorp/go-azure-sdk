package embeddedsimactivationcodepool

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type EmbeddedSIMActivationCodePoolOperationPredicate struct {
}

func (p EmbeddedSIMActivationCodePoolOperationPredicate) Matches(input beta.EmbeddedSIMActivationCodePool) bool {

	return true
}

type EmbeddedSIMActivationCodePoolAssignmentOperationPredicate struct {
}

func (p EmbeddedSIMActivationCodePoolAssignmentOperationPredicate) Matches(input beta.EmbeddedSIMActivationCodePoolAssignment) bool {

	return true
}
