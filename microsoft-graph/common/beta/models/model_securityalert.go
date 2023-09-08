package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlert struct {
	ActorDisplayName      *string                       `json:"actorDisplayName,omitempty"`
	AdditionalData        *SecurityDictionary           `json:"additionalData,omitempty"`
	AlertPolicyId         *string                       `json:"alertPolicyId,omitempty"`
	AlertWebUrl           *string                       `json:"alertWebUrl,omitempty"`
	AssignedTo            *string                       `json:"assignedTo,omitempty"`
	Category              *string                       `json:"category,omitempty"`
	Classification        *SecurityAlertClassification  `json:"classification,omitempty"`
	Comments              *[]SecurityAlertComment       `json:"comments,omitempty"`
	CreatedDateTime       *string                       `json:"createdDateTime,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	DetectionSource       *SecurityAlertDetectionSource `json:"detectionSource,omitempty"`
	DetectorId            *string                       `json:"detectorId,omitempty"`
	Determination         *SecurityAlertDetermination   `json:"determination,omitempty"`
	Evidence              *[]SecurityAlertEvidence      `json:"evidence,omitempty"`
	FirstActivityDateTime *string                       `json:"firstActivityDateTime,omitempty"`
	Id                    *string                       `json:"id,omitempty"`
	IncidentId            *string                       `json:"incidentId,omitempty"`
	IncidentWebUrl        *string                       `json:"incidentWebUrl,omitempty"`
	LastActivityDateTime  *string                       `json:"lastActivityDateTime,omitempty"`
	LastUpdateDateTime    *string                       `json:"lastUpdateDateTime,omitempty"`
	MitreTechniques       *[]string                     `json:"mitreTechniques,omitempty"`
	ODataType             *string                       `json:"@odata.type,omitempty"`
	ProviderAlertId       *string                       `json:"providerAlertId,omitempty"`
	RecommendedActions    *string                       `json:"recommendedActions,omitempty"`
	ResolvedDateTime      *string                       `json:"resolvedDateTime,omitempty"`
	ServiceSource         *SecurityAlertServiceSource   `json:"serviceSource,omitempty"`
	Severity              *SecurityAlertSeverity        `json:"severity,omitempty"`
	Status                *SecurityAlertStatus          `json:"status,omitempty"`
	SystemTags            *[]string                     `json:"systemTags,omitempty"`
	TenantId              *string                       `json:"tenantId,omitempty"`
	ThreatDisplayName     *string                       `json:"threatDisplayName,omitempty"`
	ThreatFamilyName      *string                       `json:"threatFamilyName,omitempty"`
	Title                 *string                       `json:"title,omitempty"`
}
