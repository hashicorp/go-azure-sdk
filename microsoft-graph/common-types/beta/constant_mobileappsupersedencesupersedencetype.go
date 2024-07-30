package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedenceSupersedenceType string

const (
	MobileAppSupersedenceSupersedenceType_Replace MobileAppSupersedenceSupersedenceType = "replace"
	MobileAppSupersedenceSupersedenceType_Update  MobileAppSupersedenceSupersedenceType = "update"
)

func PossibleValuesForMobileAppSupersedenceSupersedenceType() []string {
	return []string{
		string(MobileAppSupersedenceSupersedenceType_Replace),
		string(MobileAppSupersedenceSupersedenceType_Update),
	}
}

func (s *MobileAppSupersedenceSupersedenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppSupersedenceSupersedenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppSupersedenceSupersedenceType(input string) (*MobileAppSupersedenceSupersedenceType, error) {
	vals := map[string]MobileAppSupersedenceSupersedenceType{
		"replace": MobileAppSupersedenceSupersedenceType_Replace,
		"update":  MobileAppSupersedenceSupersedenceType_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupersedenceSupersedenceType(input)
	return &out, nil
}
