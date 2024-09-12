package embeddedsimactivationcodepool

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignEmbeddedSIMActivationCodePoolsRequest struct {
	Assignments *[]beta.EmbeddedSIMActivationCodePoolAssignment `json:"assignments,omitempty"`
}
