package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateResult struct {
	Error      *Error             `json:"error,omitempty"`
	ResourceId *string            `json:"resourceId,omitempty"`
	Result     *MigrateResultType `json:"result,omitempty"`
}
