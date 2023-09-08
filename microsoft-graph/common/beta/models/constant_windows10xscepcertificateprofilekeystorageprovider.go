package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10XSCEPCertificateProfileKeyStorageProvider string

const (
	Windows10XSCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail Windows10XSCEPCertificateProfileKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	Windows10XSCEPCertificateProfileKeyStorageProvideruseSoftwareKsp                     Windows10XSCEPCertificateProfileKeyStorageProvider = "UseSoftwareKsp"
	Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail             Windows10XSCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseFail"
	Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   Windows10XSCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10XSCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10XSCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(Windows10XSCEPCertificateProfileKeyStorageProvideruseSoftwareKsp),
		string(Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail),
		string(Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindows10XSCEPCertificateProfileKeyStorageProvider(input string) (*Windows10XSCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10XSCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10XSCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10XSCEPCertificateProfileKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10XSCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10XSCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
