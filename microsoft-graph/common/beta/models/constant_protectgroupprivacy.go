package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectGroupPrivacy string

const (
	ProtectGroupPrivacyprivate     ProtectGroupPrivacy = "Private"
	ProtectGroupPrivacypublic      ProtectGroupPrivacy = "Public"
	ProtectGroupPrivacyunspecified ProtectGroupPrivacy = "Unspecified"
)

func PossibleValuesForProtectGroupPrivacy() []string {
	return []string{
		string(ProtectGroupPrivacyprivate),
		string(ProtectGroupPrivacypublic),
		string(ProtectGroupPrivacyunspecified),
	}
}

func parseProtectGroupPrivacy(input string) (*ProtectGroupPrivacy, error) {
	vals := map[string]ProtectGroupPrivacy{
		"private":     ProtectGroupPrivacyprivate,
		"public":      ProtectGroupPrivacypublic,
		"unspecified": ProtectGroupPrivacyunspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectGroupPrivacy(input)
	return &out, nil
}
