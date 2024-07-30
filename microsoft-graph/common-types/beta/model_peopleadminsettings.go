package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PeopleAdminSettings struct {
	Id                    *string                `json:"id,omitempty"`
	ODataType             *string                `json:"@odata.type,omitempty"`
	ProfileCardProperties *[]ProfileCardProperty `json:"profileCardProperties,omitempty"`
	Pronouns              *PronounsSettings      `json:"pronouns,omitempty"`
}
