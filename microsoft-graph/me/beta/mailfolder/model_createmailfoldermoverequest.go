package mailfolder

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMailFolderMoveRequest struct {
	DestinationId *string `json:"DestinationId,omitempty"`
}
