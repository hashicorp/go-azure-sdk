package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileKeyStorageProvider string

const (
	WindowsPhone81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp                     WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "UseSoftwareKsp"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail             WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseFail"
	WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   WindowsPhone81SCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail),
		string(WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindowsPhone81SCEPCertificateProfileKeyStorageProvider(input string) (*WindowsPhone81SCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": WindowsPhone81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   WindowsPhone81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
