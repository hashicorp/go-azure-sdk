package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactedResource struct {
	AddedDateTime         *string                 `json:"addedDateTime,omitempty"`
	AdditionalDetails     *[]KeyValue             `json:"additionalDetails,omitempty"`
	ApiUrl                *string                 `json:"apiUrl,omitempty"`
	DisplayName           *string                 `json:"displayName,omitempty"`
	Id                    *string                 `json:"id,omitempty"`
	LastModifiedBy        *string                 `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime  *string                 `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                 `json:"@odata.type,omitempty"`
	Owner                 *string                 `json:"owner,omitempty"`
	PortalUrl             *string                 `json:"portalUrl,omitempty"`
	PostponeUntilDateTime *string                 `json:"postponeUntilDateTime,omitempty"`
	Rank                  *int64                  `json:"rank,omitempty"`
	RecommendationId      *string                 `json:"recommendationId,omitempty"`
	ResourceType          *string                 `json:"resourceType,omitempty"`
	Status                *ImpactedResourceStatus `json:"status,omitempty"`
	SubjectId             *string                 `json:"subjectId,omitempty"`
}
