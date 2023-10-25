package distributedavailabilitygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DistributedAvailabilityGroupProperties struct {
	Databases                        *[]DistributedAvailabilityGroupDatabase `json:"databases,omitempty"`
	DistributedAvailabilityGroupId   *string                                 `json:"distributedAvailabilityGroupId,omitempty"`
	DistributedAvailabilityGroupName *string                                 `json:"distributedAvailabilityGroupName,omitempty"`
	InstanceAvailabilityGroupName    *string                                 `json:"instanceAvailabilityGroupName,omitempty"`
	InstanceLinkRole                 *LinkRole                               `json:"instanceLinkRole,omitempty"`
	PartnerAvailabilityGroupName     *string                                 `json:"partnerAvailabilityGroupName,omitempty"`
	PartnerEndpoint                  *string                                 `json:"partnerEndpoint,omitempty"`
	PartnerLinkRole                  *LinkRole                               `json:"partnerLinkRole,omitempty"`
	ReplicationMode                  *ReplicationModeType                    `json:"replicationMode,omitempty"`
}
