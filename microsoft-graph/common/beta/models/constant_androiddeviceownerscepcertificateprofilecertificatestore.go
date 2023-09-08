package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateStore string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateStoremachine AndroidDeviceOwnerScepCertificateProfileCertificateStore = "Machine"
	AndroidDeviceOwnerScepCertificateProfileCertificateStoreuser    AndroidDeviceOwnerScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateStoremachine),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateStoreuser),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateStore(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateStore{
		"machine": AndroidDeviceOwnerScepCertificateProfileCertificateStoremachine,
		"user":    AndroidDeviceOwnerScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateStore(input)
	return &out, nil
}
