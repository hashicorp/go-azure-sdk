package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseKeyStorageProvider string

const (
	WindowsPhone81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail WindowsPhone81CertificateProfileBaseKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	WindowsPhone81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp                     WindowsPhone81CertificateProfileBaseKeyStorageProvider = "UseSoftwareKsp"
	WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail             WindowsPhone81CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseFail"
	WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81CertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail),
		string(WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindowsPhone81CertificateProfileBaseKeyStorageProvider(input string) (*WindowsPhone81CertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81CertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81CertificateProfileBaseKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81CertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
