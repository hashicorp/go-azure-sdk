package intent

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateIntentMigrateToTemplateRequest struct {
	NewTemplateId        nullable.Type[string] `json:"newTemplateId,omitempty"`
	PreserveCustomValues *bool                 `json:"preserveCustomValues,omitempty"`
}
