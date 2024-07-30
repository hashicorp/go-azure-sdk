package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileKeyStorageProvider string

const (
	Windows10PkcsCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows10PkcsCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows10PkcsCertificateProfileKeyStorageProvider_UseSoftwareKsp                     Windows10PkcsCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             Windows10PkcsCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows10PkcsCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10PkcsCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10PkcsCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows10PkcsCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows10PkcsCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PkcsCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PkcsCertificateProfileKeyStorageProvider(input string) (*Windows10PkcsCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10PkcsCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10PkcsCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10PkcsCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10PkcsCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
