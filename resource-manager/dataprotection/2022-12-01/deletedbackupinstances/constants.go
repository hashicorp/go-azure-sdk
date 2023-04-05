package deletedbackupinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CurrentProtectionState string

const (
	CurrentProtectionStateBackupSchedulesSuspended    CurrentProtectionState = "BackupSchedulesSuspended"
	CurrentProtectionStateConfiguringProtection       CurrentProtectionState = "ConfiguringProtection"
	CurrentProtectionStateConfiguringProtectionFailed CurrentProtectionState = "ConfiguringProtectionFailed"
	CurrentProtectionStateInvalid                     CurrentProtectionState = "Invalid"
	CurrentProtectionStateNotProtected                CurrentProtectionState = "NotProtected"
	CurrentProtectionStateProtectionConfigured        CurrentProtectionState = "ProtectionConfigured"
	CurrentProtectionStateProtectionError             CurrentProtectionState = "ProtectionError"
	CurrentProtectionStateProtectionStopped           CurrentProtectionState = "ProtectionStopped"
	CurrentProtectionStateRetentionSchedulesSuspended CurrentProtectionState = "RetentionSchedulesSuspended"
	CurrentProtectionStateSoftDeleted                 CurrentProtectionState = "SoftDeleted"
	CurrentProtectionStateSoftDeleting                CurrentProtectionState = "SoftDeleting"
	CurrentProtectionStateUpdatingProtection          CurrentProtectionState = "UpdatingProtection"
)

func PossibleValuesForCurrentProtectionState() []string {
	return []string{
		string(CurrentProtectionStateBackupSchedulesSuspended),
		string(CurrentProtectionStateConfiguringProtection),
		string(CurrentProtectionStateConfiguringProtectionFailed),
		string(CurrentProtectionStateInvalid),
		string(CurrentProtectionStateNotProtected),
		string(CurrentProtectionStateProtectionConfigured),
		string(CurrentProtectionStateProtectionError),
		string(CurrentProtectionStateProtectionStopped),
		string(CurrentProtectionStateRetentionSchedulesSuspended),
		string(CurrentProtectionStateSoftDeleted),
		string(CurrentProtectionStateSoftDeleting),
		string(CurrentProtectionStateUpdatingProtection),
	}
}

type DataStoreTypes string

const (
	DataStoreTypesArchiveStore     DataStoreTypes = "ArchiveStore"
	DataStoreTypesOperationalStore DataStoreTypes = "OperationalStore"
	DataStoreTypesVaultStore       DataStoreTypes = "VaultStore"
)

func PossibleValuesForDataStoreTypes() []string {
	return []string{
		string(DataStoreTypesArchiveStore),
		string(DataStoreTypesOperationalStore),
		string(DataStoreTypesVaultStore),
	}
}

type SecretStoreType string

const (
	SecretStoreTypeAzureKeyVault SecretStoreType = "AzureKeyVault"
	SecretStoreTypeInvalid       SecretStoreType = "Invalid"
)

func PossibleValuesForSecretStoreType() []string {
	return []string{
		string(SecretStoreTypeAzureKeyVault),
		string(SecretStoreTypeInvalid),
	}
}

type Status string

const (
	StatusConfiguringProtection       Status = "ConfiguringProtection"
	StatusConfiguringProtectionFailed Status = "ConfiguringProtectionFailed"
	StatusProtectionConfigured        Status = "ProtectionConfigured"
	StatusProtectionStopped           Status = "ProtectionStopped"
	StatusSoftDeleted                 Status = "SoftDeleted"
	StatusSoftDeleting                Status = "SoftDeleting"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusConfiguringProtection),
		string(StatusConfiguringProtectionFailed),
		string(StatusProtectionConfigured),
		string(StatusProtectionStopped),
		string(StatusSoftDeleted),
		string(StatusSoftDeleting),
	}
}

type ValidationType string

const (
	ValidationTypeDeepValidation    ValidationType = "DeepValidation"
	ValidationTypeShallowValidation ValidationType = "ShallowValidation"
)

func PossibleValuesForValidationType() []string {
	return []string{
		string(ValidationTypeDeepValidation),
		string(ValidationTypeShallowValidation),
	}
}
