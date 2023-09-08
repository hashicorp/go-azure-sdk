package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingSingleSignOnSingleSignOnMode string

const (
	OnPremisesPublishingSingleSignOnSingleSignOnModeaadHeaderBased     OnPremisesPublishingSingleSignOnSingleSignOnMode = "AadHeaderBased"
	OnPremisesPublishingSingleSignOnSingleSignOnModenone               OnPremisesPublishingSingleSignOnSingleSignOnMode = "None"
	OnPremisesPublishingSingleSignOnSingleSignOnModeoAuthToken         OnPremisesPublishingSingleSignOnSingleSignOnMode = "OAuthToken"
	OnPremisesPublishingSingleSignOnSingleSignOnModeonPremisesKerberos OnPremisesPublishingSingleSignOnSingleSignOnMode = "OnPremisesKerberos"
	OnPremisesPublishingSingleSignOnSingleSignOnModepingHeaderBased    OnPremisesPublishingSingleSignOnSingleSignOnMode = "PingHeaderBased"
	OnPremisesPublishingSingleSignOnSingleSignOnModesaml               OnPremisesPublishingSingleSignOnSingleSignOnMode = "Saml"
)

func PossibleValuesForOnPremisesPublishingSingleSignOnSingleSignOnMode() []string {
	return []string{
		string(OnPremisesPublishingSingleSignOnSingleSignOnModeaadHeaderBased),
		string(OnPremisesPublishingSingleSignOnSingleSignOnModenone),
		string(OnPremisesPublishingSingleSignOnSingleSignOnModeoAuthToken),
		string(OnPremisesPublishingSingleSignOnSingleSignOnModeonPremisesKerberos),
		string(OnPremisesPublishingSingleSignOnSingleSignOnModepingHeaderBased),
		string(OnPremisesPublishingSingleSignOnSingleSignOnModesaml),
	}
}

func parseOnPremisesPublishingSingleSignOnSingleSignOnMode(input string) (*OnPremisesPublishingSingleSignOnSingleSignOnMode, error) {
	vals := map[string]OnPremisesPublishingSingleSignOnSingleSignOnMode{
		"aadheaderbased":     OnPremisesPublishingSingleSignOnSingleSignOnModeaadHeaderBased,
		"none":               OnPremisesPublishingSingleSignOnSingleSignOnModenone,
		"oauthtoken":         OnPremisesPublishingSingleSignOnSingleSignOnModeoAuthToken,
		"onpremiseskerberos": OnPremisesPublishingSingleSignOnSingleSignOnModeonPremisesKerberos,
		"pingheaderbased":    OnPremisesPublishingSingleSignOnSingleSignOnModepingHeaderBased,
		"saml":               OnPremisesPublishingSingleSignOnSingleSignOnModesaml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesPublishingSingleSignOnSingleSignOnMode(input)
	return &out, nil
}
