package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidence struct {
	ClusterBy                *string                                       `json:"clusterBy,omitempty"`
	ClusterByValue           *string                                       `json:"clusterByValue,omitempty"`
	CreatedDateTime          *string                                       `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                     `json:"detailedRoles,omitempty"`
	EmailCount               *int64                                        `json:"emailCount,omitempty"`
	NetworkMessageIds        *[]string                                     `json:"networkMessageIds,omitempty"`
	ODataType                *string                                       `json:"@odata.type,omitempty"`
	Query                    *string                                       `json:"query,omitempty"`
	RemediationStatus        *SecurityMailClusterEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                       `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityMailClusterEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                     `json:"tags,omitempty"`
	Urn                      *string                                       `json:"urn,omitempty"`
	Verdict                  *SecurityMailClusterEvidenceVerdict           `json:"verdict,omitempty"`
}
