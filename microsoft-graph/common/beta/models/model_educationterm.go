package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationTerm struct {
	DisplayName *string `json:"displayName,omitempty"`
	EndDate     *string `json:"endDate,omitempty"`
	ExternalId  *string `json:"externalId,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	StartDate   *string `json:"startDate,omitempty"`
}
