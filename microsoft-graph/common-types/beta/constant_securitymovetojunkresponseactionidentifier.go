package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMoveToJunkResponseActionIdentifier string

const (
	SecurityMoveToJunkResponseActionIdentifier_NetworkMessageId      SecurityMoveToJunkResponseActionIdentifier = "networkMessageId"
	SecurityMoveToJunkResponseActionIdentifier_RecipientEmailAddress SecurityMoveToJunkResponseActionIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecurityMoveToJunkResponseActionIdentifier() []string {
	return []string{
		string(SecurityMoveToJunkResponseActionIdentifier_NetworkMessageId),
		string(SecurityMoveToJunkResponseActionIdentifier_RecipientEmailAddress),
	}
}

func (s *SecurityMoveToJunkResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMoveToJunkResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMoveToJunkResponseActionIdentifier(input string) (*SecurityMoveToJunkResponseActionIdentifier, error) {
	vals := map[string]SecurityMoveToJunkResponseActionIdentifier{
		"networkmessageid":      SecurityMoveToJunkResponseActionIdentifier_NetworkMessageId,
		"recipientemailaddress": SecurityMoveToJunkResponseActionIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMoveToJunkResponseActionIdentifier(input)
	return &out, nil
}
