package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHardDeleteResponseActionIdentifier string

const (
	SecurityHardDeleteResponseActionIdentifier_NetworkMessageId      SecurityHardDeleteResponseActionIdentifier = "networkMessageId"
	SecurityHardDeleteResponseActionIdentifier_RecipientEmailAddress SecurityHardDeleteResponseActionIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecurityHardDeleteResponseActionIdentifier() []string {
	return []string{
		string(SecurityHardDeleteResponseActionIdentifier_NetworkMessageId),
		string(SecurityHardDeleteResponseActionIdentifier_RecipientEmailAddress),
	}
}

func (s *SecurityHardDeleteResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHardDeleteResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHardDeleteResponseActionIdentifier(input string) (*SecurityHardDeleteResponseActionIdentifier, error) {
	vals := map[string]SecurityHardDeleteResponseActionIdentifier{
		"networkmessageid":      SecurityHardDeleteResponseActionIdentifier_NetworkMessageId,
		"recipientemailaddress": SecurityHardDeleteResponseActionIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHardDeleteResponseActionIdentifier(input)
	return &out, nil
}
