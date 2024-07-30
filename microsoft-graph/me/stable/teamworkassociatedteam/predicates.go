package teamworkassociatedteam

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type AssociatedTeamInfoOperationPredicate struct {
}

func (p AssociatedTeamInfoOperationPredicate) Matches(input stable.AssociatedTeamInfo) bool {

	return true
}
