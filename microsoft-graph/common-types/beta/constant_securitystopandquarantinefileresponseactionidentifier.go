package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityStopAndQuarantineFileResponseActionIdentifier string

const (
	SecurityStopAndQuarantineFileResponseActionIdentifier_DeviceId              SecurityStopAndQuarantineFileResponseActionIdentifier = "deviceId"
	SecurityStopAndQuarantineFileResponseActionIdentifier_InitiatingProcessSHA1 SecurityStopAndQuarantineFileResponseActionIdentifier = "initiatingProcessSHA1"
	SecurityStopAndQuarantineFileResponseActionIdentifier_Sha1                  SecurityStopAndQuarantineFileResponseActionIdentifier = "sha1"
)

func PossibleValuesForSecurityStopAndQuarantineFileResponseActionIdentifier() []string {
	return []string{
		string(SecurityStopAndQuarantineFileResponseActionIdentifier_DeviceId),
		string(SecurityStopAndQuarantineFileResponseActionIdentifier_InitiatingProcessSHA1),
		string(SecurityStopAndQuarantineFileResponseActionIdentifier_Sha1),
	}
}

func (s *SecurityStopAndQuarantineFileResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityStopAndQuarantineFileResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityStopAndQuarantineFileResponseActionIdentifier(input string) (*SecurityStopAndQuarantineFileResponseActionIdentifier, error) {
	vals := map[string]SecurityStopAndQuarantineFileResponseActionIdentifier{
		"deviceid":              SecurityStopAndQuarantineFileResponseActionIdentifier_DeviceId,
		"initiatingprocesssha1": SecurityStopAndQuarantineFileResponseActionIdentifier_InitiatingProcessSHA1,
		"sha1":                  SecurityStopAndQuarantineFileResponseActionIdentifier_Sha1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityStopAndQuarantineFileResponseActionIdentifier(input)
	return &out, nil
}
