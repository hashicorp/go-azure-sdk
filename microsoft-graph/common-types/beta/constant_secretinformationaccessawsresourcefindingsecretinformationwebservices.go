package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretInformationAccessAwsResourceFindingSecretInformationWebServices string

const (
	SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateAuthority SecretInformationAccessAwsResourceFindingSecretInformationWebServices = "certificateAuthority"
	SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateManager   SecretInformationAccessAwsResourceFindingSecretInformationWebServices = "certificateManager"
	SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CloudHsm             SecretInformationAccessAwsResourceFindingSecretInformationWebServices = "cloudHsm"
	SecretInformationAccessAwsResourceFindingSecretInformationWebServices_SecretsManager       SecretInformationAccessAwsResourceFindingSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForSecretInformationAccessAwsResourceFindingSecretInformationWebServices() []string {
	return []string{
		string(SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateAuthority),
		string(SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateManager),
		string(SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CloudHsm),
		string(SecretInformationAccessAwsResourceFindingSecretInformationWebServices_SecretsManager),
	}
}

func (s *SecretInformationAccessAwsResourceFindingSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecretInformationAccessAwsResourceFindingSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecretInformationAccessAwsResourceFindingSecretInformationWebServices(input string) (*SecretInformationAccessAwsResourceFindingSecretInformationWebServices, error) {
	vals := map[string]SecretInformationAccessAwsResourceFindingSecretInformationWebServices{
		"certificateauthority": SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CertificateManager,
		"cloudhsm":             SecretInformationAccessAwsResourceFindingSecretInformationWebServices_CloudHsm,
		"secretsmanager":       SecretInformationAccessAwsResourceFindingSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretInformationAccessAwsResourceFindingSecretInformationWebServices(input)
	return &out, nil
}
