package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySoftDeleteResponseActionIdentifier string

const (
	SecuritySoftDeleteResponseActionIdentifier_NetworkMessageId      SecuritySoftDeleteResponseActionIdentifier = "networkMessageId"
	SecuritySoftDeleteResponseActionIdentifier_RecipientEmailAddress SecuritySoftDeleteResponseActionIdentifier = "recipientEmailAddress"
)

func PossibleValuesForSecuritySoftDeleteResponseActionIdentifier() []string {
	return []string{
		string(SecuritySoftDeleteResponseActionIdentifier_NetworkMessageId),
		string(SecuritySoftDeleteResponseActionIdentifier_RecipientEmailAddress),
	}
}

func (s *SecuritySoftDeleteResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySoftDeleteResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySoftDeleteResponseActionIdentifier(input string) (*SecuritySoftDeleteResponseActionIdentifier, error) {
	vals := map[string]SecuritySoftDeleteResponseActionIdentifier{
		"networkmessageid":      SecuritySoftDeleteResponseActionIdentifier_NetworkMessageId,
		"recipientemailaddress": SecuritySoftDeleteResponseActionIdentifier_RecipientEmailAddress,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySoftDeleteResponseActionIdentifier(input)
	return &out, nil
}
