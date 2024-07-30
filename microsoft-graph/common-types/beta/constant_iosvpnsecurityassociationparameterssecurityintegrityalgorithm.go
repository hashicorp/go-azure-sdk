package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm string

const (
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Md5     IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "md5"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha1160 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "sha1_160"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha196  IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "sha1_96"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2256 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "sha2_256"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2384 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "sha2_384"
	IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2512 IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm = "sha2_512"
)

func PossibleValuesForIosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm() []string {
	return []string{
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Md5),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha1160),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha196),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2256),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2384),
		string(IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2512),
	}
}

func (s *IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm(input string) (*IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm, error) {
	vals := map[string]IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm{
		"md5":      IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Md5,
		"sha1_160": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha1160,
		"sha1_96":  IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha196,
		"sha2_256": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2256,
		"sha2_384": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2384,
		"sha2_512": IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm_Sha2512,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnSecurityAssociationParametersSecurityIntegrityAlgorithm(input)
	return &out, nil
}
