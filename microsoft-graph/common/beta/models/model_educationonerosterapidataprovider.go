package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationOneRosterApiDataProvider struct {
	ConnectionSettings *EducationSynchronizationConnectionSettings `json:"connectionSettings,omitempty"`
	ConnectionUrl      *string                                     `json:"connectionUrl,omitempty"`
	Customizations     *EducationSynchronizationCustomizations     `json:"customizations,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	ProviderName       *string                                     `json:"providerName,omitempty"`
	SchoolsIds         *[]string                                   `json:"schoolsIds,omitempty"`
	TermIds            *[]string                                   `json:"termIds,omitempty"`
}
