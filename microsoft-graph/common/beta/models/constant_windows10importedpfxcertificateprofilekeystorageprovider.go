package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileKeyStorageProvider string

const (
	Windows10ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail Windows10ImportedPFXCertificateProfileKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	Windows10ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp                     Windows10ImportedPFXCertificateProfileKeyStorageProvider = "UseSoftwareKsp"
	Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail             Windows10ImportedPFXCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseFail"
	Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   Windows10ImportedPFXCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail),
		string(Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindows10ImportedPFXCertificateProfileKeyStorageProvider(input string) (*Windows10ImportedPFXCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows10ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows10ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows10ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
