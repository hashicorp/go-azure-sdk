package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleGcpStorageBucketFindingAccessibility string

const (
	ExternallyAccessibleGcpStorageBucketFindingAccessibility_Private             ExternallyAccessibleGcpStorageBucketFindingAccessibility = "private"
	ExternallyAccessibleGcpStorageBucketFindingAccessibility_Public              ExternallyAccessibleGcpStorageBucketFindingAccessibility = "public"
	ExternallyAccessibleGcpStorageBucketFindingAccessibility_SubjectToObjectAcls ExternallyAccessibleGcpStorageBucketFindingAccessibility = "subjectToObjectAcls"
)

func PossibleValuesForExternallyAccessibleGcpStorageBucketFindingAccessibility() []string {
	return []string{
		string(ExternallyAccessibleGcpStorageBucketFindingAccessibility_Private),
		string(ExternallyAccessibleGcpStorageBucketFindingAccessibility_Public),
		string(ExternallyAccessibleGcpStorageBucketFindingAccessibility_SubjectToObjectAcls),
	}
}

func (s *ExternallyAccessibleGcpStorageBucketFindingAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternallyAccessibleGcpStorageBucketFindingAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternallyAccessibleGcpStorageBucketFindingAccessibility(input string) (*ExternallyAccessibleGcpStorageBucketFindingAccessibility, error) {
	vals := map[string]ExternallyAccessibleGcpStorageBucketFindingAccessibility{
		"private":             ExternallyAccessibleGcpStorageBucketFindingAccessibility_Private,
		"public":              ExternallyAccessibleGcpStorageBucketFindingAccessibility_Public,
		"subjecttoobjectacls": ExternallyAccessibleGcpStorageBucketFindingAccessibility_SubjectToObjectAcls,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternallyAccessibleGcpStorageBucketFindingAccessibility(input)
	return &out, nil
}
