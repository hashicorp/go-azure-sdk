package fileshares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileShareProperties struct {
	AccessTier             *ShareAccessTier    `json:"accessTier,omitempty"`
	AccessTierChangeTime   *string             `json:"accessTierChangeTime,omitempty"`
	AccessTierStatus       *string             `json:"accessTierStatus,omitempty"`
	Deleted                *bool               `json:"deleted,omitempty"`
	DeletedTime            *string             `json:"deletedTime,omitempty"`
	EnabledProtocols       *EnabledProtocols   `json:"enabledProtocols,omitempty"`
	LastModifiedTime       *string             `json:"lastModifiedTime,omitempty"`
	LeaseDuration          *LeaseDuration      `json:"leaseDuration,omitempty"`
	LeaseState             *LeaseState         `json:"leaseState,omitempty"`
	LeaseStatus            *LeaseStatus        `json:"leaseStatus,omitempty"`
	Metadata               *map[string]string  `json:"metadata,omitempty"`
	RemainingRetentionDays *int64              `json:"remainingRetentionDays,omitempty"`
	RootSquash             *RootSquashType     `json:"rootSquash,omitempty"`
	ShareQuota             *int64              `json:"shareQuota,omitempty"`
	ShareUsageBytes        *int64              `json:"shareUsageBytes,omitempty"`
	SignedIdentifiers      *[]SignedIdentifier `json:"signedIdentifiers,omitempty"`
	SnapshotTime           *string             `json:"snapshotTime,omitempty"`
	Version                *string             `json:"version,omitempty"`
}
