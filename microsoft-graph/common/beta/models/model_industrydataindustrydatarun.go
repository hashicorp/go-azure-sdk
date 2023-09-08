package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataIndustryDataRun struct {
	Activities    *[]IndustryDataIndustryDataRunActivity `json:"activities,omitempty"`
	BlockingError *PublicError                           `json:"blockingError,omitempty"`
	DisplayName   *string                                `json:"displayName,omitempty"`
	EndDateTime   *string                                `json:"endDateTime,omitempty"`
	Id            *string                                `json:"id,omitempty"`
	ODataType     *string                                `json:"@odata.type,omitempty"`
	StartDateTime *string                                `json:"startDateTime,omitempty"`
	Status        *IndustryDataIndustryDataRunStatus     `json:"status,omitempty"`
}
