package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityImpactedDeviceAssetIdentifier string

const (
	SecurityImpactedDeviceAssetIdentifier_DestinationDeviceName SecurityImpactedDeviceAssetIdentifier = "destinationDeviceName"
	SecurityImpactedDeviceAssetIdentifier_DeviceId              SecurityImpactedDeviceAssetIdentifier = "deviceId"
	SecurityImpactedDeviceAssetIdentifier_DeviceName            SecurityImpactedDeviceAssetIdentifier = "deviceName"
	SecurityImpactedDeviceAssetIdentifier_RemoteDeviceName      SecurityImpactedDeviceAssetIdentifier = "remoteDeviceName"
	SecurityImpactedDeviceAssetIdentifier_TargetDeviceName      SecurityImpactedDeviceAssetIdentifier = "targetDeviceName"
)

func PossibleValuesForSecurityImpactedDeviceAssetIdentifier() []string {
	return []string{
		string(SecurityImpactedDeviceAssetIdentifier_DestinationDeviceName),
		string(SecurityImpactedDeviceAssetIdentifier_DeviceId),
		string(SecurityImpactedDeviceAssetIdentifier_DeviceName),
		string(SecurityImpactedDeviceAssetIdentifier_RemoteDeviceName),
		string(SecurityImpactedDeviceAssetIdentifier_TargetDeviceName),
	}
}

func (s *SecurityImpactedDeviceAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityImpactedDeviceAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityImpactedDeviceAssetIdentifier(input string) (*SecurityImpactedDeviceAssetIdentifier, error) {
	vals := map[string]SecurityImpactedDeviceAssetIdentifier{
		"destinationdevicename": SecurityImpactedDeviceAssetIdentifier_DestinationDeviceName,
		"deviceid":              SecurityImpactedDeviceAssetIdentifier_DeviceId,
		"devicename":            SecurityImpactedDeviceAssetIdentifier_DeviceName,
		"remotedevicename":      SecurityImpactedDeviceAssetIdentifier_RemoteDeviceName,
		"targetdevicename":      SecurityImpactedDeviceAssetIdentifier_TargetDeviceName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityImpactedDeviceAssetIdentifier(input)
	return &out, nil
}
