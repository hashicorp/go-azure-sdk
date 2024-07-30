package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedAwsStorageBucketFindingAccessibility string

const (
	EncryptedAwsStorageBucketFindingAccessibility_CrossAccount EncryptedAwsStorageBucketFindingAccessibility = "crossAccount"
	EncryptedAwsStorageBucketFindingAccessibility_Private      EncryptedAwsStorageBucketFindingAccessibility = "private"
	EncryptedAwsStorageBucketFindingAccessibility_Public       EncryptedAwsStorageBucketFindingAccessibility = "public"
	EncryptedAwsStorageBucketFindingAccessibility_Restricted   EncryptedAwsStorageBucketFindingAccessibility = "restricted"
)

func PossibleValuesForEncryptedAwsStorageBucketFindingAccessibility() []string {
	return []string{
		string(EncryptedAwsStorageBucketFindingAccessibility_CrossAccount),
		string(EncryptedAwsStorageBucketFindingAccessibility_Private),
		string(EncryptedAwsStorageBucketFindingAccessibility_Public),
		string(EncryptedAwsStorageBucketFindingAccessibility_Restricted),
	}
}

func (s *EncryptedAwsStorageBucketFindingAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptedAwsStorageBucketFindingAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptedAwsStorageBucketFindingAccessibility(input string) (*EncryptedAwsStorageBucketFindingAccessibility, error) {
	vals := map[string]EncryptedAwsStorageBucketFindingAccessibility{
		"crossaccount": EncryptedAwsStorageBucketFindingAccessibility_CrossAccount,
		"private":      EncryptedAwsStorageBucketFindingAccessibility_Private,
		"public":       EncryptedAwsStorageBucketFindingAccessibility_Public,
		"restricted":   EncryptedAwsStorageBucketFindingAccessibility_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptedAwsStorageBucketFindingAccessibility(input)
	return &out, nil
}
