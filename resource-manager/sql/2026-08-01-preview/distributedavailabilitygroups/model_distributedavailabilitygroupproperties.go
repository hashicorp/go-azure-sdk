package distributedavailabilitygroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributedAvailabilityGroupProperties struct {
	Databases                        *[]DistributedAvailabilityGroupDatabase `json:"databases,omitempty"`
	DistributedAvailabilityGroupId   *string                                 `json:"distributedAvailabilityGroupId,omitempty"`
	DistributedAvailabilityGroupName *string                                 `json:"distributedAvailabilityGroupName,omitempty"`
	FailoverMode                     *FailoverModeType                       `json:"failoverMode,omitempty"`
	InstanceAvailabilityGroupName    *string                                 `json:"instanceAvailabilityGroupName,omitempty"`
	InstanceLinkRole                 *LinkRole                               `json:"instanceLinkRole,omitempty"`
	LinkMode                         *LinkModeType                           `json:"linkMode,omitempty"`
	MostRecentError                  *string                                 `json:"mostRecentError,omitempty"`
	MostRecentErrorMessage           *string                                 `json:"mostRecentErrorMessage,omitempty"`
	MostRecentErrorTime              *string                                 `json:"mostRecentErrorTime,omitempty"`
	PartnerAvailabilityGroupName     *string                                 `json:"partnerAvailabilityGroupName,omitempty"`
	PartnerEndpoint                  *string                                 `json:"partnerEndpoint,omitempty"`
	PartnerLinkRole                  *LinkRole                               `json:"partnerLinkRole,omitempty"`
	ReplicationMode                  *ReplicationModeType                    `json:"replicationMode,omitempty"`
	SeedingMode                      *SeedingModeType                        `json:"seedingMode,omitempty"`
}

func (o *DistributedAvailabilityGroupProperties) GetMostRecentErrorTimeAsTime() (*time.Time, error) {
	if o.MostRecentErrorTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MostRecentErrorTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DistributedAvailabilityGroupProperties) SetMostRecentErrorTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MostRecentErrorTime = &formatted
}
