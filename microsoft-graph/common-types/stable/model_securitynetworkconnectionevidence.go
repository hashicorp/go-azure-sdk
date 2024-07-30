package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkConnectionEvidence struct {
	CreatedDateTime          *string                                             `json:"createdDateTime,omitempty"`
	DestinationAddress       *SecurityIpEvidence                                 `json:"destinationAddress,omitempty"`
	DestinationPort          *int64                                              `json:"destinationPort,omitempty"`
	DetailedRoles            *[]string                                           `json:"detailedRoles,omitempty"`
	ODataType                *string                                             `json:"@odata.type,omitempty"`
	Protocol                 *SecurityNetworkConnectionEvidenceProtocol          `json:"protocol,omitempty"`
	RemediationStatus        *SecurityNetworkConnectionEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                             `json:"remediationStatusDetails,omitempty"`
	Roles                    *SecurityNetworkConnectionEvidenceRoles             `json:"roles,omitempty"`
	SourceAddress            *SecurityIpEvidence                                 `json:"sourceAddress,omitempty"`
	SourcePort               *int64                                              `json:"sourcePort,omitempty"`
	Tags                     *[]string                                           `json:"tags,omitempty"`
	Verdict                  *SecurityNetworkConnectionEvidenceVerdict           `json:"verdict,omitempty"`
}
