package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_AllFeatures             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "allFeatures"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Biometrics              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "biometrics"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Camera                  AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "camera"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Face                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "face"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Fingerprint             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "fingerprint"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Iris                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "iris"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_NotConfigured           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Notifications           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "notifications"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_RemoteInput             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "remoteInput"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_TrustAgents             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "trustAgents"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_UnredactedNotifications AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "unredactedNotifications"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_AllFeatures),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Biometrics),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Camera),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Face),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Fingerprint),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Iris),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Notifications),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_RemoteInput),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_TrustAgents),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_UnredactedNotifications),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures{
		"allfeatures":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_AllFeatures,
		"biometrics":              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Biometrics,
		"camera":                  AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Camera,
		"face":                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Face,
		"fingerprint":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Fingerprint,
		"iris":                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Iris,
		"notconfigured":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_NotConfigured,
		"notifications":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_Notifications,
		"remoteinput":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_RemoteInput,
		"trustagents":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_TrustAgents,
		"unredactednotifications": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures_UnredactedNotifications,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures(input)
	return &out, nil
}
