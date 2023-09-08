package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederationPromptLoginBehavior string

const (
	InternalDomainFederationPromptLoginBehaviordisabled                               InternalDomainFederationPromptLoginBehavior = "Disabled"
	InternalDomainFederationPromptLoginBehaviornativeSupport                          InternalDomainFederationPromptLoginBehavior = "NativeSupport"
	InternalDomainFederationPromptLoginBehaviortranslateToFreshPasswordAuthentication InternalDomainFederationPromptLoginBehavior = "TranslateToFreshPasswordAuthentication"
)

func PossibleValuesForInternalDomainFederationPromptLoginBehavior() []string {
	return []string{
		string(InternalDomainFederationPromptLoginBehaviordisabled),
		string(InternalDomainFederationPromptLoginBehaviornativeSupport),
		string(InternalDomainFederationPromptLoginBehaviortranslateToFreshPasswordAuthentication),
	}
}

func parseInternalDomainFederationPromptLoginBehavior(input string) (*InternalDomainFederationPromptLoginBehavior, error) {
	vals := map[string]InternalDomainFederationPromptLoginBehavior{
		"disabled":                               InternalDomainFederationPromptLoginBehaviordisabled,
		"nativesupport":                          InternalDomainFederationPromptLoginBehaviornativeSupport,
		"translatetofreshpasswordauthentication": InternalDomainFederationPromptLoginBehaviortranslateToFreshPasswordAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternalDomainFederationPromptLoginBehavior(input)
	return &out, nil
}
