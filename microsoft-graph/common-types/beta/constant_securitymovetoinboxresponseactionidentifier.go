package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMoveToInboxResponseActionIdentifier string

const (
	SecurityMoveToInboxResponseActionIdentifier_NetworkMessageId      SecurityMoveToInboxResponseActionIdentifier = "networkMessageId"
	SecurityMoveToInboxResponseActionIdentifier_RecipientEmailAddress SecurityMoveToInboxResponseActionIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecurityMoveToInboxResponseActionIdentifier() []string {
	return []string{
		string(SecurityMoveToInboxResponseActionIdentifier_NetworkMessageId),
		string(SecurityMoveToInboxResponseActionIdentifier_RecipientEmailAddress),
	}
}

func (s *SecurityMoveToInboxResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMoveToInboxResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMoveToInboxResponseActionIdentifier(input string) (*SecurityMoveToInboxResponseActionIdentifier, error) {
	vals := map[string]SecurityMoveToInboxResponseActionIdentifier{
		"networkmessageid":      SecurityMoveToInboxResponseActionIdentifier_NetworkMessageId,
		"recipientemailaddress": SecurityMoveToInboxResponseActionIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMoveToInboxResponseActionIdentifier(input)
	return &out, nil
}
