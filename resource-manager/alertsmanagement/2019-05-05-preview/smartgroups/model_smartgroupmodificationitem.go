package smartgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmartGroupModificationItem struct {
	Comments          *string                      `json:"comments,omitempty"`
	Description       *string                      `json:"description,omitempty"`
	ModificationEvent *SmartGroupModificationEvent `json:"modificationEvent,omitempty"`
	ModifiedAt        *string                      `json:"modifiedAt,omitempty"`
	ModifiedBy        *string                      `json:"modifiedBy,omitempty"`
	NewValue          *string                      `json:"newValue,omitempty"`
	OldValue          *string                      `json:"oldValue,omitempty"`
}
