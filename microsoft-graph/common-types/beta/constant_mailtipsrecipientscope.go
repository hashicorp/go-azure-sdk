package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTipsRecipientScope string

const (
	MailTipsRecipientScope_External           MailTipsRecipientScope = "external"
	MailTipsRecipientScope_ExternalNonPartner MailTipsRecipientScope = "externalNonPartner"
	MailTipsRecipientScope_ExternalPartner    MailTipsRecipientScope = "externalPartner"
	MailTipsRecipientScope_Internal           MailTipsRecipientScope = "internal"
	MailTipsRecipientScope_None               MailTipsRecipientScope = "none"
)

func PossibleValuesForMailTipsRecipientScope() []string {
	return []string{
		string(MailTipsRecipientScope_External),
		string(MailTipsRecipientScope_ExternalNonPartner),
		string(MailTipsRecipientScope_ExternalPartner),
		string(MailTipsRecipientScope_Internal),
		string(MailTipsRecipientScope_None),
	}
}

func (s *MailTipsRecipientScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMailTipsRecipientScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMailTipsRecipientScope(input string) (*MailTipsRecipientScope, error) {
	vals := map[string]MailTipsRecipientScope{
		"external":           MailTipsRecipientScope_External,
		"externalnonpartner": MailTipsRecipientScope_ExternalNonPartner,
		"externalpartner":    MailTipsRecipientScope_ExternalPartner,
		"internal":           MailTipsRecipientScope_Internal,
		"none":               MailTipsRecipientScope_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MailTipsRecipientScope(input)
	return &out, nil
}
