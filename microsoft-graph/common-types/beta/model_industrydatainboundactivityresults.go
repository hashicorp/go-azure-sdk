package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundActivityResults struct {
	ActivityId            *string                                       `json:"activityId,omitempty"`
	DisplayName           *string                                       `json:"displayName,omitempty"`
	Errors                *int64                                        `json:"errors,omitempty"`
	Groups                *IndustryDataIndustryDataRunEntityCountMetric `json:"groups,omitempty"`
	MatchedPeopleByRole   *[]IndustryDataIndustryDataRunRoleCountMetric `json:"matchedPeopleByRole,omitempty"`
	Memberships           *IndustryDataIndustryDataRunEntityCountMetric `json:"memberships,omitempty"`
	ODataType             *string                                       `json:"@odata.type,omitempty"`
	Organizations         *IndustryDataIndustryDataRunEntityCountMetric `json:"organizations,omitempty"`
	People                *IndustryDataIndustryDataRunEntityCountMetric `json:"people,omitempty"`
	Status                *IndustryDataInboundActivityResultsStatus     `json:"status,omitempty"`
	UnmatchedPeopleByRole *[]IndustryDataIndustryDataRunRoleCountMetric `json:"unmatchedPeopleByRole,omitempty"`
	Warnings              *int64                                        `json:"warnings,omitempty"`
}
