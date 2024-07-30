package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices string

const (
	SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateAuthority SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices = "certificateAuthority"
	SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateManager   SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices = "certificateManager"
	SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CloudHsm             SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices = "cloudHsm"
	SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_SecretsManager       SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices = "secretsManager"
)

func PossibleValuesForSecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices() []string {
	return []string{
		string(SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateAuthority),
		string(SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateManager),
		string(SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CloudHsm),
		string(SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_SecretsManager),
	}
}

func (s *SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices(input string) (*SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices, error) {
	vals := map[string]SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices{
		"certificateauthority": SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateAuthority,
		"certificatemanager":   SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CertificateManager,
		"cloudhsm":             SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_CloudHsm,
		"secretsmanager":       SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices_SecretsManager,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretInformationAccessAwsServerlessFunctionFindingSecretInformationWebServices(input)
	return &out, nil
}
