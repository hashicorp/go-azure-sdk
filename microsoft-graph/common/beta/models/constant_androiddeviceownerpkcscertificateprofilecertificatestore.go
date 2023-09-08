package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateStore string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateStoremachine AndroidDeviceOwnerPkcsCertificateProfileCertificateStore = "Machine"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateStoreuser    AndroidDeviceOwnerPkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateStoremachine),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateStoreuser),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateStore(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateStore{
		"machine": AndroidDeviceOwnerPkcsCertificateProfileCertificateStoremachine,
		"user":    AndroidDeviceOwnerPkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
