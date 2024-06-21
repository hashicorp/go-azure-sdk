package blobcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateHistoryProperty struct {
	AllowProtectedAppendWrites            *bool                         `json:"allowProtectedAppendWrites,omitempty"`
	AllowProtectedAppendWritesAll         *bool                         `json:"allowProtectedAppendWritesAll,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int64                        `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	ObjectIdentifier                      *string                       `json:"objectIdentifier,omitempty"`
	TenantId                              *string                       `json:"tenantId,omitempty"`
	Timestamp                             *string                       `json:"timestamp,omitempty"`
	Update                                *ImmutabilityPolicyUpdateType `json:"update,omitempty"`
	Upn                                   *string                       `json:"upn,omitempty"`
}
