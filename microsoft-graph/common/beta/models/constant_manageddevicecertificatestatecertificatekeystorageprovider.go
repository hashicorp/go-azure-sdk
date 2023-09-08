package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateKeyStorageProvider string

const (
	ManagedDeviceCertificateStateCertificateKeyStorageProviderusePassportForWorkKspOtherwiseFail ManagedDeviceCertificateStateCertificateKeyStorageProvider = "UsePassportForWorkKspOtherwiseFail"
	ManagedDeviceCertificateStateCertificateKeyStorageProvideruseSoftwareKsp                     ManagedDeviceCertificateStateCertificateKeyStorageProvider = "UseSoftwareKsp"
	ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseFail             ManagedDeviceCertificateStateCertificateKeyStorageProvider = "UseTpmKspOtherwiseFail"
	ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp   ManagedDeviceCertificateStateCertificateKeyStorageProvider = "UseTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateKeyStorageProvider() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateKeyStorageProviderusePassportForWorkKspOtherwiseFail),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvideruseSoftwareKsp),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseFail),
		string(ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func parseManagedDeviceCertificateStateCertificateKeyStorageProvider(input string) (*ManagedDeviceCertificateStateCertificateKeyStorageProvider, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateKeyStorageProvider{
		"usepassportforworkkspotherwisefail": ManagedDeviceCertificateStateCertificateKeyStorageProviderusePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     ManagedDeviceCertificateStateCertificateKeyStorageProvideruseSoftwareKsp,
		"usetpmkspotherwisefail":             ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   ManagedDeviceCertificateStateCertificateKeyStorageProvideruseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateKeyStorageProvider(input)
	return &out, nil
}
