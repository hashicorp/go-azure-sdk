package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipOutlierInsightOutlierMemberType string

const (
	MembershipOutlierInsightOutlierMemberType_User MembershipOutlierInsightOutlierMemberType = "user"
)

func PossibleValuesForMembershipOutlierInsightOutlierMemberType() []string {
	return []string{
		string(MembershipOutlierInsightOutlierMemberType_User),
	}
}

func (s *MembershipOutlierInsightOutlierMemberType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMembershipOutlierInsightOutlierMemberType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMembershipOutlierInsightOutlierMemberType(input string) (*MembershipOutlierInsightOutlierMemberType, error) {
	vals := map[string]MembershipOutlierInsightOutlierMemberType{
		"user": MembershipOutlierInsightOutlierMemberType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MembershipOutlierInsightOutlierMemberType(input)
	return &out, nil
}
