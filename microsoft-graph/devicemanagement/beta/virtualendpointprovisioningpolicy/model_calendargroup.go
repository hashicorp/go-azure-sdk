package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroup struct {
	Calendars *[]Calendar `json:"calendars,omitempty"`
	ChangeKey *string     `json:"changeKey,omitempty"`
	ClassId   *string     `json:"classId,omitempty"`
	Id        *string     `json:"id,omitempty"`
	Name      *string     `json:"name,omitempty"`
	ODataType *string     `json:"@odata.type,omitempty"`
}
