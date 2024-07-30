package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecuritySessionControlCloudAppSecurityType string

const (
	CloudAppSecuritySessionControlCloudAppSecurityType_BlockDownloads CloudAppSecuritySessionControlCloudAppSecurityType = "blockDownloads"
	CloudAppSecuritySessionControlCloudAppSecurityType_McasConfigured CloudAppSecuritySessionControlCloudAppSecurityType = "mcasConfigured"
	CloudAppSecuritySessionControlCloudAppSecurityType_MonitorOnly    CloudAppSecuritySessionControlCloudAppSecurityType = "monitorOnly"
)

func PossibleValuesForCloudAppSecuritySessionControlCloudAppSecurityType() []string {
	return []string{
		string(CloudAppSecuritySessionControlCloudAppSecurityType_BlockDownloads),
		string(CloudAppSecuritySessionControlCloudAppSecurityType_McasConfigured),
		string(CloudAppSecuritySessionControlCloudAppSecurityType_MonitorOnly),
	}
}

func (s *CloudAppSecuritySessionControlCloudAppSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudAppSecuritySessionControlCloudAppSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudAppSecuritySessionControlCloudAppSecurityType(input string) (*CloudAppSecuritySessionControlCloudAppSecurityType, error) {
	vals := map[string]CloudAppSecuritySessionControlCloudAppSecurityType{
		"blockdownloads": CloudAppSecuritySessionControlCloudAppSecurityType_BlockDownloads,
		"mcasconfigured": CloudAppSecuritySessionControlCloudAppSecurityType_McasConfigured,
		"monitoronly":    CloudAppSecuritySessionControlCloudAppSecurityType_MonitorOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudAppSecuritySessionControlCloudAppSecurityType(input)
	return &out, nil
}
