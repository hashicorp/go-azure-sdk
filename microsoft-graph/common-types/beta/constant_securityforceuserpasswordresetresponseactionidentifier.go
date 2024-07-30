package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityForceUserPasswordResetResponseActionIdentifier string

const (
	SecurityForceUserPasswordResetResponseActionIdentifier_AccountSid                  SecurityForceUserPasswordResetResponseActionIdentifier = "accountSid"
	SecurityForceUserPasswordResetResponseActionIdentifier_InitiatingProcessAccountSid SecurityForceUserPasswordResetResponseActionIdentifier = "initiatingProcessAccountSid"
	SecurityForceUserPasswordResetResponseActionIdentifier_OnPremSid                   SecurityForceUserPasswordResetResponseActionIdentifier = "onPremSid"
	SecurityForceUserPasswordResetResponseActionIdentifier_RequestAccountSid           SecurityForceUserPasswordResetResponseActionIdentifier = "requestAccountSid"
)

func PossibleValuesForSecurityForceUserPasswordResetResponseActionIdentifier() []string {
	return []string{
		string(SecurityForceUserPasswordResetResponseActionIdentifier_AccountSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifier_InitiatingProcessAccountSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifier_OnPremSid),
		string(SecurityForceUserPasswordResetResponseActionIdentifier_RequestAccountSid),
	}
}

func (s *SecurityForceUserPasswordResetResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityForceUserPasswordResetResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityForceUserPasswordResetResponseActionIdentifier(input string) (*SecurityForceUserPasswordResetResponseActionIdentifier, error) {
	vals := map[string]SecurityForceUserPasswordResetResponseActionIdentifier{
		"accountsid":                  SecurityForceUserPasswordResetResponseActionIdentifier_AccountSid,
		"initiatingprocessaccountsid": SecurityForceUserPasswordResetResponseActionIdentifier_InitiatingProcessAccountSid,
		"onpremsid":                   SecurityForceUserPasswordResetResponseActionIdentifier_OnPremSid,
		"requestaccountsid":           SecurityForceUserPasswordResetResponseActionIdentifier_RequestAccountSid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityForceUserPasswordResetResponseActionIdentifier(input)
	return &out, nil
}
