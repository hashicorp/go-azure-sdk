package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider string

const (
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp                     WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "UseSoftwareKsp"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail             WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseFail"
	WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail),
		string(WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider(input string) (*WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81ImportedPFXCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
