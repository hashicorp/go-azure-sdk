package protectionintentresources

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BackupManagementType string

const (
	BackupManagementTypeAzureBackupServer BackupManagementType = "AzureBackupServer"
	BackupManagementTypeAzureIaasVM       BackupManagementType = "AzureIaasVM"
	BackupManagementTypeAzureSql          BackupManagementType = "AzureSql"
	BackupManagementTypeAzureStorage      BackupManagementType = "AzureStorage"
	BackupManagementTypeAzureWorkload     BackupManagementType = "AzureWorkload"
	BackupManagementTypeDPM               BackupManagementType = "DPM"
	BackupManagementTypeDefaultBackup     BackupManagementType = "DefaultBackup"
	BackupManagementTypeInvalid           BackupManagementType = "Invalid"
	BackupManagementTypeMAB               BackupManagementType = "MAB"
)

func PossibleValuesForBackupManagementType() []string {
	return []string{
		string(BackupManagementTypeAzureBackupServer),
		string(BackupManagementTypeAzureIaasVM),
		string(BackupManagementTypeAzureSql),
		string(BackupManagementTypeAzureStorage),
		string(BackupManagementTypeAzureWorkload),
		string(BackupManagementTypeDPM),
		string(BackupManagementTypeDefaultBackup),
		string(BackupManagementTypeInvalid),
		string(BackupManagementTypeMAB),
	}
}

func parseBackupManagementType(input string) (*BackupManagementType, error) {
	vals := map[string]BackupManagementType{
		"azurebackupserver": BackupManagementTypeAzureBackupServer,
		"azureiaasvm":       BackupManagementTypeAzureIaasVM,
		"azuresql":          BackupManagementTypeAzureSql,
		"azurestorage":      BackupManagementTypeAzureStorage,
		"azureworkload":     BackupManagementTypeAzureWorkload,
		"dpm":               BackupManagementTypeDPM,
		"defaultbackup":     BackupManagementTypeDefaultBackup,
		"invalid":           BackupManagementTypeInvalid,
		"mab":               BackupManagementTypeMAB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupManagementType(input)
	return &out, nil
}

type ProtectionIntentItemType string

const (
	ProtectionIntentItemTypeAzureResourceItem                          ProtectionIntentItemType = "AzureResourceItem"
	ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent          ProtectionIntentItemType = "AzureWorkloadAutoProtectionIntent"
	ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent ProtectionIntentItemType = "AzureWorkloadContainerAutoProtectionIntent"
	ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent       ProtectionIntentItemType = "AzureWorkloadSQLAutoProtectionIntent"
	ProtectionIntentItemTypeInvalid                                    ProtectionIntentItemType = "Invalid"
	ProtectionIntentItemTypeRecoveryServiceVaultItem                   ProtectionIntentItemType = "RecoveryServiceVaultItem"
)

func PossibleValuesForProtectionIntentItemType() []string {
	return []string{
		string(ProtectionIntentItemTypeAzureResourceItem),
		string(ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent),
		string(ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent),
		string(ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent),
		string(ProtectionIntentItemTypeInvalid),
		string(ProtectionIntentItemTypeRecoveryServiceVaultItem),
	}
}

func parseProtectionIntentItemType(input string) (*ProtectionIntentItemType, error) {
	vals := map[string]ProtectionIntentItemType{
		"azureresourceitem":                          ProtectionIntentItemTypeAzureResourceItem,
		"azureworkloadautoprotectionintent":          ProtectionIntentItemTypeAzureWorkloadAutoProtectionIntent,
		"azureworkloadcontainerautoprotectionintent": ProtectionIntentItemTypeAzureWorkloadContainerAutoProtectionIntent,
		"azureworkloadsqlautoprotectionintent":       ProtectionIntentItemTypeAzureWorkloadSQLAutoProtectionIntent,
		"invalid":                                    ProtectionIntentItemTypeInvalid,
		"recoveryservicevaultitem":                   ProtectionIntentItemTypeRecoveryServiceVaultItem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionIntentItemType(input)
	return &out, nil
}

type ProtectionStatus string

const (
	ProtectionStatusInvalid          ProtectionStatus = "Invalid"
	ProtectionStatusNotProtected     ProtectionStatus = "NotProtected"
	ProtectionStatusProtected        ProtectionStatus = "Protected"
	ProtectionStatusProtecting       ProtectionStatus = "Protecting"
	ProtectionStatusProtectionFailed ProtectionStatus = "ProtectionFailed"
)

func PossibleValuesForProtectionStatus() []string {
	return []string{
		string(ProtectionStatusInvalid),
		string(ProtectionStatusNotProtected),
		string(ProtectionStatusProtected),
		string(ProtectionStatusProtecting),
		string(ProtectionStatusProtectionFailed),
	}
}

func parseProtectionStatus(input string) (*ProtectionStatus, error) {
	vals := map[string]ProtectionStatus{
		"invalid":          ProtectionStatusInvalid,
		"notprotected":     ProtectionStatusNotProtected,
		"protected":        ProtectionStatusProtected,
		"protecting":       ProtectionStatusProtecting,
		"protectionfailed": ProtectionStatusProtectionFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionStatus(input)
	return &out, nil
}

type WorkloadItemType string

const (
	WorkloadItemTypeInvalid           WorkloadItemType = "Invalid"
	WorkloadItemTypeSAPAseDatabase    WorkloadItemType = "SAPAseDatabase"
	WorkloadItemTypeSAPAseSystem      WorkloadItemType = "SAPAseSystem"
	WorkloadItemTypeSAPHanaDBInstance WorkloadItemType = "SAPHanaDBInstance"
	WorkloadItemTypeSAPHanaDatabase   WorkloadItemType = "SAPHanaDatabase"
	WorkloadItemTypeSAPHanaSystem     WorkloadItemType = "SAPHanaSystem"
	WorkloadItemTypeSQLDataBase       WorkloadItemType = "SQLDataBase"
	WorkloadItemTypeSQLInstance       WorkloadItemType = "SQLInstance"
)

func PossibleValuesForWorkloadItemType() []string {
	return []string{
		string(WorkloadItemTypeInvalid),
		string(WorkloadItemTypeSAPAseDatabase),
		string(WorkloadItemTypeSAPAseSystem),
		string(WorkloadItemTypeSAPHanaDBInstance),
		string(WorkloadItemTypeSAPHanaDatabase),
		string(WorkloadItemTypeSAPHanaSystem),
		string(WorkloadItemTypeSQLDataBase),
		string(WorkloadItemTypeSQLInstance),
	}
}

func parseWorkloadItemType(input string) (*WorkloadItemType, error) {
	vals := map[string]WorkloadItemType{
		"invalid":           WorkloadItemTypeInvalid,
		"sapasedatabase":    WorkloadItemTypeSAPAseDatabase,
		"sapasesystem":      WorkloadItemTypeSAPAseSystem,
		"saphanadbinstance": WorkloadItemTypeSAPHanaDBInstance,
		"saphanadatabase":   WorkloadItemTypeSAPHanaDatabase,
		"saphanasystem":     WorkloadItemTypeSAPHanaSystem,
		"sqldatabase":       WorkloadItemTypeSQLDataBase,
		"sqlinstance":       WorkloadItemTypeSQLInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadItemType(input)
	return &out, nil
}
