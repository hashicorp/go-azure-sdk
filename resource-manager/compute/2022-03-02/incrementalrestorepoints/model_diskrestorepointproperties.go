package incrementalrestorepoints

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskRestorePointProperties struct {
	CompletionPercent      *float64               `json:"completionPercent,omitempty"`
	DiskAccessId           *string                `json:"diskAccessId,omitempty"`
	Encryption             *Encryption            `json:"encryption,omitempty"`
	FamilyId               *string                `json:"familyId,omitempty"`
	HyperVGeneration       *HyperVGeneration      `json:"hyperVGeneration,omitempty"`
	NetworkAccessPolicy    *NetworkAccessPolicy   `json:"networkAccessPolicy,omitempty"`
	OsType                 *OperatingSystemTypes  `json:"osType,omitempty"`
	PublicNetworkAccess    *PublicNetworkAccess   `json:"publicNetworkAccess,omitempty"`
	PurchasePlan           *PurchasePlan          `json:"purchasePlan,omitempty"`
	ReplicationState       *string                `json:"replicationState,omitempty"`
	SecurityProfile        *DiskSecurityProfile   `json:"securityProfile,omitempty"`
	SourceResourceId       *string                `json:"sourceResourceId,omitempty"`
	SourceResourceLocation *string                `json:"sourceResourceLocation,omitempty"`
	SourceUniqueId         *string                `json:"sourceUniqueId,omitempty"`
	SupportedCapabilities  *SupportedCapabilities `json:"supportedCapabilities,omitempty"`
	SupportsHibernation    *bool                  `json:"supportsHibernation,omitempty"`
	TimeCreated            *string                `json:"timeCreated,omitempty"`
}

func (o *DiskRestorePointProperties) GetTimeCreatedAsTime() (*time.Time, error) {
	if o.TimeCreated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeCreated, "2006-01-02T15:04:05Z07:00")
}
