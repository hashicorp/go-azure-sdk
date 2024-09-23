package intent

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateIntentCopyRequest struct {
	Description nullable.Type[string] `json:"description,omitempty"`
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`
}
