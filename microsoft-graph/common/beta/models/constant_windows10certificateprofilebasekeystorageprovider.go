package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseKeyStorageProvider string

const (
	Windows10CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail Windows10CertificateProfileBaseKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	Windows10CertificateProfileBaseKeyStorageProvideruseSoftwareKsp                     Windows10CertificateProfileBaseKeyStorageProvider = "UseSoftwareKsp"
	Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail             Windows10CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseFail"
	Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   Windows10CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(Windows10CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(Windows10CertificateProfileBaseKeyStorageProvideruseSoftwareKsp),
		string(Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail),
		string(Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindows10CertificateProfileBaseKeyStorageProvider(input string) (*Windows10CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]Windows10CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10CertificateProfileBaseKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
