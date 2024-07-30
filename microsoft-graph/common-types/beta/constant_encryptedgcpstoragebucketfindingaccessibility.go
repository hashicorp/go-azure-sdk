package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedGcpStorageBucketFindingAccessibility string

const (
	EncryptedGcpStorageBucketFindingAccessibility_Private             EncryptedGcpStorageBucketFindingAccessibility = "private"
	EncryptedGcpStorageBucketFindingAccessibility_Public              EncryptedGcpStorageBucketFindingAccessibility = "public"
	EncryptedGcpStorageBucketFindingAccessibility_SubjectToObjectAcls EncryptedGcpStorageBucketFindingAccessibility = "subjectToObjectAcls"
)

func PossibleValuesForEncryptedGcpStorageBucketFindingAccessibility() []string {
	return []string{
		string(EncryptedGcpStorageBucketFindingAccessibility_Private),
		string(EncryptedGcpStorageBucketFindingAccessibility_Public),
		string(EncryptedGcpStorageBucketFindingAccessibility_SubjectToObjectAcls),
	}
}

func (s *EncryptedGcpStorageBucketFindingAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptedGcpStorageBucketFindingAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptedGcpStorageBucketFindingAccessibility(input string) (*EncryptedGcpStorageBucketFindingAccessibility, error) {
	vals := map[string]EncryptedGcpStorageBucketFindingAccessibility{
		"private":             EncryptedGcpStorageBucketFindingAccessibility_Private,
		"public":              EncryptedGcpStorageBucketFindingAccessibility_Public,
		"subjecttoobjectacls": EncryptedGcpStorageBucketFindingAccessibility_SubjectToObjectAcls,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptedGcpStorageBucketFindingAccessibility(input)
	return &out, nil
}
