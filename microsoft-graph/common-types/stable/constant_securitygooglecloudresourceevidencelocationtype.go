package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceLocationType string

const (
	SecurityGoogleCloudResourceEvidenceLocationType_Global   SecurityGoogleCloudResourceEvidenceLocationType = "global"
	SecurityGoogleCloudResourceEvidenceLocationType_Regional SecurityGoogleCloudResourceEvidenceLocationType = "regional"
	SecurityGoogleCloudResourceEvidenceLocationType_Unknown  SecurityGoogleCloudResourceEvidenceLocationType = "unknown"
	SecurityGoogleCloudResourceEvidenceLocationType_Zonal    SecurityGoogleCloudResourceEvidenceLocationType = "zonal"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceLocationType() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceLocationType_Global),
		string(SecurityGoogleCloudResourceEvidenceLocationType_Regional),
		string(SecurityGoogleCloudResourceEvidenceLocationType_Unknown),
		string(SecurityGoogleCloudResourceEvidenceLocationType_Zonal),
	}
}

func (s *SecurityGoogleCloudResourceEvidenceLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGoogleCloudResourceEvidenceLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGoogleCloudResourceEvidenceLocationType(input string) (*SecurityGoogleCloudResourceEvidenceLocationType, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceLocationType{
		"global":   SecurityGoogleCloudResourceEvidenceLocationType_Global,
		"regional": SecurityGoogleCloudResourceEvidenceLocationType_Regional,
		"unknown":  SecurityGoogleCloudResourceEvidenceLocationType_Unknown,
		"zonal":    SecurityGoogleCloudResourceEvidenceLocationType_Zonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceLocationType(input)
	return &out, nil
}
