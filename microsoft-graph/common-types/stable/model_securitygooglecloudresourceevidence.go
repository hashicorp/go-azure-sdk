package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidence struct {
	CreatedDateTime          *string                                               `json:"createdDateTime,omitempty"`
	DetailedRoles            *[]string                                             `json:"detailedRoles,omitempty"`
	Location                 *string                                               `json:"location,omitempty"`
	LocationType             *SecurityGoogleCloudResourceEvidenceLocationType      `json:"locationType,omitempty"`
	ODataType                *string                                               `json:"@odata.type,omitempty"`
	ProjectId                *string                                               `json:"projectId,omitempty"`
	ProjectNumber            *int64                                                `json:"projectNumber,omitempty"`
	RemediationStatus        *SecurityGoogleCloudResourceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                               `json:"remediationStatusDetails,omitempty"`
	ResourceName             *string                                               `json:"resourceName,omitempty"`
	ResourceType             *string                                               `json:"resourceType,omitempty"`
	Roles                    *SecurityGoogleCloudResourceEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                             `json:"tags,omitempty"`
	Verdict                  *SecurityGoogleCloudResourceEvidenceVerdict           `json:"verdict,omitempty"`
}
