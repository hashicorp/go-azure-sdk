package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm string

const (
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes128"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes128Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes192"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes192Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes256"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Aes256Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmchaCha20Poly1305 IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "ChaCha20Poly1305"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmdes              IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "Des"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmtripleDes        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "TripleDes"
)

func PossibleValuesForIosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm() []string {
	return []string{
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmchaCha20Poly1305),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmdes),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmtripleDes),
	}
}

func parseIosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm(input string) (*IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm, error) {
	vals := map[string]IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm{
		"aes128":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128,
		"aes128gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes128Gcm,
		"aes192":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192,
		"aes192gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes192Gcm,
		"aes256":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256,
		"aes256gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmaes256Gcm,
		"chacha20poly1305": IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmchaCha20Poly1305,
		"des":              IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmdes,
		"tripledes":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithmtripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm(input)
	return &out, nil
}
