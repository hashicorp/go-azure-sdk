package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired string

const (
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Block AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "block"
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Warn  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "warn"
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Wipe  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Block),
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Warn),
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired(input string) (*AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired{
		"block": AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired(input)
	return &out, nil
}
