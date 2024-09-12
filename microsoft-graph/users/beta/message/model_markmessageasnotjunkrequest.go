package message

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarkMessageAsNotJunkRequest struct {
	MoveToInbox *bool `json:"MoveToInbox,omitempty"`
}
