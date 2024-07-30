package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateKeyStorageProvider string

const (
	ManagedDeviceCertificateStateCertificateKeyStorageProvider_UsePassportForWorkKspOtherwiseFail ManagedDeviceCertificateStateCertificateKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseSoftwareKsp                     ManagedDeviceCertificateStateCertificateKeyStorageProvider = "useSoftwareKsp"
	ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseFail             ManagedDeviceCertificateStateCertificateKeyStorageProvider = "useTpmKspOtherwiseFail"
	ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   ManagedDeviceCertificateStateCertificateKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateKeyStorageProvider() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseSoftwareKsp),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *ManagedDeviceCertificateStateCertificateKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateKeyStorageProvider(input string) (*ManagedDeviceCertificateStateCertificateKeyStorageProvider, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateKeyStorageProvider{
		"usepassportforworkkspotherwisefail": ManagedDeviceCertificateStateCertificateKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   ManagedDeviceCertificateStateCertificateKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateKeyStorageProvider(input)
	return &out, nil
}
