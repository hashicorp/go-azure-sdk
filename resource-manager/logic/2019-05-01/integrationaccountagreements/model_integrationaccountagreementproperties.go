package integrationaccountagreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountAgreementProperties struct {
	AgreementType AgreementType    `json:"agreementType"`
	ChangedTime   *string          `json:"changedTime,omitempty"`
	Content       AgreementContent `json:"content"`
	CreatedTime   *string          `json:"createdTime,omitempty"`
	GuestIdentity BusinessIdentity `json:"guestIdentity"`
	GuestPartner  string           `json:"guestPartner"`
	HostIdentity  BusinessIdentity `json:"hostIdentity"`
	HostPartner   string           `json:"hostPartner"`
	Metadata      *interface{}     `json:"metadata,omitempty"`
}
