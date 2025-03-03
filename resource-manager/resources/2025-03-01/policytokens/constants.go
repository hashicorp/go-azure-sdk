package policytokens

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalEndpointResult string

const (
	ExternalEndpointResultFailed    ExternalEndpointResult = "Failed"
	ExternalEndpointResultSucceeded ExternalEndpointResult = "Succeeded"
)

func PossibleValuesForExternalEndpointResult() []string {
	return []string{
		string(ExternalEndpointResultFailed),
		string(ExternalEndpointResultSucceeded),
	}
}

func (s *ExternalEndpointResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalEndpointResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalEndpointResult(input string) (*ExternalEndpointResult, error) {
	vals := map[string]ExternalEndpointResult{
		"failed":    ExternalEndpointResultFailed,
		"succeeded": ExternalEndpointResultSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalEndpointResult(input)
	return &out, nil
}

type PolicyTokenResult string

const (
	PolicyTokenResultFailed    PolicyTokenResult = "Failed"
	PolicyTokenResultSucceeded PolicyTokenResult = "Succeeded"
)

func PossibleValuesForPolicyTokenResult() []string {
	return []string{
		string(PolicyTokenResultFailed),
		string(PolicyTokenResultSucceeded),
	}
}

func (s *PolicyTokenResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePolicyTokenResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePolicyTokenResult(input string) (*PolicyTokenResult, error) {
	vals := map[string]PolicyTokenResult{
		"failed":    PolicyTokenResultFailed,
		"succeeded": PolicyTokenResultSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PolicyTokenResult(input)
	return &out, nil
}
