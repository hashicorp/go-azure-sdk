package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm string

const (
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmmd5      IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Md5"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_160 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Sha1160"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_96  IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Sha196"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_256 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Sha2256"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_384 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Sha2384"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_512 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "Sha2512"
)

func PossibleValuesForIosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm() []string {
	return []string{
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmmd5),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_160),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_96),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_256),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_384),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_512),
	}
}

func parseIosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm(input string) (*IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm, error) {
	vals := map[string]IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm{
		"md5":     IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmmd5,
		"sha1160": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_160,
		"sha196":  IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha1_96,
		"sha2256": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_256,
		"sha2384": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_384,
		"sha2512": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithmsha2_512,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm(input)
	return &out, nil
}
