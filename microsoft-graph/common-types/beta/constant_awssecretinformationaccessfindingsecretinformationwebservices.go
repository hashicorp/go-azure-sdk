package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSecretInformationAccessFindingSecretInformationWebServices string

const (
	AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateAuthority AwsSecretInformationAccessFindingSecretInformationWebServices = "certificateAuthority"
	AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateManager   AwsSecretInformationAccessFindingSecretInformationWebServices = "certificateManager"
	AwsSecretInformationAccessFindingSecretInformationWebServices_CloudHsm             AwsSecretInformationAccessFindingSecretInformationWebServices = "cloudHsm"
	AwsSecretInformationAccessFindingSecretInformationWebServices_SecretsManager       AwsSecretInformationAccessFindingSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForAwsSecretInformationAccessFindingSecretInformationWebServices() []string {
	return []string{
		string(AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateAuthority),
		string(AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateManager),
		string(AwsSecretInformationAccessFindingSecretInformationWebServices_CloudHsm),
		string(AwsSecretInformationAccessFindingSecretInformationWebServices_SecretsManager),
	}
}

func (s *AwsSecretInformationAccessFindingSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsSecretInformationAccessFindingSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsSecretInformationAccessFindingSecretInformationWebServices(input string) (*AwsSecretInformationAccessFindingSecretInformationWebServices, error) {
	vals := map[string]AwsSecretInformationAccessFindingSecretInformationWebServices{
		"certificateauthority": AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   AwsSecretInformationAccessFindingSecretInformationWebServices_CertificateManager,
		"cloudhsm":             AwsSecretInformationAccessFindingSecretInformationWebServices_CloudHsm,
		"secretsmanager":       AwsSecretInformationAccessFindingSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsSecretInformationAccessFindingSecretInformationWebServices(input)
	return &out, nil
}
