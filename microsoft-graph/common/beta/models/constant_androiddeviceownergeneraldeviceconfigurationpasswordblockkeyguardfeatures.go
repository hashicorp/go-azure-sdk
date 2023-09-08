package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesallFeatures             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "AllFeatures"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesbiometrics              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Biometrics"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturescamera                  AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Camera"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesface                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Face"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesfingerprint             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Fingerprint"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesiris                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Iris"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotConfigured           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotifications           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "Notifications"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesremoteInput             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "RemoteInput"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturestrustAgents             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "TrustAgents"
	AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesunredactedNotifications AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures = "UnredactedNotifications"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesallFeatures),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesbiometrics),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturescamera),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesface),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesfingerprint),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesiris),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotifications),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesremoteInput),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturestrustAgents),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesunredactedNotifications),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures{
		"allfeatures":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesallFeatures,
		"biometrics":              AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesbiometrics,
		"camera":                  AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturescamera,
		"face":                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesface,
		"fingerprint":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesfingerprint,
		"iris":                    AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesiris,
		"notconfigured":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotConfigured,
		"notifications":           AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesnotifications,
		"remoteinput":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesremoteInput,
		"trustagents":             AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturestrustAgents,
		"unredactednotifications": AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeaturesunredactedNotifications,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationPasswordBlockKeyguardFeatures(input)
	return &out, nil
}
