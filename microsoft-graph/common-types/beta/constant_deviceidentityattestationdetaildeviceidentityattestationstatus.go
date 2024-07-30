package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus string

const (
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_IncompleteData DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "incompleteData"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_NotSupported   DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "notSupported"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Trusted        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "trusted"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_UnTrusted      DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "unTrusted"
	DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Unknown        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus = "unknown"
)

func PossibleValuesForDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus() []string {
	return []string{
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_IncompleteData),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_NotSupported),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Trusted),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_UnTrusted),
		string(DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Unknown),
	}
}

func (s *DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(input string) (*DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus, error) {
	vals := map[string]DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus{
		"incompletedata": DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_IncompleteData,
		"notsupported":   DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_NotSupported,
		"trusted":        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Trusted,
		"untrusted":      DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_UnTrusted,
		"unknown":        DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceIdentityAttestationDetailDeviceIdentityAttestationStatus(input)
	return &out, nil
}
