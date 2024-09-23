package teamscheduletimeoffrequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type TimeOffRequestOperationPredicate struct {
}

func (p TimeOffRequestOperationPredicate) Matches(input stable.TimeOffRequest) bool {

	return true
}
