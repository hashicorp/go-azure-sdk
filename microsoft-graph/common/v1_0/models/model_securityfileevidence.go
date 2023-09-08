package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidence struct {
	CreatedDateTime          *string                                `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                              `json:"detailedRoles,omitempty"`
	DetectionStatus          *SecurityFileEvidenceDetectionStatus   `json:"detectionStatus,omitempty"`
	FileDetails              *SecurityFileDetails                   `json:"fileDetails,omitempty"`
	MdeDeviceId              *string                                `json:"mdeDeviceId,omitempty"`
	ODataType                *string                                `json:"@odata.type,omitempty"`
	RemediationStatus        *SecurityFileEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                `json:"remediationStatusDetails,omitempty"`
	Roles                    *[]SecurityFileEvidenceRoles           `json:"roles,omitempty"`
	Tags                     *[]string                              `json:"tags,omitempty"`
	Verdict                  *SecurityFileEvidenceVerdict           `json:"verdict,omitempty"`
}
