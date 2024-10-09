package sqlpoolsblobauditing

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolBlobAuditingPolicyProperties struct {
	AuditActionsAndGroups        *[]string               `json:"auditActionsAndGroups,omitempty"`
	IsAzureMonitorTargetEnabled  *bool                   `json:"isAzureMonitorTargetEnabled,omitempty"`
	IsStorageSecondaryKeyInUse   *bool                   `json:"isStorageSecondaryKeyInUse,omitempty"`
	RetentionDays                *int64                  `json:"retentionDays,omitempty"`
	State                        BlobAuditingPolicyState `json:"state"`
	StorageAccountAccessKey      *string                 `json:"storageAccountAccessKey,omitempty"`
	StorageAccountSubscriptionId *string                 `json:"storageAccountSubscriptionId,omitempty"`
	StorageEndpoint              *string                 `json:"storageEndpoint,omitempty"`
}
