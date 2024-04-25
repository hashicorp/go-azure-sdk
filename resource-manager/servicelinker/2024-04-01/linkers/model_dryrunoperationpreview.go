package linkers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DryrunOperationPreview struct {
	Action        *string                     `json:"action,omitempty"`
	Description   *string                     `json:"description,omitempty"`
	Name          *string                     `json:"name,omitempty"`
	OperationType *DryrunPreviewOperationType `json:"operationType,omitempty"`
	Scope         *string                     `json:"scope,omitempty"`
}
