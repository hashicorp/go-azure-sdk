package marketplaceregistrationdefinitions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiFactorAuthProvider string

const (
	MultiFactorAuthProviderAzure MultiFactorAuthProvider = "Azure"
	MultiFactorAuthProviderNone  MultiFactorAuthProvider = "None"
)

func PossibleValuesForMultiFactorAuthProvider() []string {
	return []string{
		string(MultiFactorAuthProviderAzure),
		string(MultiFactorAuthProviderNone),
	}
}

func (s *MultiFactorAuthProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiFactorAuthProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiFactorAuthProvider(input string) (*MultiFactorAuthProvider, error) {
	vals := map[string]MultiFactorAuthProvider{
		"azure": MultiFactorAuthProviderAzure,
		"none":  MultiFactorAuthProviderNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiFactorAuthProvider(input)
	return &out, nil
}
