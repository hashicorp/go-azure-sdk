package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseKeyStorageProvider string

const (
	Windows10CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows10CertificateProfileBaseKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows10CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp                     Windows10CertificateProfileBaseKeyStorageProvider = "useSoftwareKsp"
	Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail             Windows10CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows10CertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(Windows10CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows10CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp),
		string(Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows10CertificateProfileBaseKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CertificateProfileBaseKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CertificateProfileBaseKeyStorageProvider(input string) (*Windows10CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]Windows10CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10CertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10CertificateProfileBaseKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10CertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
