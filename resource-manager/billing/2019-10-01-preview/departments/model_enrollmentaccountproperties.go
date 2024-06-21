package departments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentAccountProperties struct {
	AccountName  *string     `json:"accountName,omitempty"`
	AccountOwner *string     `json:"accountOwner,omitempty"`
	CostCenter   *string     `json:"costCenter,omitempty"`
	Department   *Department `json:"department,omitempty"`
	EndDate      *string     `json:"endDate,omitempty"`
	StartDate    *string     `json:"startDate,omitempty"`
	Status       *string     `json:"status,omitempty"`
}
