package videos

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicyEccAlgo string

const (
	AccessPolicyEccAlgoESFiveOneTwo     AccessPolicyEccAlgo = "ES512"
	AccessPolicyEccAlgoESThreeEightFour AccessPolicyEccAlgo = "ES384"
	AccessPolicyEccAlgoESTwoFiveSix     AccessPolicyEccAlgo = "ES256"
)

func PossibleValuesForAccessPolicyEccAlgo() []string {
	return []string{
		string(AccessPolicyEccAlgoESFiveOneTwo),
		string(AccessPolicyEccAlgoESThreeEightFour),
		string(AccessPolicyEccAlgoESTwoFiveSix),
	}
}

func (s *AccessPolicyEccAlgo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPolicyEccAlgo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPolicyEccAlgo(input string) (*AccessPolicyEccAlgo, error) {
	vals := map[string]AccessPolicyEccAlgo{
		"es512": AccessPolicyEccAlgoESFiveOneTwo,
		"es384": AccessPolicyEccAlgoESThreeEightFour,
		"es256": AccessPolicyEccAlgoESTwoFiveSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPolicyEccAlgo(input)
	return &out, nil
}

type AccessPolicyRole string

const (
	AccessPolicyRoleReader AccessPolicyRole = "Reader"
)

func PossibleValuesForAccessPolicyRole() []string {
	return []string{
		string(AccessPolicyRoleReader),
	}
}

func (s *AccessPolicyRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPolicyRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPolicyRole(input string) (*AccessPolicyRole, error) {
	vals := map[string]AccessPolicyRole{
		"reader": AccessPolicyRoleReader,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPolicyRole(input)
	return &out, nil
}

type AccessPolicyRsaAlgo string

const (
	AccessPolicyRsaAlgoRSFiveOneTwo     AccessPolicyRsaAlgo = "RS512"
	AccessPolicyRsaAlgoRSThreeEightFour AccessPolicyRsaAlgo = "RS384"
	AccessPolicyRsaAlgoRSTwoFiveSix     AccessPolicyRsaAlgo = "RS256"
)

func PossibleValuesForAccessPolicyRsaAlgo() []string {
	return []string{
		string(AccessPolicyRsaAlgoRSFiveOneTwo),
		string(AccessPolicyRsaAlgoRSThreeEightFour),
		string(AccessPolicyRsaAlgoRSTwoFiveSix),
	}
}

func (s *AccessPolicyRsaAlgo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPolicyRsaAlgo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPolicyRsaAlgo(input string) (*AccessPolicyRsaAlgo, error) {
	vals := map[string]AccessPolicyRsaAlgo{
		"rs512": AccessPolicyRsaAlgoRSFiveOneTwo,
		"rs384": AccessPolicyRsaAlgoRSThreeEightFour,
		"rs256": AccessPolicyRsaAlgoRSTwoFiveSix,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPolicyRsaAlgo(input)
	return &out, nil
}

type VideoType string

const (
	VideoTypeArchive VideoType = "Archive"
)

func PossibleValuesForVideoType() []string {
	return []string{
		string(VideoTypeArchive),
	}
}

func (s *VideoType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVideoType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVideoType(input string) (*VideoType, error) {
	vals := map[string]VideoType{
		"archive": VideoTypeArchive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VideoType(input)
	return &out, nil
}
