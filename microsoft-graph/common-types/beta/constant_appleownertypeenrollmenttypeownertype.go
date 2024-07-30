package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleOwnerTypeEnrollmentTypeOwnerType string

const (
	AppleOwnerTypeEnrollmentTypeOwnerType_Company  AppleOwnerTypeEnrollmentTypeOwnerType = "company"
	AppleOwnerTypeEnrollmentTypeOwnerType_Personal AppleOwnerTypeEnrollmentTypeOwnerType = "personal"
	AppleOwnerTypeEnrollmentTypeOwnerType_Unknown  AppleOwnerTypeEnrollmentTypeOwnerType = "unknown"
)

func PossibleValuesForAppleOwnerTypeEnrollmentTypeOwnerType() []string {
	return []string{
		string(AppleOwnerTypeEnrollmentTypeOwnerType_Company),
		string(AppleOwnerTypeEnrollmentTypeOwnerType_Personal),
		string(AppleOwnerTypeEnrollmentTypeOwnerType_Unknown),
	}
}

func (s *AppleOwnerTypeEnrollmentTypeOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleOwnerTypeEnrollmentTypeOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleOwnerTypeEnrollmentTypeOwnerType(input string) (*AppleOwnerTypeEnrollmentTypeOwnerType, error) {
	vals := map[string]AppleOwnerTypeEnrollmentTypeOwnerType{
		"company":  AppleOwnerTypeEnrollmentTypeOwnerType_Company,
		"personal": AppleOwnerTypeEnrollmentTypeOwnerType_Personal,
		"unknown":  AppleOwnerTypeEnrollmentTypeOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleOwnerTypeEnrollmentTypeOwnerType(input)
	return &out, nil
}
