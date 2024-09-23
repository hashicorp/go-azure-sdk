package siteoperation

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type RichLongRunningOperationOperationPredicate struct {
}

func (p RichLongRunningOperationOperationPredicate) Matches(input beta.RichLongRunningOperation) bool {

	return true
}
