package operationapprovalrequest

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOperationApprovalRequestRetrieveRequestStatusRequest struct {
	EntityId   nullable.Type[string] `json:"entityId,omitempty"`
	EntityType nullable.Type[string] `json:"entityType,omitempty"`
}
