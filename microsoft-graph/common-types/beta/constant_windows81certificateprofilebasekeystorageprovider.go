package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseKeyStorageProvider string

const (
	Windows81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows81CertificateProfileBaseKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp                     Windows81CertificateProfileBaseKeyStorageProvider = "useSoftwareKsp"
	Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail             Windows81CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows81CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows81CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(Windows81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp),
		string(Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows81CertificateProfileBaseKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81CertificateProfileBaseKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81CertificateProfileBaseKeyStorageProvider(input string) (*Windows81CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]Windows81CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows81CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows81CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows81CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
