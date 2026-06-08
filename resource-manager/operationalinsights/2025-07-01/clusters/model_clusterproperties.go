package clusters

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	AssociatedWorkspaces          *[]AssociatedWorkspace         `json:"associatedWorkspaces,omitempty"`
	BillingType                   *BillingType                   `json:"billingType,omitempty"`
	CapacityReservationProperties *CapacityReservationProperties `json:"capacityReservationProperties,omitempty"`
	ClusterId                     *string                        `json:"clusterId,omitempty"`
	CreatedDate                   *string                        `json:"createdDate,omitempty"`
	IsAvailabilityZonesEnabled    *bool                          `json:"isAvailabilityZonesEnabled,omitempty"`
	IsDoubleEncryptionEnabled     *bool                          `json:"isDoubleEncryptionEnabled,omitempty"`
	KeyVaultProperties            *KeyVaultProperties            `json:"keyVaultProperties,omitempty"`
	LastModifiedDate              *string                        `json:"lastModifiedDate,omitempty"`
	ProvisioningState             *ClusterEntityStatus           `json:"provisioningState,omitempty"`
	Replication                   *ClusterReplicationProperties  `json:"replication,omitempty"`
}

func (o *ClusterProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ClusterProperties) SetCreatedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDate = &formatted
}

func (o *ClusterProperties) GetLastModifiedDateAsTime() (*time.Time, error) {
	if o.LastModifiedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ClusterProperties) SetLastModifiedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedDate = &formatted
}
