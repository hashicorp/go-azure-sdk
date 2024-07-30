package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationPromptLoginBehavior string

const (
	InternalDomainFederationPromptLoginBehavior_Disabled                               InternalDomainFederationPromptLoginBehavior = "disabled"
	InternalDomainFederationPromptLoginBehavior_NativeSupport                          InternalDomainFederationPromptLoginBehavior = "nativeSupport"
	InternalDomainFederationPromptLoginBehavior_TranslateToFreshPasswordAuthentication InternalDomainFederationPromptLoginBehavior = "translateToFreshPasswordAuthentication"
)

func PossibleValuesForInternalDomainFederationPromptLoginBehavior() []string {
	return []string{
		string(InternalDomainFederationPromptLoginBehavior_Disabled),
		string(InternalDomainFederationPromptLoginBehavior_NativeSupport),
		string(InternalDomainFederationPromptLoginBehavior_TranslateToFreshPasswordAuthentication),
	}
}

func (s *InternalDomainFederationPromptLoginBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInternalDomainFederationPromptLoginBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInternalDomainFederationPromptLoginBehavior(input string) (*InternalDomainFederationPromptLoginBehavior, error) {
	vals := map[string]InternalDomainFederationPromptLoginBehavior{
		"disabled":                               InternalDomainFederationPromptLoginBehavior_Disabled,
		"nativesupport":                          InternalDomainFederationPromptLoginBehavior_NativeSupport,
		"translatetofreshpasswordauthentication": InternalDomainFederationPromptLoginBehavior_TranslateToFreshPasswordAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationPromptLoginBehavior(input)
	return &out, nil
}
