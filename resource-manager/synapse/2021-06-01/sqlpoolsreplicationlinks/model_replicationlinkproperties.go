package sqlpoolsreplicationlinks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationLinkProperties struct {
	IsTerminationAllowed *bool             `json:"isTerminationAllowed,omitempty"`
	PartnerDatabase      *string           `json:"partnerDatabase,omitempty"`
	PartnerLocation      *string           `json:"partnerLocation,omitempty"`
	PartnerRole          *ReplicationRole  `json:"partnerRole,omitempty"`
	PartnerServer        *string           `json:"partnerServer,omitempty"`
	PercentComplete      *int64            `json:"percentComplete,omitempty"`
	ReplicationMode      *string           `json:"replicationMode,omitempty"`
	ReplicationState     *ReplicationState `json:"replicationState,omitempty"`
	Role                 *ReplicationRole  `json:"role,omitempty"`
	StartTime            *string           `json:"startTime,omitempty"`
}
