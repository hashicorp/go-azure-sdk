package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedAuthenticationEventListenerEventType string

const (
	AppliedAuthenticationEventListenerEventTypepageRenderStart    AppliedAuthenticationEventListenerEventType = "PageRenderStart"
	AppliedAuthenticationEventListenerEventTypetokenIssuanceStart AppliedAuthenticationEventListenerEventType = "TokenIssuanceStart"
)

func PossibleValuesForAppliedAuthenticationEventListenerEventType() []string {
	return []string{
		string(AppliedAuthenticationEventListenerEventTypepageRenderStart),
		string(AppliedAuthenticationEventListenerEventTypetokenIssuanceStart),
	}
}

func parseAppliedAuthenticationEventListenerEventType(input string) (*AppliedAuthenticationEventListenerEventType, error) {
	vals := map[string]AppliedAuthenticationEventListenerEventType{
		"pagerenderstart":    AppliedAuthenticationEventListenerEventTypepageRenderStart,
		"tokenissuancestart": AppliedAuthenticationEventListenerEventTypetokenIssuanceStart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedAuthenticationEventListenerEventType(input)
	return &out, nil
}
