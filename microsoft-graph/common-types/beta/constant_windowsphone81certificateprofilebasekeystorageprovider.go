package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseKeyStorageProvider string

const (
	WindowsPhone81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail WindowsPhone81CertificateProfileBaseKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp                     WindowsPhone81CertificateProfileBaseKeyStorageProvider = "useSoftwareKsp"
	WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail             WindowsPhone81CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseFail"
	WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *WindowsPhone81CertificateProfileBaseKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81CertificateProfileBaseKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81CertificateProfileBaseKeyStorageProvider(input string) (*WindowsPhone81CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
