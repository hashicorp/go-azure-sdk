package blobcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagProperty struct {
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`
	Tag              *string `json:"tag,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
	Upn              *string `json:"upn,omitempty"`
}
