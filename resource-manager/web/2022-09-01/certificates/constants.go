package certificates

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultSecretStatus string

const (
	KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault KeyVaultSecretStatus = "AzureServiceUnauthorizedToAccessKeyVault"
	KeyVaultSecretStatusCertificateOrderFailed                   KeyVaultSecretStatus = "CertificateOrderFailed"
	KeyVaultSecretStatusExternalPrivateKey                       KeyVaultSecretStatus = "ExternalPrivateKey"
	KeyVaultSecretStatusInitialized                              KeyVaultSecretStatus = "Initialized"
	KeyVaultSecretStatusKeyVaultDoesNotExist                     KeyVaultSecretStatus = "KeyVaultDoesNotExist"
	KeyVaultSecretStatusKeyVaultSecretDoesNotExist               KeyVaultSecretStatus = "KeyVaultSecretDoesNotExist"
	KeyVaultSecretStatusOperationNotPermittedOnKeyVault          KeyVaultSecretStatus = "OperationNotPermittedOnKeyVault"
	KeyVaultSecretStatusSucceeded                                KeyVaultSecretStatus = "Succeeded"
	KeyVaultSecretStatusUnknown                                  KeyVaultSecretStatus = "Unknown"
	KeyVaultSecretStatusUnknownError                             KeyVaultSecretStatus = "UnknownError"
	KeyVaultSecretStatusWaitingOnCertificateOrder                KeyVaultSecretStatus = "WaitingOnCertificateOrder"
)

func PossibleValuesForKeyVaultSecretStatus() []string {
	return []string{
		string(KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault),
		string(KeyVaultSecretStatusCertificateOrderFailed),
		string(KeyVaultSecretStatusExternalPrivateKey),
		string(KeyVaultSecretStatusInitialized),
		string(KeyVaultSecretStatusKeyVaultDoesNotExist),
		string(KeyVaultSecretStatusKeyVaultSecretDoesNotExist),
		string(KeyVaultSecretStatusOperationNotPermittedOnKeyVault),
		string(KeyVaultSecretStatusSucceeded),
		string(KeyVaultSecretStatusUnknown),
		string(KeyVaultSecretStatusUnknownError),
		string(KeyVaultSecretStatusWaitingOnCertificateOrder),
	}
}

func parseKeyVaultSecretStatus(input string) (*KeyVaultSecretStatus, error) {
	vals := map[string]KeyVaultSecretStatus{
		"azureserviceunauthorizedtoaccesskeyvault": KeyVaultSecretStatusAzureServiceUnauthorizedToAccessKeyVault,
		"certificateorderfailed":                   KeyVaultSecretStatusCertificateOrderFailed,
		"externalprivatekey":                       KeyVaultSecretStatusExternalPrivateKey,
		"initialized":                              KeyVaultSecretStatusInitialized,
		"keyvaultdoesnotexist":                     KeyVaultSecretStatusKeyVaultDoesNotExist,
		"keyvaultsecretdoesnotexist":               KeyVaultSecretStatusKeyVaultSecretDoesNotExist,
		"operationnotpermittedonkeyvault":          KeyVaultSecretStatusOperationNotPermittedOnKeyVault,
		"succeeded":                                KeyVaultSecretStatusSucceeded,
		"unknown":                                  KeyVaultSecretStatusUnknown,
		"unknownerror":                             KeyVaultSecretStatusUnknownError,
		"waitingoncertificateorder":                KeyVaultSecretStatusWaitingOnCertificateOrder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyVaultSecretStatus(input)
	return &out, nil
}
