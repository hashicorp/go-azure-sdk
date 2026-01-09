package messageattachment

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMessageAttachmentsUploadSessionRequest struct {
	AttachmentItem *stable.AttachmentItem `json:"AttachmentItem,omitempty"`
}
