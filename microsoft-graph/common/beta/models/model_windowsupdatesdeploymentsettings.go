package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesDeploymentSettings struct {
	ContentApplicability *WindowsUpdatesContentApplicabilitySettings `json:"contentApplicability,omitempty"`
	Expedite             *WindowsUpdatesExpediteSettings             `json:"expedite,omitempty"`
	Monitoring           *WindowsUpdatesMonitoringSettings           `json:"monitoring,omitempty"`
	ODataType            *string                                     `json:"@odata.type,omitempty"`
	Schedule             *WindowsUpdatesScheduleSettings             `json:"schedule,omitempty"`
	UserExperience       *WindowsUpdatesUserExperienceSettings       `json:"userExperience,omitempty"`
}
