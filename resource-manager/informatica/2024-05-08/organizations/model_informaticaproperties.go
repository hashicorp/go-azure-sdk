package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformaticaProperties struct {
	InformaticaRegion *string `json:"informaticaRegion,omitempty"`
	OrganizationId    *string `json:"organizationId,omitempty"`
	OrganizationName  *string `json:"organizationName,omitempty"`
	SingleSignOnURL   *string `json:"singleSignOnUrl,omitempty"`
}
