package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseKeyStorageProvider string

const (
	WindowsCertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail WindowsCertificateProfileBaseKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	WindowsCertificateProfileBaseKeyStorageProvider_UseSoftwareKsp                     WindowsCertificateProfileBaseKeyStorageProvider = "useSoftwareKsp"
	WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail             WindowsCertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseFail"
	WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   WindowsCertificateProfileBaseKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsCertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(WindowsCertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(WindowsCertificateProfileBaseKeyStorageProvider_UseSoftwareKsp),
		string(WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *WindowsCertificateProfileBaseKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsCertificateProfileBaseKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsCertificateProfileBaseKeyStorageProvider(input string) (*WindowsCertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]WindowsCertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsCertificateProfileBaseKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsCertificateProfileBaseKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsCertificateProfileBaseKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
