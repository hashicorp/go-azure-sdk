package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingSingleSignOnSingleSignOnMode string

const (
	OnPremisesPublishingSingleSignOnSingleSignOnMode_AadHeaderBased     OnPremisesPublishingSingleSignOnSingleSignOnMode = "aadHeaderBased"
	OnPremisesPublishingSingleSignOnSingleSignOnMode_None               OnPremisesPublishingSingleSignOnSingleSignOnMode = "none"
	OnPremisesPublishingSingleSignOnSingleSignOnMode_OAuthToken         OnPremisesPublishingSingleSignOnSingleSignOnMode = "oAuthToken"
	OnPremisesPublishingSingleSignOnSingleSignOnMode_OnPremisesKerberos OnPremisesPublishingSingleSignOnSingleSignOnMode = "onPremisesKerberos"
	OnPremisesPublishingSingleSignOnSingleSignOnMode_PingHeaderBased    OnPremisesPublishingSingleSignOnSingleSignOnMode = "pingHeaderBased"
	OnPremisesPublishingSingleSignOnSingleSignOnMode_Saml               OnPremisesPublishingSingleSignOnSingleSignOnMode = "saml"
)

func PossibleValuesForOnPremisesPublishingSingleSignOnSingleSignOnMode() []string {
	return []string{
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_AadHeaderBased),
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_None),
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_OAuthToken),
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_OnPremisesKerberos),
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_PingHeaderBased),
		string(OnPremisesPublishingSingleSignOnSingleSignOnMode_Saml),
	}
}

func (s *OnPremisesPublishingSingleSignOnSingleSignOnMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnPremisesPublishingSingleSignOnSingleSignOnMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnPremisesPublishingSingleSignOnSingleSignOnMode(input string) (*OnPremisesPublishingSingleSignOnSingleSignOnMode, error) {
	vals := map[string]OnPremisesPublishingSingleSignOnSingleSignOnMode{
		"aadheaderbased":     OnPremisesPublishingSingleSignOnSingleSignOnMode_AadHeaderBased,
		"none":               OnPremisesPublishingSingleSignOnSingleSignOnMode_None,
		"oauthtoken":         OnPremisesPublishingSingleSignOnSingleSignOnMode_OAuthToken,
		"onpremiseskerberos": OnPremisesPublishingSingleSignOnSingleSignOnMode_OnPremisesKerberos,
		"pingheaderbased":    OnPremisesPublishingSingleSignOnSingleSignOnMode_PingHeaderBased,
		"saml":               OnPremisesPublishingSingleSignOnSingleSignOnMode_Saml,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesPublishingSingleSignOnSingleSignOnMode(input)
	return &out, nil
}
