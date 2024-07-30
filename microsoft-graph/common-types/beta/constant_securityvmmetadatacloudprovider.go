package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityVmMetadataCloudProvider string

const (
	SecurityVmMetadataCloudProvider_Azure   SecurityVmMetadataCloudProvider = "azure"
	SecurityVmMetadataCloudProvider_Unknown SecurityVmMetadataCloudProvider = "unknown"
)

func PossibleValuesForSecurityVmMetadataCloudProvider() []string {
	return []string{
		string(SecurityVmMetadataCloudProvider_Azure),
		string(SecurityVmMetadataCloudProvider_Unknown),
	}
}

func (s *SecurityVmMetadataCloudProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityVmMetadataCloudProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityVmMetadataCloudProvider(input string) (*SecurityVmMetadataCloudProvider, error) {
	vals := map[string]SecurityVmMetadataCloudProvider{
		"azure":   SecurityVmMetadataCloudProvider_Azure,
		"unknown": SecurityVmMetadataCloudProvider_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityVmMetadataCloudProvider(input)
	return &out, nil
}
