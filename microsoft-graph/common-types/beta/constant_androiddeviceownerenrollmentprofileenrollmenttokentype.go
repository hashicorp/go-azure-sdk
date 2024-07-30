package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType string

const (
	AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType = "corporateOwnedDedicatedDeviceWithAzureADSharedMode"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_Default                                            AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType = "default"
	AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_DeviceStaging                                      AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType = "deviceStaging"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_Default),
		string(AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_DeviceStaging),
	}
}

func (s *AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType(input string) (*AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType{
		"corporateowneddedicateddevicewithazureadsharedmode": AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_CorporateOwnedDedicatedDeviceWithAzureADSharedMode,
		"default":       AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_Default,
		"devicestaging": AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType_DeviceStaging,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileEnrollmentTokenType(input)
	return &out, nil
}
