package manageddatabasesecurityevents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventSqlInjectionAdditionalProperties struct {
	ErrorCode                *int64  `json:"errorCode,omitempty"`
	ErrorMessage             *string `json:"errorMessage,omitempty"`
	ErrorSeverity            *int64  `json:"errorSeverity,omitempty"`
	Statement                *string `json:"statement,omitempty"`
	StatementHighlightLength *int64  `json:"statementHighlightLength,omitempty"`
	StatementHighlightOffset *int64  `json:"statementHighlightOffset,omitempty"`
	ThreatId                 *string `json:"threatId,omitempty"`
}
