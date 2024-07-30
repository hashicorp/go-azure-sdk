package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatusUserAccessLevel string

const (
	CloudPCReviewStatusUserAccessLevel_Restricted   CloudPCReviewStatusUserAccessLevel = "restricted"
	CloudPCReviewStatusUserAccessLevel_Unrestricted CloudPCReviewStatusUserAccessLevel = "unrestricted"
)

func PossibleValuesForCloudPCReviewStatusUserAccessLevel() []string {
	return []string{
		string(CloudPCReviewStatusUserAccessLevel_Restricted),
		string(CloudPCReviewStatusUserAccessLevel_Unrestricted),
	}
}

func (s *CloudPCReviewStatusUserAccessLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCReviewStatusUserAccessLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCReviewStatusUserAccessLevel(input string) (*CloudPCReviewStatusUserAccessLevel, error) {
	vals := map[string]CloudPCReviewStatusUserAccessLevel{
		"restricted":   CloudPCReviewStatusUserAccessLevel_Restricted,
		"unrestricted": CloudPCReviewStatusUserAccessLevel_Unrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCReviewStatusUserAccessLevel(input)
	return &out, nil
}
