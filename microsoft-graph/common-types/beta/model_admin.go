package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Admin struct {
	AppsAndServices     *AdminAppsAndServices `json:"appsAndServices,omitempty"`
	Dynamics            *AdminDynamics        `json:"dynamics,omitempty"`
	Edge                *Edge                 `json:"edge,omitempty"`
	Forms               *AdminForms           `json:"forms,omitempty"`
	ODataType           *string               `json:"@odata.type,omitempty"`
	People              *PeopleAdminSettings  `json:"people,omitempty"`
	ReportSettings      *AdminReportSettings  `json:"reportSettings,omitempty"`
	ServiceAnnouncement *ServiceAnnouncement  `json:"serviceAnnouncement,omitempty"`
	Sharepoint          *Sharepoint           `json:"sharepoint,omitempty"`
	Todo                *AdminTodo            `json:"todo,omitempty"`
	Windows             *AdminWindows         `json:"windows,omitempty"`
}
