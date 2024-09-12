package plannerplan

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePlannerPlanArchiveRequest struct {
	Justification nullable.Type[string] `json:"justification,omitempty"`
}
