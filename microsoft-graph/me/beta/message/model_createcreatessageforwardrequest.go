package message

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateCreatessageForwardRequest struct {
	Comment      nullable.Type[string] `json:"Comment,omitempty"`
	Message      *beta.Message         `json:"Message,omitempty"`
	ToRecipients *[]beta.Recipient     `json:"ToRecipients,omitempty"`
}
