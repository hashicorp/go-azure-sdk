package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy string

const (
	ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_Customer          ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy = "customer"
	ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftKeyVault ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy = "microsoftKeyVault"
	ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftStorage  ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy = "microsoftStorage"
)

func PossibleValuesForExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy() []string {
	return []string{
		string(ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_Customer),
		string(ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftKeyVault),
		string(ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftStorage),
	}
}

func (s *ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy(input string) (*ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy, error) {
	vals := map[string]ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy{
		"customer":          ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_Customer,
		"microsoftkeyvault": ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftKeyVault,
		"microsoftstorage":  ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy_MicrosoftStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternallyAccessibleAzureBlobContainerFindingEncryptionManagedBy(input)
	return &out, nil
}
