package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretInformationAccessAwsUserFindingSecretInformationWebServices string

const (
	SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateAuthority SecretInformationAccessAwsUserFindingSecretInformationWebServices = "certificateAuthority"
	SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateManager   SecretInformationAccessAwsUserFindingSecretInformationWebServices = "certificateManager"
	SecretInformationAccessAwsUserFindingSecretInformationWebServices_CloudHsm             SecretInformationAccessAwsUserFindingSecretInformationWebServices = "cloudHsm"
	SecretInformationAccessAwsUserFindingSecretInformationWebServices_SecretsManager       SecretInformationAccessAwsUserFindingSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForSecretInformationAccessAwsUserFindingSecretInformationWebServices() []string {
	return []string{
		string(SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateAuthority),
		string(SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateManager),
		string(SecretInformationAccessAwsUserFindingSecretInformationWebServices_CloudHsm),
		string(SecretInformationAccessAwsUserFindingSecretInformationWebServices_SecretsManager),
	}
}

func (s *SecretInformationAccessAwsUserFindingSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecretInformationAccessAwsUserFindingSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecretInformationAccessAwsUserFindingSecretInformationWebServices(input string) (*SecretInformationAccessAwsUserFindingSecretInformationWebServices, error) {
	vals := map[string]SecretInformationAccessAwsUserFindingSecretInformationWebServices{
		"certificateauthority": SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   SecretInformationAccessAwsUserFindingSecretInformationWebServices_CertificateManager,
		"cloudhsm":             SecretInformationAccessAwsUserFindingSecretInformationWebServices_CloudHsm,
		"secretsmanager":       SecretInformationAccessAwsUserFindingSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretInformationAccessAwsUserFindingSecretInformationWebServices(input)
	return &out, nil
}
