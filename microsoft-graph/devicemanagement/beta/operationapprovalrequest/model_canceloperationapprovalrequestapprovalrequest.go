package operationapprovalrequest

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelOperationApprovalRequestApprovalRequest struct {
	ApprovalSource *beta.OperationApprovalSource `json:"approvalSource,omitempty"`
	Justification  nullable.Type[string]         `json:"justification,omitempty"`
}
