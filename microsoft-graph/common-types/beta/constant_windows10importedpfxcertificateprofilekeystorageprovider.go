package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileKeyStorageProvider string

const (
	Windows10ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows10ImportedPFXCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp                     Windows10ImportedPFXCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             Windows10ImportedPFXCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows10ImportedPFXCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows10ImportedPFXCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10ImportedPFXCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10ImportedPFXCertificateProfileKeyStorageProvider(input string) (*Windows10ImportedPFXCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
