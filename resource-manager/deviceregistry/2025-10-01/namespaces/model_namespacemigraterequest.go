package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceMigrateRequest struct {
	ResourceIds *[]string `json:"resourceIds,omitempty"`
	Scope       *Scope    `json:"scope,omitempty"`
}
