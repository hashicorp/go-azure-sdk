package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkSettingsBindStatus string

const (
	AndroidForWorkSettingsBindStatus_Bound             AndroidForWorkSettingsBindStatus = "bound"
	AndroidForWorkSettingsBindStatus_BoundAndValidated AndroidForWorkSettingsBindStatus = "boundAndValidated"
	AndroidForWorkSettingsBindStatus_NotBound          AndroidForWorkSettingsBindStatus = "notBound"
	AndroidForWorkSettingsBindStatus_Unbinding         AndroidForWorkSettingsBindStatus = "unbinding"
)

func PossibleValuesForAndroidForWorkSettingsBindStatus() []string {
	return []string{
		string(AndroidForWorkSettingsBindStatus_Bound),
		string(AndroidForWorkSettingsBindStatus_BoundAndValidated),
		string(AndroidForWorkSettingsBindStatus_NotBound),
		string(AndroidForWorkSettingsBindStatus_Unbinding),
	}
}

func (s *AndroidForWorkSettingsBindStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkSettingsBindStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkSettingsBindStatus(input string) (*AndroidForWorkSettingsBindStatus, error) {
	vals := map[string]AndroidForWorkSettingsBindStatus{
		"bound":             AndroidForWorkSettingsBindStatus_Bound,
		"boundandvalidated": AndroidForWorkSettingsBindStatus_BoundAndValidated,
		"notbound":          AndroidForWorkSettingsBindStatus_NotBound,
		"unbinding":         AndroidForWorkSettingsBindStatus_Unbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkSettingsBindStatus(input)
	return &out, nil
}
