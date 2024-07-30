package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailboxSettingsDelegateMeetingMessageDeliveryOptions string

const (
	MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndInformationToPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal              MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateAndPrincipal"
	MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateOnly                      MailboxSettingsDelegateMeetingMessageDeliveryOptions = "sendToDelegateOnly"
)

func PossibleValuesForMailboxSettingsDelegateMeetingMessageDeliveryOptions() []string {
	return []string{
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal),
		string(MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateOnly),
	}
}

func (s *MailboxSettingsDelegateMeetingMessageDeliveryOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailboxSettingsDelegateMeetingMessageDeliveryOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailboxSettingsDelegateMeetingMessageDeliveryOptions(input string) (*MailboxSettingsDelegateMeetingMessageDeliveryOptions, error) {
	vals := map[string]MailboxSettingsDelegateMeetingMessageDeliveryOptions{
		"sendtodelegateandinformationtoprincipal": MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndInformationToPrincipal,
		"sendtodelegateandprincipal":              MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateAndPrincipal,
		"sendtodelegateonly":                      MailboxSettingsDelegateMeetingMessageDeliveryOptions_SendToDelegateOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailboxSettingsDelegateMeetingMessageDeliveryOptions(input)
	return &out, nil
}
