package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileKeyStorageProvider string

const (
	Windows81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail Windows81SCEPCertificateProfileKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	Windows81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp                     Windows81SCEPCertificateProfileKeyStorageProvider = "UseSoftwareKsp"
	Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail             Windows81SCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseFail"
	Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   Windows81SCEPCertificateProfileKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForWindows81SCEPCertificateProfileKeyStorageProvider() []string {
	return []string{
		string(Windows81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(Windows81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp),
		string(Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail),
		string(Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseWindows81SCEPCertificateProfileKeyStorageProvider(input string) (*Windows81SCEPCertificateProfileKeyStorageProvider, error) {
	vals := map[string]Windows81SCEPCertificateProfileKeyStorageProvider{
		"usepassportforworkkspotherwisefail": Windows81SCEPCertificateProfileKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     Windows81SCEPCertificateProfileKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   Windows81SCEPCertificateProfileKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileKeyStorageProvider(input)
	return &out, nil
}
