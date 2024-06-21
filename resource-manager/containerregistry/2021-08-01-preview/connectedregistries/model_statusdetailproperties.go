package connectedregistries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetailProperties struct {
	Code          *string `json:"code,omitempty"`
	CorrelationId *string `json:"correlationId,omitempty"`
	Description   *string `json:"description,omitempty"`
	Timestamp     *string `json:"timestamp,omitempty"`
	Type          *string `json:"type,omitempty"`
}
