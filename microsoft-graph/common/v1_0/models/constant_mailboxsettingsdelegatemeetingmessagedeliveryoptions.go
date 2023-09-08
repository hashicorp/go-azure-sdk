package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsDelegateMeetingMessageDeliveryOptions string

const (
	MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndInformationToPrincipal MailboxSettingsDelegateMeetingMessageDeliveryOptions = "SendToDelegateAndInformationToPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndPrincipal              MailboxSettingsDelegateMeetingMessageDeliveryOptions = "SendToDelegateAndPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateOnly                      MailboxSettingsDelegateMeetingMessageDeliveryOptions = "SendToDelegateOnly"
)

func PossibleValuesForMailboxSettingsDelegateMeetingMessageDeliveryOptions() []string {
	return []string{
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndInformationToPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateOnly),
	}
}

func parseMailboxSettingsDelegateMeetingMessageDeliveryOptions(input string) (*MailboxSettingsDelegateMeetingMessageDeliveryOptions, error) {
	vals := map[string]MailboxSettingsDelegateMeetingMessageDeliveryOptions{
		"sendtodelegateandinformationtoprincipal": MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndInformationToPrincipal,
		"sendtodelegateandprincipal":              MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateAndPrincipal,
		"sendtodelegateonly":                      MailboxSettingsDelegateMeetingMessageDeliveryOptionssendToDelegateOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsDelegateMeetingMessageDeliveryOptions(input)
	return &out, nil
}
