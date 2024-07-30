package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleAzureBlobContainerFindingAccessibility string

const (
	ExternallyAccessibleAzureBlobContainerFindingAccessibility_Private ExternallyAccessibleAzureBlobContainerFindingAccessibility = "private"
	ExternallyAccessibleAzureBlobContainerFindingAccessibility_Public  ExternallyAccessibleAzureBlobContainerFindingAccessibility = "public"
)

func PossibleValuesForExternallyAccessibleAzureBlobContainerFindingAccessibility() []string {
	return []string{
		string(ExternallyAccessibleAzureBlobContainerFindingAccessibility_Private),
		string(ExternallyAccessibleAzureBlobContainerFindingAccessibility_Public),
	}
}

func (s *ExternallyAccessibleAzureBlobContainerFindingAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternallyAccessibleAzureBlobContainerFindingAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternallyAccessibleAzureBlobContainerFindingAccessibility(input string) (*ExternallyAccessibleAzureBlobContainerFindingAccessibility, error) {
	vals := map[string]ExternallyAccessibleAzureBlobContainerFindingAccessibility{
		"private": ExternallyAccessibleAzureBlobContainerFindingAccessibility_Private,
		"public":  ExternallyAccessibleAzureBlobContainerFindingAccessibility_Public,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternallyAccessibleAzureBlobContainerFindingAccessibility(input)
	return &out, nil
}
