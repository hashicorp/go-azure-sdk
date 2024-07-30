package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAggregatedInboundStatistics struct {
	Errors                *int64                                        `json:"errors,omitempty"`
	Groups                *int64                                        `json:"groups,omitempty"`
	MatchedPeopleByRole   *[]IndustryDataIndustryDataRunRoleCountMetric `json:"matchedPeopleByRole,omitempty"`
	Memberships           *int64                                        `json:"memberships,omitempty"`
	ODataType             *string                                       `json:"@odata.type,omitempty"`
	Organizations         *int64                                        `json:"organizations,omitempty"`
	People                *int64                                        `json:"people,omitempty"`
	UnmatchedPeopleByRole *[]IndustryDataIndustryDataRunRoleCountMetric `json:"unmatchedPeopleByRole,omitempty"`
	Warnings              *int64                                        `json:"warnings,omitempty"`
}
