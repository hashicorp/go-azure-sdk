package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOsVersion string

const (
	CloudPCOsVersion_Windows10 CloudPCOsVersion = "windows10"
	CloudPCOsVersion_Windows11 CloudPCOsVersion = "windows11"
)

func PossibleValuesForCloudPCOsVersion() []string {
	return []string{
		string(CloudPCOsVersion_Windows10),
		string(CloudPCOsVersion_Windows11),
	}
}

func (s *CloudPCOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOsVersion(input string) (*CloudPCOsVersion, error) {
	vals := map[string]CloudPCOsVersion{
		"windows10": CloudPCOsVersion_Windows10,
		"windows11": CloudPCOsVersion_Windows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOsVersion(input)
	return &out, nil
}
