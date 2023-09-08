package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedDeviceAssetIdentifier string

const (
	SecurityImpactedDeviceAssetIdentifierdestinationDeviceName SecurityImpactedDeviceAssetIdentifier = "DestinationDeviceName"
	SecurityImpactedDeviceAssetIdentifierdeviceId              SecurityImpactedDeviceAssetIdentifier = "DeviceId"
	SecurityImpactedDeviceAssetIdentifierdeviceName            SecurityImpactedDeviceAssetIdentifier = "DeviceName"
	SecurityImpactedDeviceAssetIdentifierremoteDeviceName      SecurityImpactedDeviceAssetIdentifier = "RemoteDeviceName"
	SecurityImpactedDeviceAssetIdentifiertargetDeviceName      SecurityImpactedDeviceAssetIdentifier = "TargetDeviceName"
)

func PossibleValuesForSecurityImpactedDeviceAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedDeviceAssetIdentifierdestinationDeviceName),
		string(SecurityImpactedDeviceAssetIdentifierdeviceId),
		string(SecurityImpactedDeviceAssetIdentifierdeviceName),
		string(SecurityImpactedDeviceAssetIdentifierremoteDeviceName),
		string(SecurityImpactedDeviceAssetIdentifiertargetDeviceName),
	}
}

func parseSecurityImpactedDeviceAssetIdentifier(input string) (*SecurityImpactedDeviceAssetIdentifier, error) {
	vals := map[string]SecurityImpactedDeviceAssetIdentifier{
		"destinationdevicename": SecurityImpactedDeviceAssetIdentifierdestinationDeviceName,
		"deviceid":              SecurityImpactedDeviceAssetIdentifierdeviceId,
		"devicename":            SecurityImpactedDeviceAssetIdentifierdeviceName,
		"remotedevicename":      SecurityImpactedDeviceAssetIdentifierremoteDeviceName,
		"targetdevicename":      SecurityImpactedDeviceAssetIdentifiertargetDeviceName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedDeviceAssetIdentifier(input)
	return &out, nil
}
