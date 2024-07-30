package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificatePaddingScheme string

const (
	UserPFXCertificatePaddingScheme_None       UserPFXCertificatePaddingScheme = "none"
	UserPFXCertificatePaddingScheme_OaepSha1   UserPFXCertificatePaddingScheme = "oaepSha1"
	UserPFXCertificatePaddingScheme_OaepSha256 UserPFXCertificatePaddingScheme = "oaepSha256"
	UserPFXCertificatePaddingScheme_OaepSha384 UserPFXCertificatePaddingScheme = "oaepSha384"
	UserPFXCertificatePaddingScheme_OaepSha512 UserPFXCertificatePaddingScheme = "oaepSha512"
	UserPFXCertificatePaddingScheme_Pkcs1      UserPFXCertificatePaddingScheme = "pkcs1"
)

func PossibleValuesForUserPFXCertificatePaddingScheme() []string {
	return []string{
		string(UserPFXCertificatePaddingScheme_None),
		string(UserPFXCertificatePaddingScheme_OaepSha1),
		string(UserPFXCertificatePaddingScheme_OaepSha256),
		string(UserPFXCertificatePaddingScheme_OaepSha384),
		string(UserPFXCertificatePaddingScheme_OaepSha512),
		string(UserPFXCertificatePaddingScheme_Pkcs1),
	}
}

func (s *UserPFXCertificatePaddingScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserPFXCertificatePaddingScheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserPFXCertificatePaddingScheme(input string) (*UserPFXCertificatePaddingScheme, error) {
	vals := map[string]UserPFXCertificatePaddingScheme{
		"none":       UserPFXCertificatePaddingScheme_None,
		"oaepsha1":   UserPFXCertificatePaddingScheme_OaepSha1,
		"oaepsha256": UserPFXCertificatePaddingScheme_OaepSha256,
		"oaepsha384": UserPFXCertificatePaddingScheme_OaepSha384,
		"oaepsha512": UserPFXCertificatePaddingScheme_OaepSha512,
		"pkcs1":      UserPFXCertificatePaddingScheme_Pkcs1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPFXCertificatePaddingScheme(input)
	return &out, nil
}
