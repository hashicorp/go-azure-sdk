package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm string

const (
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes128"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes128Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes192"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes192Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes256"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256Gcm        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "aes256Gcm"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_ChaCha20Poly1305 IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "chaCha20Poly1305"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Des              IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "des"
	IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_TripleDes        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm = "tripleDes"
)

func PossibleValuesForIosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm() []string {
	return []string{
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256Gcm),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_ChaCha20Poly1305),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Des),
		string(IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_TripleDes),
	}
}

func (s *IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm(input string) (*IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm, error) {
	vals := map[string]IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm{
		"aes128":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128,
		"aes128gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes128Gcm,
		"aes192":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192,
		"aes192gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes192Gcm,
		"aes256":           IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256,
		"aes256gcm":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Aes256Gcm,
		"chacha20poly1305": IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_ChaCha20Poly1305,
		"des":              IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_Des,
		"tripledes":        IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm_TripleDes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnSecurityAssociationParametersSecurityEncryptionAlgorithm(input)
	return &out, nil
}
