package mepresence

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetMePresenceStatusMessageRequest struct {
	StatusMessage *PresenceStatusMessage `json:"statusMessage,omitempty"`
}
