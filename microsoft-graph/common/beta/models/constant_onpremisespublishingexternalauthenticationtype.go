package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingExternalAuthenticationType string

const (
	OnPremisesPublishingExternalAuthenticationTypeaadPreAuthentication OnPremisesPublishingExternalAuthenticationType = "AadPreAuthentication"
	OnPremisesPublishingExternalAuthenticationTypepassthru             OnPremisesPublishingExternalAuthenticationType = "Passthru"
)

func PossibleValuesForOnPremisesPublishingExternalAuthenticationType() []string {
	return []string{
		string(OnPremisesPublishingExternalAuthenticationTypeaadPreAuthentication),
		string(OnPremisesPublishingExternalAuthenticationTypepassthru),
	}
}

func parseOnPremisesPublishingExternalAuthenticationType(input string) (*OnPremisesPublishingExternalAuthenticationType, error) {
	vals := map[string]OnPremisesPublishingExternalAuthenticationType{
		"aadpreauthentication": OnPremisesPublishingExternalAuthenticationTypeaadPreAuthentication,
		"passthru":             OnPremisesPublishingExternalAuthenticationTypepassthru,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesPublishingExternalAuthenticationType(input)
	return &out, nil
}
