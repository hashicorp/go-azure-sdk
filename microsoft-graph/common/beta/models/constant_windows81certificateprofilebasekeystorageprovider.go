package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseKeyStorageProvider string

const (
	Windows81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail Windows81CertificateProfileBaseKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	Windows81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp                     Windows81CertificateProfileBaseKeyStorageProvider = "UseSoftwareKsp"
	Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail             Windows81CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseFail"
	Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   Windows81CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows81CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(Windows81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(Windows81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp),
		string(Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail),
		string(Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindows81CertificateProfileBaseKeyStorageProvider(input string) (*Windows81CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]Windows81CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
