package mailfolder

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMailFolderAllMessagesReadStateRequest struct {
	IsRead               nullable.Type[bool] `json:"isRead,omitempty"`
	SuppressReadReceipts nullable.Type[bool] `json:"suppressReadReceipts,omitempty"`
}
