package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProviderResponseMode string

const (
	OpenIdConnectProviderResponseModeform_post OpenIdConnectProviderResponseMode = "Formpost"
	OpenIdConnectProviderResponseModequery     OpenIdConnectProviderResponseMode = "Query"
)

func PossibleValuesForOpenIdConnectProviderResponseMode() []string {
	return []string{
		string(OpenIdConnectProviderResponseModeform_post),
		string(OpenIdConnectProviderResponseModequery),
	}
}

func parseOpenIdConnectProviderResponseMode(input string) (*OpenIdConnectProviderResponseMode, error) {
	vals := map[string]OpenIdConnectProviderResponseMode{
		"formpost": OpenIdConnectProviderResponseModeform_post,
		"query":    OpenIdConnectProviderResponseModequery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OpenIdConnectProviderResponseMode(input)
	return &out, nil
}
