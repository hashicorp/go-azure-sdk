package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyDetails struct {
	Business          *string `json:"business,omitempty"`
	CompanyName       *string `json:"companyName,omitempty"`
	Country           *string `json:"country,omitempty"`
	Domain            *string `json:"domain,omitempty"`
	NumberOfEmployees *int64  `json:"numberOfEmployees,omitempty"`
	OfficeAddress     *string `json:"officeAddress,omitempty"`
}
