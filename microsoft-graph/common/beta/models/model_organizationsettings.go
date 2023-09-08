package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationSettings struct {
	ContactInsights                *InsightsSettings                       `json:"contactInsights,omitempty"`
	Id                             *string                                 `json:"id,omitempty"`
	ItemInsights                   *InsightsSettings                       `json:"itemInsights,omitempty"`
	MicrosoftApplicationDataAccess *MicrosoftApplicationDataAccessSettings `json:"microsoftApplicationDataAccess,omitempty"`
	ODataType                      *string                                 `json:"@odata.type,omitempty"`
	PeopleInsights                 *InsightsSettings                       `json:"peopleInsights,omitempty"`
	ProfileCardProperties          *[]ProfileCardProperty                  `json:"profileCardProperties,omitempty"`
	Pronouns                       *PronounsSettings                       `json:"pronouns,omitempty"`
}
