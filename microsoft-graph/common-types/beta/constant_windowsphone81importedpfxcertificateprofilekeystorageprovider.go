package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider string

const (
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp                     WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider(input string) (*WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
