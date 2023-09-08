package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificateStore string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificateStoremachine AospDeviceOwnerPkcsCertificateProfileCertificateStore = "Machine"
	AospDeviceOwnerPkcsCertificateProfileCertificateStoreuser    AospDeviceOwnerPkcsCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificateStore() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificateStoremachine),
		string(AospDeviceOwnerPkcsCertificateProfileCertificateStoreuser),
	}
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificateStore(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificateStore, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificateStore{
		"machine": AospDeviceOwnerPkcsCertificateProfileCertificateStoremachine,
		"user":    AospDeviceOwnerPkcsCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificateStore(input)
	return &out, nil
}
