package plan

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountCreationSource string

const (
	AccountCreationSourceLIFTR    AccountCreationSource = "LIFTR"
	AccountCreationSourceNEWRELIC AccountCreationSource = "NEWRELIC"
)

func PossibleValuesForAccountCreationSource() []string {
	return []string{
		string(AccountCreationSourceLIFTR),
		string(AccountCreationSourceNEWRELIC),
	}
}

func (s *AccountCreationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccountCreationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccountCreationSource(input string) (*AccountCreationSource, error) {
	vals := map[string]AccountCreationSource{
		"liftr":    AccountCreationSourceLIFTR,
		"newrelic": AccountCreationSourceNEWRELIC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountCreationSource(input)
	return &out, nil
}

type OrgCreationSource string

const (
	OrgCreationSourceLIFTR    OrgCreationSource = "LIFTR"
	OrgCreationSourceNEWRELIC OrgCreationSource = "NEWRELIC"
)

func PossibleValuesForOrgCreationSource() []string {
	return []string{
		string(OrgCreationSourceLIFTR),
		string(OrgCreationSourceNEWRELIC),
	}
}

func (s *OrgCreationSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrgCreationSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrgCreationSource(input string) (*OrgCreationSource, error) {
	vals := map[string]OrgCreationSource{
		"liftr":    OrgCreationSourceLIFTR,
		"newrelic": OrgCreationSourceNEWRELIC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrgCreationSource(input)
	return &out, nil
}

type UsageType string

const (
	UsageTypeCOMMITTED UsageType = "COMMITTED"
	UsageTypePAYG      UsageType = "PAYG"
)

func PossibleValuesForUsageType() []string {
	return []string{
		string(UsageTypeCOMMITTED),
		string(UsageTypePAYG),
	}
}

func (s *UsageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageType(input string) (*UsageType, error) {
	vals := map[string]UsageType{
		"committed": UsageTypeCOMMITTED,
		"payg":      UsageTypePAYG,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageType(input)
	return &out, nil
}
