package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PFXImportCertificateProfileKeyStorageProvider string

const (
	Windows10PFXImportCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows10PFXImportCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows10PFXImportCertificateProfileKeyStorageProvider_UseSoftwareKsp                     Windows10PFXImportCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             Windows10PFXImportCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows10PFXImportCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10PFXImportCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10PFXImportCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows10PFXImportCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows10PFXImportCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PFXImportCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PFXImportCertificateProfileKeyStorageProvider(input string) (*Windows10PFXImportCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10PFXImportCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10PFXImportCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10PFXImportCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10PFXImportCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PFXImportCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
