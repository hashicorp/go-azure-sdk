package datareference

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataReferenceCredentialType string

const (
	DataReferenceCredentialTypeDockerCredentials DataReferenceCredentialType = "DockerCredentials"
	DataReferenceCredentialTypeManagedIdentity   DataReferenceCredentialType = "ManagedIdentity"
	DataReferenceCredentialTypeNoCredentials     DataReferenceCredentialType = "NoCredentials"
	DataReferenceCredentialTypeSAS               DataReferenceCredentialType = "SAS"
)

func PossibleValuesForDataReferenceCredentialType() []string {
	return []string{
		string(DataReferenceCredentialTypeDockerCredentials),
		string(DataReferenceCredentialTypeManagedIdentity),
		string(DataReferenceCredentialTypeNoCredentials),
		string(DataReferenceCredentialTypeSAS),
	}
}

func (s *DataReferenceCredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDataReferenceCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDataReferenceCredentialType(input string) (*DataReferenceCredentialType, error) {
	vals := map[string]DataReferenceCredentialType{
		"dockercredentials": DataReferenceCredentialTypeDockerCredentials,
		"managedidentity":   DataReferenceCredentialTypeManagedIdentity,
		"nocredentials":     DataReferenceCredentialTypeNoCredentials,
		"sas":               DataReferenceCredentialTypeSAS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataReferenceCredentialType(input)
	return &out, nil
}
