package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PositionDetail struct {
	Company        *CompanyDetail `json:"company,omitempty"`
	Description    *string        `json:"description,omitempty"`
	EndMonthYear   *string        `json:"endMonthYear,omitempty"`
	JobTitle       *string        `json:"jobTitle,omitempty"`
	ODataType      *string        `json:"@odata.type,omitempty"`
	Role           *string        `json:"role,omitempty"`
	StartMonthYear *string        `json:"startMonthYear,omitempty"`
	Summary        *string        `json:"summary,omitempty"`
}
