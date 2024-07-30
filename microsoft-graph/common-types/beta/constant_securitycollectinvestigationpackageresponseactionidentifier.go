package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCollectInvestigationPackageResponseActionIdentifier string

const (
	SecurityCollectInvestigationPackageResponseActionIdentifier_DeviceId SecurityCollectInvestigationPackageResponseActionIdentifier = "deviceId"
)

func PossibleValuesForSecurityCollectInvestigationPackageResponseActionIdentifier() []string {
	return []string{
		string(SecurityCollectInvestigationPackageResponseActionIdentifier_DeviceId),
	}
}

func (s *SecurityCollectInvestigationPackageResponseActionIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCollectInvestigationPackageResponseActionIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCollectInvestigationPackageResponseActionIdentifier(input string) (*SecurityCollectInvestigationPackageResponseActionIdentifier, error) {
	vals := map[string]SecurityCollectInvestigationPackageResponseActionIdentifier{
		"deviceid": SecurityCollectInvestigationPackageResponseActionIdentifier_DeviceId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCollectInvestigationPackageResponseActionIdentifier(input)
	return &out, nil
}
