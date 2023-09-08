package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantExecutionError struct {
	Error          *string `json:"error,omitempty"`
	ErrorDetails   *string `json:"errorDetails,omitempty"`
	NodeId         *int64  `json:"nodeId,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	RawToken       *string `json:"rawToken,omitempty"`
	StatementIndex *int64  `json:"statementIndex,omitempty"`
	TenantId       *string `json:"tenantId,omitempty"`
}
