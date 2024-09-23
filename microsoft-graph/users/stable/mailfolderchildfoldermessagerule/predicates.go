package mailfolderchildfoldermessagerule

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type MessageRuleOperationPredicate struct {
}

func (p MessageRuleOperationPredicate) Matches(input stable.MessageRule) bool {

	return true
}
