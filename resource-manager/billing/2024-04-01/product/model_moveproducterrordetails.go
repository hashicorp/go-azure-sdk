package product

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveProductErrorDetails struct {
	Code    *MoveValidationErrorCode `json:"code,omitempty"`
	Details *string                  `json:"details,omitempty"`
	Message *string                  `json:"message,omitempty"`
}
