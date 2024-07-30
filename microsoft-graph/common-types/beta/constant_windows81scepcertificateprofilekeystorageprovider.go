package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileKeyStorageProvider string

const (
	Windows81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail Windows81SCEPCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	Windows81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp                     Windows81SCEPCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             Windows81SCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   Windows81SCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows81SCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(Windows81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *Windows81SCEPCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileKeyStorageProvider(input string) (*Windows81SCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows81SCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
