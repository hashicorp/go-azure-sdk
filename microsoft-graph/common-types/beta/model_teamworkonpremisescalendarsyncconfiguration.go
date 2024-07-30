package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkOnPremisesCalendarSyncConfiguration struct {
	Domain         *string `json:"domain,omitempty"`
	DomainUserName *string `json:"domainUserName,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	SmtpAddress    *string `json:"smtpAddress,omitempty"`
}
