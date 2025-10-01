package bms

import (
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AcquireStorageAccountLock string

const (
	AcquireStorageAccountLockAcquire    AcquireStorageAccountLock = "Acquire"
	AcquireStorageAccountLockNotAcquire AcquireStorageAccountLock = "NotAcquire"
)

func PossibleValuesForAcquireStorageAccountLock() []string {
	return []string{
		string(AcquireStorageAccountLockAcquire),
		string(AcquireStorageAccountLockNotAcquire),
	}
}

func parseAcquireStorageAccountLock(input string) (*AcquireStorageAccountLock, error) {
	vals := map[string]AcquireStorageAccountLock{
		"acquire":    AcquireStorageAccountLockAcquire,
		"notacquire": AcquireStorageAccountLockNotAcquire,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AcquireStorageAccountLock(input)
	return &out, nil
}

type AzureFileShareType string

const (
	AzureFileShareTypeInvalid AzureFileShareType = "Invalid"
	AzureFileShareTypeXSMB    AzureFileShareType = "XSMB"
	AzureFileShareTypeXSync   AzureFileShareType = "XSync"
)

func PossibleValuesForAzureFileShareType() []string {
	return []string{
		string(AzureFileShareTypeInvalid),
		string(AzureFileShareTypeXSMB),
		string(AzureFileShareTypeXSync),
	}
}

func parseAzureFileShareType(input string) (*AzureFileShareType, error) {
	vals := map[string]AzureFileShareType{
		"invalid": AzureFileShareTypeInvalid,
		"xsmb":    AzureFileShareTypeXSMB,
		"xsync":   AzureFileShareTypeXSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AzureFileShareType(input)
	return &out, nil
}

type BackupItemType string

const (
	BackupItemTypeAzureFileShare    BackupItemType = "AzureFileShare"
	BackupItemTypeAzureSqlDb        BackupItemType = "AzureSqlDb"
	BackupItemTypeClient            BackupItemType = "Client"
	BackupItemTypeExchange          BackupItemType = "Exchange"
	BackupItemTypeFileFolder        BackupItemType = "FileFolder"
	BackupItemTypeGenericDataSource BackupItemType = "GenericDataSource"
	BackupItemTypeInvalid           BackupItemType = "Invalid"
	BackupItemTypeSAPAseDatabase    BackupItemType = "SAPAseDatabase"
	BackupItemTypeSAPHanaDBInstance BackupItemType = "SAPHanaDBInstance"
	BackupItemTypeSAPHanaDatabase   BackupItemType = "SAPHanaDatabase"
	BackupItemTypeSQLDB             BackupItemType = "SQLDB"
	BackupItemTypeSQLDataBase       BackupItemType = "SQLDataBase"
	BackupItemTypeSharepoint        BackupItemType = "Sharepoint"
	BackupItemTypeSystemState       BackupItemType = "SystemState"
	BackupItemTypeVM                BackupItemType = "VM"
	BackupItemTypeVMwareVM          BackupItemType = "VMwareVM"
)

func PossibleValuesForBackupItemType() []string {
	return []string{
		string(BackupItemTypeAzureFileShare),
		string(BackupItemTypeAzureSqlDb),
		string(BackupItemTypeClient),
		string(BackupItemTypeExchange),
		string(BackupItemTypeFileFolder),
		string(BackupItemTypeGenericDataSource),
		string(BackupItemTypeInvalid),
		string(BackupItemTypeSAPAseDatabase),
		string(BackupItemTypeSAPHanaDBInstance),
		string(BackupItemTypeSAPHanaDatabase),
		string(BackupItemTypeSQLDB),
		string(BackupItemTypeSQLDataBase),
		string(BackupItemTypeSharepoint),
		string(BackupItemTypeSystemState),
		string(BackupItemTypeVM),
		string(BackupItemTypeVMwareVM),
	}
}

func parseBackupItemType(input string) (*BackupItemType, error) {
	vals := map[string]BackupItemType{
		"azurefileshare":    BackupItemTypeAzureFileShare,
		"azuresqldb":        BackupItemTypeAzureSqlDb,
		"client":            BackupItemTypeClient,
		"exchange":          BackupItemTypeExchange,
		"filefolder":        BackupItemTypeFileFolder,
		"genericdatasource": BackupItemTypeGenericDataSource,
		"invalid":           BackupItemTypeInvalid,
		"sapasedatabase":    BackupItemTypeSAPAseDatabase,
		"saphanadbinstance": BackupItemTypeSAPHanaDBInstance,
		"saphanadatabase":   BackupItemTypeSAPHanaDatabase,
		"sqldb":             BackupItemTypeSQLDB,
		"sqldatabase":       BackupItemTypeSQLDataBase,
		"sharepoint":        BackupItemTypeSharepoint,
		"systemstate":       BackupItemTypeSystemState,
		"vm":                BackupItemTypeVM,
		"vmwarevm":          BackupItemTypeVMwareVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BackupItemType(input)
	return &out, nil
}

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

type CopyOptions string

const (
	CopyOptionsCreateCopy     CopyOptions = "CreateCopy"
	CopyOptionsFailOnConflict CopyOptions = "FailOnConflict"
	CopyOptionsInvalid        CopyOptions = "Invalid"
	CopyOptionsOverwrite      CopyOptions = "Overwrite"
	CopyOptionsSkip           CopyOptions = "Skip"
)

func PossibleValuesForCopyOptions() []string {
	return []string{
		string(CopyOptionsCreateCopy),
		string(CopyOptionsFailOnConflict),
		string(CopyOptionsInvalid),
		string(CopyOptionsOverwrite),
		string(CopyOptionsSkip),
	}
}

func parseCopyOptions(input string) (*CopyOptions, error) {
	vals := map[string]CopyOptions{
		"createcopy":     CopyOptionsCreateCopy,
		"failonconflict": CopyOptionsFailOnConflict,
		"invalid":        CopyOptionsInvalid,
		"overwrite":      CopyOptionsOverwrite,
		"skip":           CopyOptionsSkip,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CopyOptions(input)
	return &out, nil
}

type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeInvalid CreateMode = "Invalid"
	CreateModeRecover CreateMode = "Recover"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeInvalid),
		string(CreateModeRecover),
	}
}

func parseCreateMode(input string) (*CreateMode, error) {
	vals := map[string]CreateMode{
		"default": CreateModeDefault,
		"invalid": CreateModeInvalid,
		"recover": CreateModeRecover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMode(input)
	return &out, nil
}

type DataSourceType string

const (
	DataSourceTypeAzureFileShare    DataSourceType = "AzureFileShare"
	DataSourceTypeAzureSqlDb        DataSourceType = "AzureSqlDb"
	DataSourceTypeClient            DataSourceType = "Client"
	DataSourceTypeExchange          DataSourceType = "Exchange"
	DataSourceTypeFileFolder        DataSourceType = "FileFolder"
	DataSourceTypeGenericDataSource DataSourceType = "GenericDataSource"
	DataSourceTypeInvalid           DataSourceType = "Invalid"
	DataSourceTypeSAPAseDatabase    DataSourceType = "SAPAseDatabase"
	DataSourceTypeSAPHanaDBInstance DataSourceType = "SAPHanaDBInstance"
	DataSourceTypeSAPHanaDatabase   DataSourceType = "SAPHanaDatabase"
	DataSourceTypeSQLDB             DataSourceType = "SQLDB"
	DataSourceTypeSQLDataBase       DataSourceType = "SQLDataBase"
	DataSourceTypeSharepoint        DataSourceType = "Sharepoint"
	DataSourceTypeSystemState       DataSourceType = "SystemState"
	DataSourceTypeVM                DataSourceType = "VM"
	DataSourceTypeVMwareVM          DataSourceType = "VMwareVM"
)

func PossibleValuesForDataSourceType() []string {
	return []string{
		string(DataSourceTypeAzureFileShare),
		string(DataSourceTypeAzureSqlDb),
		string(DataSourceTypeClient),
		string(DataSourceTypeExchange),
		string(DataSourceTypeFileFolder),
		string(DataSourceTypeGenericDataSource),
		string(DataSourceTypeInvalid),
		string(DataSourceTypeSAPAseDatabase),
		string(DataSourceTypeSAPHanaDBInstance),
		string(DataSourceTypeSAPHanaDatabase),
		string(DataSourceTypeSQLDB),
		string(DataSourceTypeSQLDataBase),
		string(DataSourceTypeSharepoint),
		string(DataSourceTypeSystemState),
		string(DataSourceTypeVM),
		string(DataSourceTypeVMwareVM),
	}
}

func parseDataSourceType(input string) (*DataSourceType, error) {
	vals := map[string]DataSourceType{
		"azurefileshare":    DataSourceTypeAzureFileShare,
		"azuresqldb":        DataSourceTypeAzureSqlDb,
		"client":            DataSourceTypeClient,
		"exchange":          DataSourceTypeExchange,
		"filefolder":        DataSourceTypeFileFolder,
		"genericdatasource": DataSourceTypeGenericDataSource,
		"invalid":           DataSourceTypeInvalid,
		"sapasedatabase":    DataSourceTypeSAPAseDatabase,
		"saphanadbinstance": DataSourceTypeSAPHanaDBInstance,
		"saphanadatabase":   DataSourceTypeSAPHanaDatabase,
		"sqldb":             DataSourceTypeSQLDB,
		"sqldatabase":       DataSourceTypeSQLDataBase,
		"sharepoint":        DataSourceTypeSharepoint,
		"systemstate":       DataSourceTypeSystemState,
		"vm":                DataSourceTypeVM,
		"vmwarevm":          DataSourceTypeVMwareVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DataSourceType(input)
	return &out, nil
}

type FabricName string

const (
	FabricNameAzure   FabricName = "Azure"
	FabricNameInvalid FabricName = "Invalid"
)

func PossibleValuesForFabricName() []string {
	return []string{
		string(FabricNameAzure),
		string(FabricNameInvalid),
	}
}

func parseFabricName(input string) (*FabricName, error) {
	vals := map[string]FabricName{
		"azure":   FabricNameAzure,
		"invalid": FabricNameInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FabricName(input)
	return &out, nil
}

type HealthStatus string

const (
	HealthStatusActionRequired  HealthStatus = "ActionRequired"
	HealthStatusActionSuggested HealthStatus = "ActionSuggested"
	HealthStatusInvalid         HealthStatus = "Invalid"
	HealthStatusPassed          HealthStatus = "Passed"
)

func PossibleValuesForHealthStatus() []string {
	return []string{
		string(HealthStatusActionRequired),
		string(HealthStatusActionSuggested),
		string(HealthStatusInvalid),
		string(HealthStatusPassed),
	}
}

func parseHealthStatus(input string) (*HealthStatus, error) {
	vals := map[string]HealthStatus{
		"actionrequired":  HealthStatusActionRequired,
		"actionsuggested": HealthStatusActionSuggested,
		"invalid":         HealthStatusInvalid,
		"passed":          HealthStatusPassed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthStatus(input)
	return &out, nil
}

type InquiryStatus string

const (
	InquiryStatusFailed  InquiryStatus = "Failed"
	InquiryStatusInvalid InquiryStatus = "Invalid"
	InquiryStatusSuccess InquiryStatus = "Success"
)

func PossibleValuesForInquiryStatus() []string {
	return []string{
		string(InquiryStatusFailed),
		string(InquiryStatusInvalid),
		string(InquiryStatusSuccess),
	}
}

func parseInquiryStatus(input string) (*InquiryStatus, error) {
	vals := map[string]InquiryStatus{
		"failed":  InquiryStatusFailed,
		"invalid": InquiryStatusInvalid,
		"success": InquiryStatusSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InquiryStatus(input)
	return &out, nil
}

type LastBackupStatus string

const (
	LastBackupStatusHealthy   LastBackupStatus = "Healthy"
	LastBackupStatusIRPending LastBackupStatus = "IRPending"
	LastBackupStatusInvalid   LastBackupStatus = "Invalid"
	LastBackupStatusUnhealthy LastBackupStatus = "Unhealthy"
)

func PossibleValuesForLastBackupStatus() []string {
	return []string{
		string(LastBackupStatusHealthy),
		string(LastBackupStatusIRPending),
		string(LastBackupStatusInvalid),
		string(LastBackupStatusUnhealthy),
	}
}

func parseLastBackupStatus(input string) (*LastBackupStatus, error) {
	vals := map[string]LastBackupStatus{
		"healthy":   LastBackupStatusHealthy,
		"irpending": LastBackupStatusIRPending,
		"invalid":   LastBackupStatusInvalid,
		"unhealthy": LastBackupStatusUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LastBackupStatus(input)
	return &out, nil
}

type OperationStatusValues string

const (
	OperationStatusValuesCanceled   OperationStatusValues = "Canceled"
	OperationStatusValuesFailed     OperationStatusValues = "Failed"
	OperationStatusValuesInProgress OperationStatusValues = "InProgress"
	OperationStatusValuesInvalid    OperationStatusValues = "Invalid"
	OperationStatusValuesSucceeded  OperationStatusValues = "Succeeded"
)

func PossibleValuesForOperationStatusValues() []string {
	return []string{
		string(OperationStatusValuesCanceled),
		string(OperationStatusValuesFailed),
		string(OperationStatusValuesInProgress),
		string(OperationStatusValuesInvalid),
		string(OperationStatusValuesSucceeded),
	}
}

func parseOperationStatusValues(input string) (*OperationStatusValues, error) {
	vals := map[string]OperationStatusValues{
		"canceled":   OperationStatusValuesCanceled,
		"failed":     OperationStatusValuesFailed,
		"inprogress": OperationStatusValuesInProgress,
		"invalid":    OperationStatusValuesInvalid,
		"succeeded":  OperationStatusValuesSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatusValues(input)
	return &out, nil
}

type OperationType string

const (
	OperationTypeInvalid    OperationType = "Invalid"
	OperationTypeRegister   OperationType = "Register"
	OperationTypeRehydrate  OperationType = "Rehydrate"
	OperationTypeReregister OperationType = "Reregister"
)

func PossibleValuesForOperationType() []string {
	return []string{
		string(OperationTypeInvalid),
		string(OperationTypeRegister),
		string(OperationTypeRehydrate),
		string(OperationTypeReregister),
	}
}

func parseOperationType(input string) (*OperationType, error) {
	vals := map[string]OperationType{
		"invalid":    OperationTypeInvalid,
		"register":   OperationTypeRegister,
		"rehydrate":  OperationTypeRehydrate,
		"reregister": OperationTypeReregister,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationType(input)
	return &out, nil
}

type OverwriteOptions string

const (
	OverwriteOptionsFailOnConflict OverwriteOptions = "FailOnConflict"
	OverwriteOptionsInvalid        OverwriteOptions = "Invalid"
	OverwriteOptionsOverwrite      OverwriteOptions = "Overwrite"
)

func PossibleValuesForOverwriteOptions() []string {
	return []string{
		string(OverwriteOptionsFailOnConflict),
		string(OverwriteOptionsInvalid),
		string(OverwriteOptionsOverwrite),
	}
}

func parseOverwriteOptions(input string) (*OverwriteOptions, error) {
	vals := map[string]OverwriteOptions{
		"failonconflict": OverwriteOptionsFailOnConflict,
		"invalid":        OverwriteOptionsInvalid,
		"overwrite":      OverwriteOptionsOverwrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OverwriteOptions(input)
	return &out, nil
}

type ProtectableContainerType string

const (
	ProtectableContainerTypeAzureBackupServerContainer                  ProtectableContainerType = "AzureBackupServerContainer"
	ProtectableContainerTypeAzureSqlContainer                           ProtectableContainerType = "AzureSqlContainer"
	ProtectableContainerTypeAzureWorkloadContainer                      ProtectableContainerType = "AzureWorkloadContainer"
	ProtectableContainerTypeCluster                                     ProtectableContainerType = "Cluster"
	ProtectableContainerTypeDPMContainer                                ProtectableContainerType = "DPMContainer"
	ProtectableContainerTypeGenericContainer                            ProtectableContainerType = "GenericContainer"
	ProtectableContainerTypeIaasVMContainer                             ProtectableContainerType = "IaasVMContainer"
	ProtectableContainerTypeIaasVMServiceContainer                      ProtectableContainerType = "IaasVMServiceContainer"
	ProtectableContainerTypeInvalid                                     ProtectableContainerType = "Invalid"
	ProtectableContainerTypeMABContainer                                ProtectableContainerType = "MABContainer"
	ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines ProtectableContainerType = "Microsoft.ClassicCompute/virtualMachines"
	ProtectableContainerTypeMicrosoftPointComputeVirtualMachines        ProtectableContainerType = "Microsoft.Compute/virtualMachines"
	ProtectableContainerTypeSQLAGWorkLoadContainer                      ProtectableContainerType = "SQLAGWorkLoadContainer"
	ProtectableContainerTypeStorageContainer                            ProtectableContainerType = "StorageContainer"
	ProtectableContainerTypeUnknown                                     ProtectableContainerType = "Unknown"
	ProtectableContainerTypeVCenter                                     ProtectableContainerType = "VCenter"
	ProtectableContainerTypeVMAppContainer                              ProtectableContainerType = "VMAppContainer"
	ProtectableContainerTypeWindows                                     ProtectableContainerType = "Windows"
)

func PossibleValuesForProtectableContainerType() []string {
	return []string{
		string(ProtectableContainerTypeAzureBackupServerContainer),
		string(ProtectableContainerTypeAzureSqlContainer),
		string(ProtectableContainerTypeAzureWorkloadContainer),
		string(ProtectableContainerTypeCluster),
		string(ProtectableContainerTypeDPMContainer),
		string(ProtectableContainerTypeGenericContainer),
		string(ProtectableContainerTypeIaasVMContainer),
		string(ProtectableContainerTypeIaasVMServiceContainer),
		string(ProtectableContainerTypeInvalid),
		string(ProtectableContainerTypeMABContainer),
		string(ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines),
		string(ProtectableContainerTypeMicrosoftPointComputeVirtualMachines),
		string(ProtectableContainerTypeSQLAGWorkLoadContainer),
		string(ProtectableContainerTypeStorageContainer),
		string(ProtectableContainerTypeUnknown),
		string(ProtectableContainerTypeVCenter),
		string(ProtectableContainerTypeVMAppContainer),
		string(ProtectableContainerTypeWindows),
	}
}

func parseProtectableContainerType(input string) (*ProtectableContainerType, error) {
	vals := map[string]ProtectableContainerType{
		"azurebackupservercontainer": ProtectableContainerTypeAzureBackupServerContainer,
		"azuresqlcontainer":          ProtectableContainerTypeAzureSqlContainer,
		"azureworkloadcontainer":     ProtectableContainerTypeAzureWorkloadContainer,
		"cluster":                    ProtectableContainerTypeCluster,
		"dpmcontainer":               ProtectableContainerTypeDPMContainer,
		"genericcontainer":           ProtectableContainerTypeGenericContainer,
		"iaasvmcontainer":            ProtectableContainerTypeIaasVMContainer,
		"iaasvmservicecontainer":     ProtectableContainerTypeIaasVMServiceContainer,
		"invalid":                    ProtectableContainerTypeInvalid,
		"mabcontainer":               ProtectableContainerTypeMABContainer,
		"microsoft.classiccompute/virtualmachines": ProtectableContainerTypeMicrosoftPointClassicComputeVirtualMachines,
		"microsoft.compute/virtualmachines":        ProtectableContainerTypeMicrosoftPointComputeVirtualMachines,
		"sqlagworkloadcontainer":                   ProtectableContainerTypeSQLAGWorkLoadContainer,
		"storagecontainer":                         ProtectableContainerTypeStorageContainer,
		"unknown":                                  ProtectableContainerTypeUnknown,
		"vcenter":                                  ProtectableContainerTypeVCenter,
		"vmappcontainer":                           ProtectableContainerTypeVMAppContainer,
		"windows":                                  ProtectableContainerTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectableContainerType(input)
	return &out, nil
}

type ProtectedItemHealthStatus string

const (
	ProtectedItemHealthStatusHealthy      ProtectedItemHealthStatus = "Healthy"
	ProtectedItemHealthStatusIRPending    ProtectedItemHealthStatus = "IRPending"
	ProtectedItemHealthStatusInvalid      ProtectedItemHealthStatus = "Invalid"
	ProtectedItemHealthStatusNotReachable ProtectedItemHealthStatus = "NotReachable"
	ProtectedItemHealthStatusUnhealthy    ProtectedItemHealthStatus = "Unhealthy"
)

func PossibleValuesForProtectedItemHealthStatus() []string {
	return []string{
		string(ProtectedItemHealthStatusHealthy),
		string(ProtectedItemHealthStatusIRPending),
		string(ProtectedItemHealthStatusInvalid),
		string(ProtectedItemHealthStatusNotReachable),
		string(ProtectedItemHealthStatusUnhealthy),
	}
}

func parseProtectedItemHealthStatus(input string) (*ProtectedItemHealthStatus, error) {
	vals := map[string]ProtectedItemHealthStatus{
		"healthy":      ProtectedItemHealthStatusHealthy,
		"irpending":    ProtectedItemHealthStatusIRPending,
		"invalid":      ProtectedItemHealthStatusInvalid,
		"notreachable": ProtectedItemHealthStatusNotReachable,
		"unhealthy":    ProtectedItemHealthStatusUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectedItemHealthStatus(input)
	return &out, nil
}

type ProtectedItemState string

const (
	ProtectedItemStateBackupsSuspended  ProtectedItemState = "BackupsSuspended"
	ProtectedItemStateIRPending         ProtectedItemState = "IRPending"
	ProtectedItemStateInvalid           ProtectedItemState = "Invalid"
	ProtectedItemStateProtected         ProtectedItemState = "Protected"
	ProtectedItemStateProtectionError   ProtectedItemState = "ProtectionError"
	ProtectedItemStateProtectionPaused  ProtectedItemState = "ProtectionPaused"
	ProtectedItemStateProtectionStopped ProtectedItemState = "ProtectionStopped"
)

func PossibleValuesForProtectedItemState() []string {
	return []string{
		string(ProtectedItemStateBackupsSuspended),
		string(ProtectedItemStateIRPending),
		string(ProtectedItemStateInvalid),
		string(ProtectedItemStateProtected),
		string(ProtectedItemStateProtectionError),
		string(ProtectedItemStateProtectionPaused),
		string(ProtectedItemStateProtectionStopped),
	}
}

func parseProtectedItemState(input string) (*ProtectedItemState, error) {
	vals := map[string]ProtectedItemState{
		"backupssuspended":  ProtectedItemStateBackupsSuspended,
		"irpending":         ProtectedItemStateIRPending,
		"invalid":           ProtectedItemStateInvalid,
		"protected":         ProtectedItemStateProtected,
		"protectionerror":   ProtectedItemStateProtectionError,
		"protectionpaused":  ProtectedItemStateProtectionPaused,
		"protectionstopped": ProtectedItemStateProtectionStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectedItemState(input)
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

type ProtectionState string

const (
	ProtectionStateBackupsSuspended  ProtectionState = "BackupsSuspended"
	ProtectionStateIRPending         ProtectionState = "IRPending"
	ProtectionStateInvalid           ProtectionState = "Invalid"
	ProtectionStateProtected         ProtectionState = "Protected"
	ProtectionStateProtectionError   ProtectionState = "ProtectionError"
	ProtectionStateProtectionPaused  ProtectionState = "ProtectionPaused"
	ProtectionStateProtectionStopped ProtectionState = "ProtectionStopped"
)

func PossibleValuesForProtectionState() []string {
	return []string{
		string(ProtectionStateBackupsSuspended),
		string(ProtectionStateIRPending),
		string(ProtectionStateInvalid),
		string(ProtectionStateProtected),
		string(ProtectionStateProtectionError),
		string(ProtectionStateProtectionPaused),
		string(ProtectionStateProtectionStopped),
	}
}

func parseProtectionState(input string) (*ProtectionState, error) {
	vals := map[string]ProtectionState{
		"backupssuspended":  ProtectionStateBackupsSuspended,
		"irpending":         ProtectionStateIRPending,
		"invalid":           ProtectionStateInvalid,
		"protected":         ProtectionStateProtected,
		"protectionerror":   ProtectionStateProtectionError,
		"protectionpaused":  ProtectionStateProtectionPaused,
		"protectionstopped": ProtectionStateProtectionStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionState(input)
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

type RecoveryMode string

const (
	RecoveryModeFileRecovery             RecoveryMode = "FileRecovery"
	RecoveryModeInvalid                  RecoveryMode = "Invalid"
	RecoveryModeRecoveryUsingSnapshot    RecoveryMode = "RecoveryUsingSnapshot"
	RecoveryModeSnapshotAttach           RecoveryMode = "SnapshotAttach"
	RecoveryModeSnapshotAttachAndRecover RecoveryMode = "SnapshotAttachAndRecover"
	RecoveryModeWorkloadRecovery         RecoveryMode = "WorkloadRecovery"
)

func PossibleValuesForRecoveryMode() []string {
	return []string{
		string(RecoveryModeFileRecovery),
		string(RecoveryModeInvalid),
		string(RecoveryModeRecoveryUsingSnapshot),
		string(RecoveryModeSnapshotAttach),
		string(RecoveryModeSnapshotAttachAndRecover),
		string(RecoveryModeWorkloadRecovery),
	}
}

func parseRecoveryMode(input string) (*RecoveryMode, error) {
	vals := map[string]RecoveryMode{
		"filerecovery":             RecoveryModeFileRecovery,
		"invalid":                  RecoveryModeInvalid,
		"recoveryusingsnapshot":    RecoveryModeRecoveryUsingSnapshot,
		"snapshotattach":           RecoveryModeSnapshotAttach,
		"snapshotattachandrecover": RecoveryModeSnapshotAttachAndRecover,
		"workloadrecovery":         RecoveryModeWorkloadRecovery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryMode(input)
	return &out, nil
}

type RecoveryPointTierType string

const (
	RecoveryPointTierTypeArchivedRP RecoveryPointTierType = "ArchivedRP"
	RecoveryPointTierTypeHardenedRP RecoveryPointTierType = "HardenedRP"
	RecoveryPointTierTypeInstantRP  RecoveryPointTierType = "InstantRP"
	RecoveryPointTierTypeInvalid    RecoveryPointTierType = "Invalid"
)

func PossibleValuesForRecoveryPointTierType() []string {
	return []string{
		string(RecoveryPointTierTypeArchivedRP),
		string(RecoveryPointTierTypeHardenedRP),
		string(RecoveryPointTierTypeInstantRP),
		string(RecoveryPointTierTypeInvalid),
	}
}

func parseRecoveryPointTierType(input string) (*RecoveryPointTierType, error) {
	vals := map[string]RecoveryPointTierType{
		"archivedrp": RecoveryPointTierTypeArchivedRP,
		"hardenedrp": RecoveryPointTierTypeHardenedRP,
		"instantrp":  RecoveryPointTierTypeInstantRP,
		"invalid":    RecoveryPointTierTypeInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryPointTierType(input)
	return &out, nil
}

type RecoveryType string

const (
	RecoveryTypeAlternateLocation RecoveryType = "AlternateLocation"
	RecoveryTypeInvalid           RecoveryType = "Invalid"
	RecoveryTypeOffline           RecoveryType = "Offline"
	RecoveryTypeOriginalLocation  RecoveryType = "OriginalLocation"
	RecoveryTypeRestoreDisks      RecoveryType = "RestoreDisks"
)

func PossibleValuesForRecoveryType() []string {
	return []string{
		string(RecoveryTypeAlternateLocation),
		string(RecoveryTypeInvalid),
		string(RecoveryTypeOffline),
		string(RecoveryTypeOriginalLocation),
		string(RecoveryTypeRestoreDisks),
	}
}

func parseRecoveryType(input string) (*RecoveryType, error) {
	vals := map[string]RecoveryType{
		"alternatelocation": RecoveryTypeAlternateLocation,
		"invalid":           RecoveryTypeInvalid,
		"offline":           RecoveryTypeOffline,
		"originallocation":  RecoveryTypeOriginalLocation,
		"restoredisks":      RecoveryTypeRestoreDisks,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecoveryType(input)
	return &out, nil
}

type RehydrationPriority string

const (
	RehydrationPriorityHigh     RehydrationPriority = "High"
	RehydrationPriorityStandard RehydrationPriority = "Standard"
)

func PossibleValuesForRehydrationPriority() []string {
	return []string{
		string(RehydrationPriorityHigh),
		string(RehydrationPriorityStandard),
	}
}

func parseRehydrationPriority(input string) (*RehydrationPriority, error) {
	vals := map[string]RehydrationPriority{
		"high":     RehydrationPriorityHigh,
		"standard": RehydrationPriorityStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RehydrationPriority(input)
	return &out, nil
}

type ResourceHealthStatus string

const (
	ResourceHealthStatusHealthy             ResourceHealthStatus = "Healthy"
	ResourceHealthStatusInvalid             ResourceHealthStatus = "Invalid"
	ResourceHealthStatusPersistentDegraded  ResourceHealthStatus = "PersistentDegraded"
	ResourceHealthStatusPersistentUnhealthy ResourceHealthStatus = "PersistentUnhealthy"
	ResourceHealthStatusTransientDegraded   ResourceHealthStatus = "TransientDegraded"
	ResourceHealthStatusTransientUnhealthy  ResourceHealthStatus = "TransientUnhealthy"
)

func PossibleValuesForResourceHealthStatus() []string {
	return []string{
		string(ResourceHealthStatusHealthy),
		string(ResourceHealthStatusInvalid),
		string(ResourceHealthStatusPersistentDegraded),
		string(ResourceHealthStatusPersistentUnhealthy),
		string(ResourceHealthStatusTransientDegraded),
		string(ResourceHealthStatusTransientUnhealthy),
	}
}

func parseResourceHealthStatus(input string) (*ResourceHealthStatus, error) {
	vals := map[string]ResourceHealthStatus{
		"healthy":             ResourceHealthStatusHealthy,
		"invalid":             ResourceHealthStatusInvalid,
		"persistentdegraded":  ResourceHealthStatusPersistentDegraded,
		"persistentunhealthy": ResourceHealthStatusPersistentUnhealthy,
		"transientdegraded":   ResourceHealthStatusTransientDegraded,
		"transientunhealthy":  ResourceHealthStatusTransientUnhealthy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceHealthStatus(input)
	return &out, nil
}

type RestoreRequestType string

const (
	RestoreRequestTypeFullShareRestore RestoreRequestType = "FullShareRestore"
	RestoreRequestTypeInvalid          RestoreRequestType = "Invalid"
	RestoreRequestTypeItemLevelRestore RestoreRequestType = "ItemLevelRestore"
)

func PossibleValuesForRestoreRequestType() []string {
	return []string{
		string(RestoreRequestTypeFullShareRestore),
		string(RestoreRequestTypeInvalid),
		string(RestoreRequestTypeItemLevelRestore),
	}
}

func parseRestoreRequestType(input string) (*RestoreRequestType, error) {
	vals := map[string]RestoreRequestType{
		"fullsharerestore": RestoreRequestTypeFullShareRestore,
		"invalid":          RestoreRequestTypeInvalid,
		"itemlevelrestore": RestoreRequestTypeItemLevelRestore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestoreRequestType(input)
	return &out, nil
}

type SQLDataDirectoryType string

const (
	SQLDataDirectoryTypeData    SQLDataDirectoryType = "Data"
	SQLDataDirectoryTypeInvalid SQLDataDirectoryType = "Invalid"
	SQLDataDirectoryTypeLog     SQLDataDirectoryType = "Log"
)

func PossibleValuesForSQLDataDirectoryType() []string {
	return []string{
		string(SQLDataDirectoryTypeData),
		string(SQLDataDirectoryTypeInvalid),
		string(SQLDataDirectoryTypeLog),
	}
}

func parseSQLDataDirectoryType(input string) (*SQLDataDirectoryType, error) {
	vals := map[string]SQLDataDirectoryType{
		"data":    SQLDataDirectoryTypeData,
		"invalid": SQLDataDirectoryTypeInvalid,
		"log":     SQLDataDirectoryTypeLog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SQLDataDirectoryType(input)
	return &out, nil
}

type SupportStatus string

const (
	SupportStatusDefaultOFF   SupportStatus = "DefaultOFF"
	SupportStatusDefaultON    SupportStatus = "DefaultON"
	SupportStatusInvalid      SupportStatus = "Invalid"
	SupportStatusNotSupported SupportStatus = "NotSupported"
	SupportStatusSupported    SupportStatus = "Supported"
)

func PossibleValuesForSupportStatus() []string {
	return []string{
		string(SupportStatusDefaultOFF),
		string(SupportStatusDefaultON),
		string(SupportStatusInvalid),
		string(SupportStatusNotSupported),
		string(SupportStatusSupported),
	}
}

func parseSupportStatus(input string) (*SupportStatus, error) {
	vals := map[string]SupportStatus{
		"defaultoff":   SupportStatusDefaultOFF,
		"defaulton":    SupportStatusDefaultON,
		"invalid":      SupportStatusInvalid,
		"notsupported": SupportStatusNotSupported,
		"supported":    SupportStatusSupported,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SupportStatus(input)
	return &out, nil
}

type TargetDiskNetworkAccessOption string

const (
	TargetDiskNetworkAccessOptionEnablePrivateAccessForAllDisks TargetDiskNetworkAccessOption = "EnablePrivateAccessForAllDisks"
	TargetDiskNetworkAccessOptionEnablePublicAccessForAllDisks  TargetDiskNetworkAccessOption = "EnablePublicAccessForAllDisks"
	TargetDiskNetworkAccessOptionSameAsOnSourceDisks            TargetDiskNetworkAccessOption = "SameAsOnSourceDisks"
)

func PossibleValuesForTargetDiskNetworkAccessOption() []string {
	return []string{
		string(TargetDiskNetworkAccessOptionEnablePrivateAccessForAllDisks),
		string(TargetDiskNetworkAccessOptionEnablePublicAccessForAllDisks),
		string(TargetDiskNetworkAccessOptionSameAsOnSourceDisks),
	}
}

func parseTargetDiskNetworkAccessOption(input string) (*TargetDiskNetworkAccessOption, error) {
	vals := map[string]TargetDiskNetworkAccessOption{
		"enableprivateaccessforalldisks": TargetDiskNetworkAccessOptionEnablePrivateAccessForAllDisks,
		"enablepublicaccessforalldisks":  TargetDiskNetworkAccessOptionEnablePublicAccessForAllDisks,
		"sameasonsourcedisks":            TargetDiskNetworkAccessOptionSameAsOnSourceDisks,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetDiskNetworkAccessOption(input)
	return &out, nil
}

type UsagesUnit string

const (
	UsagesUnitBytes          UsagesUnit = "Bytes"
	UsagesUnitBytesPerSecond UsagesUnit = "BytesPerSecond"
	UsagesUnitCount          UsagesUnit = "Count"
	UsagesUnitCountPerSecond UsagesUnit = "CountPerSecond"
	UsagesUnitPercent        UsagesUnit = "Percent"
	UsagesUnitSeconds        UsagesUnit = "Seconds"
)

func PossibleValuesForUsagesUnit() []string {
	return []string{
		string(UsagesUnitBytes),
		string(UsagesUnitBytesPerSecond),
		string(UsagesUnitCount),
		string(UsagesUnitCountPerSecond),
		string(UsagesUnitPercent),
		string(UsagesUnitSeconds),
	}
}

func parseUsagesUnit(input string) (*UsagesUnit, error) {
	vals := map[string]UsagesUnit{
		"bytes":          UsagesUnitBytes,
		"bytespersecond": UsagesUnitBytesPerSecond,
		"count":          UsagesUnitCount,
		"countpersecond": UsagesUnitCountPerSecond,
		"percent":        UsagesUnitPercent,
		"seconds":        UsagesUnitSeconds,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsagesUnit(input)
	return &out, nil
}

type ValidationStatus string

const (
	ValidationStatusFailed    ValidationStatus = "Failed"
	ValidationStatusInvalid   ValidationStatus = "Invalid"
	ValidationStatusSucceeded ValidationStatus = "Succeeded"
)

func PossibleValuesForValidationStatus() []string {
	return []string{
		string(ValidationStatusFailed),
		string(ValidationStatusInvalid),
		string(ValidationStatusSucceeded),
	}
}

func parseValidationStatus(input string) (*ValidationStatus, error) {
	vals := map[string]ValidationStatus{
		"failed":    ValidationStatusFailed,
		"invalid":   ValidationStatusInvalid,
		"succeeded": ValidationStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidationStatus(input)
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

type WorkloadType string

const (
	WorkloadTypeAzureFileShare    WorkloadType = "AzureFileShare"
	WorkloadTypeAzureSqlDb        WorkloadType = "AzureSqlDb"
	WorkloadTypeClient            WorkloadType = "Client"
	WorkloadTypeExchange          WorkloadType = "Exchange"
	WorkloadTypeFileFolder        WorkloadType = "FileFolder"
	WorkloadTypeGenericDataSource WorkloadType = "GenericDataSource"
	WorkloadTypeInvalid           WorkloadType = "Invalid"
	WorkloadTypeSAPAseDatabase    WorkloadType = "SAPAseDatabase"
	WorkloadTypeSAPHanaDBInstance WorkloadType = "SAPHanaDBInstance"
	WorkloadTypeSAPHanaDatabase   WorkloadType = "SAPHanaDatabase"
	WorkloadTypeSQLDB             WorkloadType = "SQLDB"
	WorkloadTypeSQLDataBase       WorkloadType = "SQLDataBase"
	WorkloadTypeSharepoint        WorkloadType = "Sharepoint"
	WorkloadTypeSystemState       WorkloadType = "SystemState"
	WorkloadTypeVM                WorkloadType = "VM"
	WorkloadTypeVMwareVM          WorkloadType = "VMwareVM"
)

func PossibleValuesForWorkloadType() []string {
	return []string{
		string(WorkloadTypeAzureFileShare),
		string(WorkloadTypeAzureSqlDb),
		string(WorkloadTypeClient),
		string(WorkloadTypeExchange),
		string(WorkloadTypeFileFolder),
		string(WorkloadTypeGenericDataSource),
		string(WorkloadTypeInvalid),
		string(WorkloadTypeSAPAseDatabase),
		string(WorkloadTypeSAPHanaDBInstance),
		string(WorkloadTypeSAPHanaDatabase),
		string(WorkloadTypeSQLDB),
		string(WorkloadTypeSQLDataBase),
		string(WorkloadTypeSharepoint),
		string(WorkloadTypeSystemState),
		string(WorkloadTypeVM),
		string(WorkloadTypeVMwareVM),
	}
}

func parseWorkloadType(input string) (*WorkloadType, error) {
	vals := map[string]WorkloadType{
		"azurefileshare":    WorkloadTypeAzureFileShare,
		"azuresqldb":        WorkloadTypeAzureSqlDb,
		"client":            WorkloadTypeClient,
		"exchange":          WorkloadTypeExchange,
		"filefolder":        WorkloadTypeFileFolder,
		"genericdatasource": WorkloadTypeGenericDataSource,
		"invalid":           WorkloadTypeInvalid,
		"sapasedatabase":    WorkloadTypeSAPAseDatabase,
		"saphanadbinstance": WorkloadTypeSAPHanaDBInstance,
		"saphanadatabase":   WorkloadTypeSAPHanaDatabase,
		"sqldb":             WorkloadTypeSQLDB,
		"sqldatabase":       WorkloadTypeSQLDataBase,
		"sharepoint":        WorkloadTypeSharepoint,
		"systemstate":       WorkloadTypeSystemState,
		"vm":                WorkloadTypeVM,
		"vmwarevm":          WorkloadTypeVMwareVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkloadType(input)
	return &out, nil
}
