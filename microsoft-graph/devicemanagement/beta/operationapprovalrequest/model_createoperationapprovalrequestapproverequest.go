package operationapprovalrequest

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOperationApprovalRequestApproveRequest struct {
	ApprovalSource *beta.OperationApprovalSource `json:"approvalSource,omitempty"`
	Justification  nullable.Type[string]         `json:"justification,omitempty"`
}
