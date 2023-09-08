package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthIssue struct {
	Classification       *ServiceHealthIssueClassification `json:"classification,omitempty"`
	Details              *[]KeyValuePair                   `json:"details,omitempty"`
	EndDateTime          *string                           `json:"endDateTime,omitempty"`
	Feature              *string                           `json:"feature,omitempty"`
	FeatureGroup         *string                           `json:"featureGroup,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	ImpactDescription    *string                           `json:"impactDescription,omitempty"`
	IsResolved           *bool                             `json:"isResolved,omitempty"`
	LastModifiedDateTime *string                           `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                           `json:"@odata.type,omitempty"`
	Origin               *ServiceHealthIssueOrigin         `json:"origin,omitempty"`
	Posts                *[]ServiceHealthIssuePost         `json:"posts,omitempty"`
	Service              *string                           `json:"service,omitempty"`
	StartDateTime        *string                           `json:"startDateTime,omitempty"`
	Status               *ServiceHealthIssueStatus         `json:"status,omitempty"`
	Title                *string                           `json:"title,omitempty"`
}
