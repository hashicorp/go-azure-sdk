package integrationaccountpartners

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountPartnerProperties struct {
	ChangedTime *string        `json:"changedTime,omitempty"`
	Content     PartnerContent `json:"content"`
	CreatedTime *string        `json:"createdTime,omitempty"`
	Metadata    *interface{}   `json:"metadata,omitempty"`
	PartnerType PartnerType    `json:"partnerType"`
}
