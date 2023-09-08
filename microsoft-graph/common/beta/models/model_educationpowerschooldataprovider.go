package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationPowerSchoolDataProvider struct {
	AllowTeachersInMultipleSchools *bool                                   `json:"allowTeachersInMultipleSchools,omitempty"`
	ClientId                       *string                                 `json:"clientId,omitempty"`
	ClientSecret                   *string                                 `json:"clientSecret,omitempty"`
	ConnectionUrl                  *string                                 `json:"connectionUrl,omitempty"`
	Customizations                 *EducationSynchronizationCustomizations `json:"customizations,omitempty"`
	ODataType                      *string                                 `json:"@odata.type,omitempty"`
	SchoolYear                     *string                                 `json:"schoolYear,omitempty"`
	SchoolsIds                     *[]string                               `json:"schoolsIds,omitempty"`
}
