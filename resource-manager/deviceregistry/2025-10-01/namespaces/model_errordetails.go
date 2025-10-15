package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ErrorDetails struct {
	Code          *string `json:"code,omitempty"`
	CorrelationId *string `json:"correlationId,omitempty"`
	Info          *string `json:"info,omitempty"`
	Message       *string `json:"message,omitempty"`
}
