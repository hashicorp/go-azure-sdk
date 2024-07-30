package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy string

const (
	ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Customer ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy = "customer"
	ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Google   ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy = "google"
)

func PossibleValuesForExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy() []string {
	return []string{
		string(ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Customer),
		string(ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Google),
	}
}

func (s *ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy(input string) (*ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy, error) {
	vals := map[string]ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy{
		"customer": ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Customer,
		"google":   ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy_Google,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternallyAccessibleGcpStorageBucketFindingEncryptionManagedBy(input)
	return &out, nil
}
