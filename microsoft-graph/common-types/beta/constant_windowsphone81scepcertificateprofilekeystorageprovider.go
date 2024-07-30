package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileKeyStorageProvider string

const (
	WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "usePassportForWorkKspOtherwiseFail"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp                     WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "useSoftwareKsp"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail             WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseFail"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileKeyStorageProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileKeyStorageProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileKeyStorageProvider(input string) (*WindowsPhone81SCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81SCEPCertificateProfileKeyStorageProvider_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
