package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedActionAction string

const (
	CompanyPortalBlockedActionAction_Remove  CompanyPortalBlockedActionAction = "remove"
	CompanyPortalBlockedActionAction_Reset   CompanyPortalBlockedActionAction = "reset"
	CompanyPortalBlockedActionAction_Unknown CompanyPortalBlockedActionAction = "unknown"
)

func PossibleValuesForCompanyPortalBlockedActionAction() []string {
	return []string{
		string(CompanyPortalBlockedActionAction_Remove),
		string(CompanyPortalBlockedActionAction_Reset),
		string(CompanyPortalBlockedActionAction_Unknown),
	}
}

func (s *CompanyPortalBlockedActionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompanyPortalBlockedActionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompanyPortalBlockedActionAction(input string) (*CompanyPortalBlockedActionAction, error) {
	vals := map[string]CompanyPortalBlockedActionAction{
		"remove":  CompanyPortalBlockedActionAction_Remove,
		"reset":   CompanyPortalBlockedActionAction_Reset,
		"unknown": CompanyPortalBlockedActionAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalBlockedActionAction(input)
	return &out, nil
}
