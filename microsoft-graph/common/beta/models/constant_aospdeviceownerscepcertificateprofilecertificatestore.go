package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileCertificateStore string

const (
	AospDeviceOwnerScepCertificateProfileCertificateStoremachine AospDeviceOwnerScepCertificateProfileCertificateStore = "Machine"
	AospDeviceOwnerScepCertificateProfileCertificateStoreuser    AospDeviceOwnerScepCertificateProfileCertificateStore = "User"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileCertificateStore() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileCertificateStoremachine),
		string(AospDeviceOwnerScepCertificateProfileCertificateStoreuser),
	}
}

func parseAospDeviceOwnerScepCertificateProfileCertificateStore(input string) (*AospDeviceOwnerScepCertificateProfileCertificateStore, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileCertificateStore{
		"machine": AospDeviceOwnerScepCertificateProfileCertificateStoremachine,
		"user":    AospDeviceOwnerScepCertificateProfileCertificateStoreuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileCertificateStore(input)
	return &out, nil
}
