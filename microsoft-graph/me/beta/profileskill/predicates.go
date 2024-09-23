package profileskill

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type SkillProficiencyOperationPredicate struct {
}

func (p SkillProficiencyOperationPredicate) Matches(input beta.SkillProficiency) bool {

	return true
}
