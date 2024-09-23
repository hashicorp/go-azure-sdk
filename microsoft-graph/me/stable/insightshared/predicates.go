package insightshared

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type SharedInsightOperationPredicate struct {
}

func (p SharedInsightOperationPredicate) Matches(input stable.SharedInsight) bool {

	return true
}
