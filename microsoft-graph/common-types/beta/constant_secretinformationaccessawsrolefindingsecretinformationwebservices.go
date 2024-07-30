package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretInformationAccessAwsRoleFindingSecretInformationWebServices string

const (
	SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateAuthority SecretInformationAccessAwsRoleFindingSecretInformationWebServices = "certificateAuthority"
	SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateManager   SecretInformationAccessAwsRoleFindingSecretInformationWebServices = "certificateManager"
	SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CloudHsm             SecretInformationAccessAwsRoleFindingSecretInformationWebServices = "cloudHsm"
	SecretInformationAccessAwsRoleFindingSecretInformationWebServices_SecretsManager       SecretInformationAccessAwsRoleFindingSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForSecretInformationAccessAwsRoleFindingSecretInformationWebServices() []string {
	return []string{
		string(SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateAuthority),
		string(SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateManager),
		string(SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CloudHsm),
		string(SecretInformationAccessAwsRoleFindingSecretInformationWebServices_SecretsManager),
	}
}

func (s *SecretInformationAccessAwsRoleFindingSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecretInformationAccessAwsRoleFindingSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecretInformationAccessAwsRoleFindingSecretInformationWebServices(input string) (*SecretInformationAccessAwsRoleFindingSecretInformationWebServices, error) {
	vals := map[string]SecretInformationAccessAwsRoleFindingSecretInformationWebServices{
		"certificateauthority": SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CertificateManager,
		"cloudhsm":             SecretInformationAccessAwsRoleFindingSecretInformationWebServices_CloudHsm,
		"secretsmanager":       SecretInformationAccessAwsRoleFindingSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretInformationAccessAwsRoleFindingSecretInformationWebServices(input)
	return &out, nil
}
