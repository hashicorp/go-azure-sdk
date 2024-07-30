package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptedAzureStorageAccountFindingEncryptionManagedBy string

const (
	EncryptedAzureStorageAccountFindingEncryptionManagedBy_Customer          EncryptedAzureStorageAccountFindingEncryptionManagedBy = "customer"
	EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftKeyVault EncryptedAzureStorageAccountFindingEncryptionManagedBy = "microsoftKeyVault"
	EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftStorage  EncryptedAzureStorageAccountFindingEncryptionManagedBy = "microsoftStorage"
)

func PossibleValuesForEncryptedAzureStorageAccountFindingEncryptionManagedBy() []string {
	return []string{
		string(EncryptedAzureStorageAccountFindingEncryptionManagedBy_Customer),
		string(EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftKeyVault),
		string(EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftStorage),
	}
}

func (s *EncryptedAzureStorageAccountFindingEncryptionManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEncryptedAzureStorageAccountFindingEncryptionManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEncryptedAzureStorageAccountFindingEncryptionManagedBy(input string) (*EncryptedAzureStorageAccountFindingEncryptionManagedBy, error) {
	vals := map[string]EncryptedAzureStorageAccountFindingEncryptionManagedBy{
		"customer":          EncryptedAzureStorageAccountFindingEncryptionManagedBy_Customer,
		"microsoftkeyvault": EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftKeyVault,
		"microsoftstorage":  EncryptedAzureStorageAccountFindingEncryptionManagedBy_MicrosoftStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptedAzureStorageAccountFindingEncryptionManagedBy(input)
	return &out, nil
}
