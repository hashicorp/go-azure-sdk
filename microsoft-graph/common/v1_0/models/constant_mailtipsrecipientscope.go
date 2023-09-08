package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTipsRecipientScope string

const (
	MailTipsRecipientScopeexternal           MailTipsRecipientScope = "External"
	MailTipsRecipientScopeexternalNonPartner MailTipsRecipientScope = "ExternalNonPartner"
	MailTipsRecipientScopeexternalPartner    MailTipsRecipientScope = "ExternalPartner"
	MailTipsRecipientScopeinternal           MailTipsRecipientScope = "Internal"
	MailTipsRecipientScopenone               MailTipsRecipientScope = "None"
)

func PossibleValuesForMailTipsRecipientScope() []string {
	return []string{
		string(MailTipsRecipientScopeexternal),
		string(MailTipsRecipientScopeexternalNonPartner),
		string(MailTipsRecipientScopeexternalPartner),
		string(MailTipsRecipientScopeinternal),
		string(MailTipsRecipientScopenone),
	}
}

func parseMailTipsRecipientScope(input string) (*MailTipsRecipientScope, error) {
	vals := map[string]MailTipsRecipientScope{
		"external":           MailTipsRecipientScopeexternal,
		"externalnonpartner": MailTipsRecipientScopeexternalNonPartner,
		"externalpartner":    MailTipsRecipientScopeexternalPartner,
		"internal":           MailTipsRecipientScopeinternal,
		"none":               MailTipsRecipientScopenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailTipsRecipientScope(input)
	return &out, nil
}
