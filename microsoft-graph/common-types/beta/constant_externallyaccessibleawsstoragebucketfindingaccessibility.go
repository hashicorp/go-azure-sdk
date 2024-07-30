package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleAwsStorageBucketFindingAccessibility string

const (
	ExternallyAccessibleAwsStorageBucketFindingAccessibility_CrossAccount ExternallyAccessibleAwsStorageBucketFindingAccessibility = "crossAccount"
	ExternallyAccessibleAwsStorageBucketFindingAccessibility_Private      ExternallyAccessibleAwsStorageBucketFindingAccessibility = "private"
	ExternallyAccessibleAwsStorageBucketFindingAccessibility_Public       ExternallyAccessibleAwsStorageBucketFindingAccessibility = "public"
	ExternallyAccessibleAwsStorageBucketFindingAccessibility_Restricted   ExternallyAccessibleAwsStorageBucketFindingAccessibility = "restricted"
)

func PossibleValuesForExternallyAccessibleAwsStorageBucketFindingAccessibility() []string {
	return []string{
		string(ExternallyAccessibleAwsStorageBucketFindingAccessibility_CrossAccount),
		string(ExternallyAccessibleAwsStorageBucketFindingAccessibility_Private),
		string(ExternallyAccessibleAwsStorageBucketFindingAccessibility_Public),
		string(ExternallyAccessibleAwsStorageBucketFindingAccessibility_Restricted),
	}
}

func (s *ExternallyAccessibleAwsStorageBucketFindingAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternallyAccessibleAwsStorageBucketFindingAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternallyAccessibleAwsStorageBucketFindingAccessibility(input string) (*ExternallyAccessibleAwsStorageBucketFindingAccessibility, error) {
	vals := map[string]ExternallyAccessibleAwsStorageBucketFindingAccessibility{
		"crossaccount": ExternallyAccessibleAwsStorageBucketFindingAccessibility_CrossAccount,
		"private":      ExternallyAccessibleAwsStorageBucketFindingAccessibility_Private,
		"public":       ExternallyAccessibleAwsStorageBucketFindingAccessibility_Public,
		"restricted":   ExternallyAccessibleAwsStorageBucketFindingAccessibility_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternallyAccessibleAwsStorageBucketFindingAccessibility(input)
	return &out, nil
}
