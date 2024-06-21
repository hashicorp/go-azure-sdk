package staticsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StaticSiteCustomDomainOverviewARMResourceProperties struct {
	CreatedOn       *string             `json:"createdOn,omitempty"`
	DomainName      *string             `json:"domainName,omitempty"`
	ErrorMessage    *string             `json:"errorMessage,omitempty"`
	Status          *CustomDomainStatus `json:"status,omitempty"`
	ValidationToken *string             `json:"validationToken,omitempty"`
}
