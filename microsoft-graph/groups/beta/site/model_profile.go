package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Profile struct {
	Account               *[]UserAccountInformation `json:"account,omitempty"`
	Addresses             *[]ItemAddress            `json:"addresses,omitempty"`
	Anniversaries         *[]PersonAnnualEvent      `json:"anniversaries,omitempty"`
	Awards                *[]PersonAward            `json:"awards,omitempty"`
	Certifications        *[]PersonCertification    `json:"certifications,omitempty"`
	EducationalActivities *[]EducationalActivity    `json:"educationalActivities,omitempty"`
	Emails                *[]ItemEmail              `json:"emails,omitempty"`
	Id                    *string                   `json:"id,omitempty"`
	Interests             *[]PersonInterest         `json:"interests,omitempty"`
	Languages             *[]LanguageProficiency    `json:"languages,omitempty"`
	Names                 *[]PersonName             `json:"names,omitempty"`
	Notes                 *[]PersonAnnotation       `json:"notes,omitempty"`
	ODataType             *string                   `json:"@odata.type,omitempty"`
	Patents               *[]ItemPatent             `json:"patents,omitempty"`
	Phones                *[]ItemPhone              `json:"phones,omitempty"`
	Positions             *[]WorkPosition           `json:"positions,omitempty"`
	Projects              *[]ProjectParticipation   `json:"projects,omitempty"`
	Publications          *[]ItemPublication        `json:"publications,omitempty"`
	Skills                *[]SkillProficiency       `json:"skills,omitempty"`
	WebAccounts           *[]WebAccount             `json:"webAccounts,omitempty"`
	Websites              *[]PersonWebsite          `json:"websites,omitempty"`
}
