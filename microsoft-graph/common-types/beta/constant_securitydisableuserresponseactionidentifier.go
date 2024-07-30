package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDisableUserResponseActionIdentifier string

const (
	SecurityDisableUserResponseActionIdentifier_AccountSid                  SecurityDisableUserResponseActionIdentifier = "accountSid"
	SecurityDisableUserResponseActionIdentifier_InitiatingProcessAccountSid SecurityDisableUserResponseActionIdentifier = "initiatingProcessAccountSid"
	SecurityDisableUserResponseActionIdentifier_OnPremSid                   SecurityDisableUserResponseActionIdentifier = "onPremSid"
	SecurityDisableUserResponseActionIdentifier_RequestAccountSid           SecurityDisableUserResponseActionIdentifier = "requestAccountSid"
)

func PossibleValuesForSecurityDisableUserResponseActionIdentifier() []string {
	return []string{
		string(SecurityDisableUserResponseActionIdentifier_AccountSid),
		string(SecurityDisableUserResponseActionIdentifier_InitiatingProcessAccountSid),
		string(SecurityDisableUserResponseActionIdentifier_OnPremSid),
		string(SecurityDisableUserResponseActionIdentifier_RequestAccountSid),
	}
}

func (s *SecurityDisableUserResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDisableUserResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDisableUserResponseActionIdentifier(input string) (*SecurityDisableUserResponseActionIdentifier, error) {
	vals := map[string]SecurityDisableUserResponseActionIdentifier{
		"accountsid":                  SecurityDisableUserResponseActionIdentifier_AccountSid,
		"initiatingprocessaccountsid": SecurityDisableUserResponseActionIdentifier_InitiatingProcessAccountSid,
		"onpremsid":                   SecurityDisableUserResponseActionIdentifier_OnPremSid,
		"requestaccountsid":           SecurityDisableUserResponseActionIdentifier_RequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDisableUserResponseActionIdentifier(input)
	return &out, nil
}
