package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMarkUserAsCompromisedResponseActionIdentifier string

const (
	SecurityMarkUserAsCompromisedResponseActionIdentifier_AccountObjectId                  SecurityMarkUserAsCompromisedResponseActionIdentifier = "accountObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifier_InitiatingProcessAccountObjectId SecurityMarkUserAsCompromisedResponseActionIdentifier = "initiatingProcessAccountObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifier_RecipientObjectId                SecurityMarkUserAsCompromisedResponseActionIdentifier = "recipientObjectId"
	SecurityMarkUserAsCompromisedResponseActionIdentifier_ServicePrincipalId               SecurityMarkUserAsCompromisedResponseActionIdentifier = "servicePrincipalId"
)

func PossibleValuesForSecurityMarkUserAsCompromisedResponseActionIdentifier() []string {
	return []string{
		string(SecurityMarkUserAsCompromisedResponseActionIdentifier_AccountObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifier_InitiatingProcessAccountObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifier_RecipientObjectId),
		string(SecurityMarkUserAsCompromisedResponseActionIdentifier_ServicePrincipalId),
	}
}

func (s *SecurityMarkUserAsCompromisedResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMarkUserAsCompromisedResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMarkUserAsCompromisedResponseActionIdentifier(input string) (*SecurityMarkUserAsCompromisedResponseActionIdentifier, error) {
	vals := map[string]SecurityMarkUserAsCompromisedResponseActionIdentifier{
		"accountobjectid":                  SecurityMarkUserAsCompromisedResponseActionIdentifier_AccountObjectId,
		"initiatingprocessaccountobjectid": SecurityMarkUserAsCompromisedResponseActionIdentifier_InitiatingProcessAccountObjectId,
		"recipientobjectid":                SecurityMarkUserAsCompromisedResponseActionIdentifier_RecipientObjectId,
		"serviceprincipalid":               SecurityMarkUserAsCompromisedResponseActionIdentifier_ServicePrincipalId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMarkUserAsCompromisedResponseActionIdentifier(input)
	return &out, nil
}
