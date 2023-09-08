package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired string

const (
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredblock AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "Block"
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwarn  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "Warn"
	AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwipe  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredblock),
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwarn),
		string(AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired(input string) (*AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired{
		"block": AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredblock,
		"warn":  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequiredwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfSamsungKnoxAttestationRequired(input)
	return &out, nil
}
