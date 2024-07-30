package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionRequestTemplateType string

const (
	DeviceLogCollectionRequestTemplateType_Predefined DeviceLogCollectionRequestTemplateType = "predefined"
)

func PossibleValuesForDeviceLogCollectionRequestTemplateType() []string {
	return []string{
		string(DeviceLogCollectionRequestTemplateType_Predefined),
	}
}

func (s *DeviceLogCollectionRequestTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLogCollectionRequestTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLogCollectionRequestTemplateType(input string) (*DeviceLogCollectionRequestTemplateType, error) {
	vals := map[string]DeviceLogCollectionRequestTemplateType{
		"predefined": DeviceLogCollectionRequestTemplateType_Predefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionRequestTemplateType(input)
	return &out, nil
}
