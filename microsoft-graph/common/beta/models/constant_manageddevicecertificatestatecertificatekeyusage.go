package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateKeyUsage string

const (
	ManagedDeviceCertificateStateCertificateKeyUsagedigitalSignature ManagedDeviceCertificateStateCertificateKeyUsage = "DigitalSignature"
	ManagedDeviceCertificateStateCertificateKeyUsagekeyEncipherment  ManagedDeviceCertificateStateCertificateKeyUsage = "KeyEncipherment"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateKeyUsage() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateKeyUsagedigitalSignature),
		string(ManagedDeviceCertificateStateCertificateKeyUsagekeyEncipherment),
	}
}

func parseManagedDeviceCertificateStateCertificateKeyUsage(input string) (*ManagedDeviceCertificateStateCertificateKeyUsage, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateKeyUsage{
		"digitalsignature": ManagedDeviceCertificateStateCertificateKeyUsagedigitalSignature,
		"keyencipherment":  ManagedDeviceCertificateStateCertificateKeyUsagekeyEncipherment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateKeyUsage(input)
	return &out, nil
}
