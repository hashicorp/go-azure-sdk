package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileKeyStorageProvider string

const (
	Windows10XSCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows10XSCEPCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows10XSCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp                     Windows10XSCEPCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             Windows10XSCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows10XSCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10XSCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows10XSCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows10XSCEPCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10XSCEPCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10XSCEPCertificateProfileKeyStorageProvider(input string) (*Windows10XSCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10XSCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10XSCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10XSCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10XSCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
