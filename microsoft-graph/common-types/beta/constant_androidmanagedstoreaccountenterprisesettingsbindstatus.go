package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAccountEnterpriseSettingsBindStatus string

const (
	AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Bound             AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "bound"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatus_BoundAndValidated AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "boundAndValidated"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatus_NotBound          AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "notBound"
	AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Unbinding         AndroidManagedStoreAccountEnterpriseSettingsBindStatus = "unbinding"
)

func PossibleValuesForAndroidManagedStoreAccountEnterpriseSettingsBindStatus() []string {
	return []string{
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Bound),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatus_BoundAndValidated),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatus_NotBound),
		string(AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Unbinding),
	}
}

func (s *AndroidManagedStoreAccountEnterpriseSettingsBindStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedStoreAccountEnterpriseSettingsBindStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedStoreAccountEnterpriseSettingsBindStatus(input string) (*AndroidManagedStoreAccountEnterpriseSettingsBindStatus, error) {
	vals := map[string]AndroidManagedStoreAccountEnterpriseSettingsBindStatus{
		"bound":             AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Bound,
		"boundandvalidated": AndroidManagedStoreAccountEnterpriseSettingsBindStatus_BoundAndValidated,
		"notbound":          AndroidManagedStoreAccountEnterpriseSettingsBindStatus_NotBound,
		"unbinding":         AndroidManagedStoreAccountEnterpriseSettingsBindStatus_Unbinding,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedStoreAccountEnterpriseSettingsBindStatus(input)
	return &out, nil
}
