package migrates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteSpnProperties struct {
	AadAuthority  *string `json:"aadAuthority,omitempty"`
	ApplicationId *string `json:"applicationId,omitempty"`
	Audience      *string `json:"audience,omitempty"`
	ObjectId      *string `json:"objectId,omitempty"`
	RawCertData   *string `json:"rawCertData,omitempty"`
	TenantId      *string `json:"tenantId,omitempty"`
}
