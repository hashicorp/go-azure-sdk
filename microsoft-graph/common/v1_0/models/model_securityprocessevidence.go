package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidence struct {
	CreatedDateTime               *string                                   `json:"createdDateTime,omitempty"`
	DetailedRoles                 *[]string                                 `json:"detailedRoles,omitempty"`
	DetectionStatus               *SecurityProcessEvidenceDetectionStatus   `json:"detectionStatus,omitempty"`
	ImageFile                     *SecurityFileDetails                      `json:"imageFile,omitempty"`
	MdeDeviceId                   *string                                   `json:"mdeDeviceId,omitempty"`
	ODataType                     *string                                   `json:"@odata.type,omitempty"`
	ParentProcessCreationDateTime *string                                   `json:"parentProcessCreationDateTime,omitempty"`
	ParentProcessId               *int64                                    `json:"parentProcessId,omitempty"`
	ParentProcessImageFile        *SecurityFileDetails                      `json:"parentProcessImageFile,omitempty"`
	ProcessCommandLine            *string                                   `json:"processCommandLine,omitempty"`
	ProcessCreationDateTime       *string                                   `json:"processCreationDateTime,omitempty"`
	ProcessId                     *int64                                    `json:"processId,omitempty"`
	RemediationStatus             *SecurityProcessEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails      *string                                   `json:"remediationStatusDetails,omitempty"`
	Roles                         *[]SecurityProcessEvidenceRoles           `json:"roles,omitempty"`
	Tags                          *[]string                                 `json:"tags,omitempty"`
	UserAccount                   *SecurityUserAccount                      `json:"userAccount,omitempty"`
	Verdict                       *SecurityProcessEvidenceVerdict           `json:"verdict,omitempty"`
}
