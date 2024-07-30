package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedGcpStorageBucketFindingEncryptionManagedBy string

const (
	EncryptedGcpStorageBucketFindingEncryptionManagedBy_Customer EncryptedGcpStorageBucketFindingEncryptionManagedBy = "customer"
	EncryptedGcpStorageBucketFindingEncryptionManagedBy_Google   EncryptedGcpStorageBucketFindingEncryptionManagedBy = "google"
)

func PossibleValuesForEncryptedGcpStorageBucketFindingEncryptionManagedBy() []string {
	return []string{
		string(EncryptedGcpStorageBucketFindingEncryptionManagedBy_Customer),
		string(EncryptedGcpStorageBucketFindingEncryptionManagedBy_Google),
	}
}

func (s *EncryptedGcpStorageBucketFindingEncryptionManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptedGcpStorageBucketFindingEncryptionManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptedGcpStorageBucketFindingEncryptionManagedBy(input string) (*EncryptedGcpStorageBucketFindingEncryptionManagedBy, error) {
	vals := map[string]EncryptedGcpStorageBucketFindingEncryptionManagedBy{
		"customer": EncryptedGcpStorageBucketFindingEncryptionManagedBy_Customer,
		"google":   EncryptedGcpStorageBucketFindingEncryptionManagedBy_Google,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptedGcpStorageBucketFindingEncryptionManagedBy(input)
	return &out, nil
}
