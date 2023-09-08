package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseKeyStorageProvider string

const (
	WindowsCertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail WindowsCertificateProfileBaseKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	WindowsCertificateProfileBaseKeyStorageProvideruseSoftwareKsp                     WindowsCertificateProfileBaseKeyStorageProvider = "UseSoftwareKsp"
	WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail             WindowsCertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseFail"
	WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   WindowsCertificateProfileBaseKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsCertificateProfileBaseKeyStorageProvider() []string {
	return []string{
		string(WindowsCertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(WindowsCertificateProfileBaseKeyStorageProvideruseSoftwareKsp),
		string(WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail),
		string(WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindowsCertificateProfileBaseKeyStorageProvider(input string) (*WindowsCertificateProfileBaseKeyStorageProvider, error) {
	vals := map[string]WindowsCertificateProfileBaseKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsCertificateProfileBaseKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsCertificateProfileBaseKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsCertificateProfileBaseKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseKeyStorageProvider(input)
	return &out, nil
}
