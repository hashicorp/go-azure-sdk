package databases

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CatalogCollationType string

const (
	CatalogCollationTypeDATABASEDEFAULT             CatalogCollationType = "DATABASE_DEFAULT"
	CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS CatalogCollationType = "SQL_Latin1_General_CP1_CI_AS"
)

func PossibleValuesForCatalogCollationType() []string {
	return []string{
		string(CatalogCollationTypeDATABASEDEFAULT),
		string(CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS),
	}
}

func (s *CatalogCollationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCatalogCollationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCatalogCollationType(input string) (*CatalogCollationType, error) {
	vals := map[string]CatalogCollationType{
		"database_default":             CatalogCollationTypeDATABASEDEFAULT,
		"sql_latin1_general_cp1_ci_as": CatalogCollationTypeSQLLatinOneGeneralCPOneCIAS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CatalogCollationType(input)
	return &out, nil
}

type CreateMode string

const (
	CreateModeCopy                           CreateMode = "Copy"
	CreateModeDefault                        CreateMode = "Default"
	CreateModeOnlineSecondary                CreateMode = "OnlineSecondary"
	CreateModePointInTimeRestore             CreateMode = "PointInTimeRestore"
	CreateModeRecovery                       CreateMode = "Recovery"
	CreateModeRestore                        CreateMode = "Restore"
	CreateModeRestoreExternalBackup          CreateMode = "RestoreExternalBackup"
	CreateModeRestoreExternalBackupSecondary CreateMode = "RestoreExternalBackupSecondary"
	CreateModeRestoreLongTermRetentionBackup CreateMode = "RestoreLongTermRetentionBackup"
	CreateModeSecondary                      CreateMode = "Secondary"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeCopy),
		string(CreateModeDefault),
		string(CreateModeOnlineSecondary),
		string(CreateModePointInTimeRestore),
		string(CreateModeRecovery),
		string(CreateModeRestore),
		string(CreateModeRestoreExternalBackup),
		string(CreateModeRestoreExternalBackupSecondary),
		string(CreateModeRestoreLongTermRetentionBackup),
		string(CreateModeSecondary),
	}
}

func (s *CreateMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMode(input string) (*CreateMode, error) {
	vals := map[string]CreateMode{
		"copy":                           CreateModeCopy,
		"default":                        CreateModeDefault,
		"onlinesecondary":                CreateModeOnlineSecondary,
		"pointintimerestore":             CreateModePointInTimeRestore,
		"recovery":                       CreateModeRecovery,
		"restore":                        CreateModeRestore,
		"restoreexternalbackup":          CreateModeRestoreExternalBackup,
		"restoreexternalbackupsecondary": CreateModeRestoreExternalBackupSecondary,
		"restorelongtermretentionbackup": CreateModeRestoreLongTermRetentionBackup,
		"secondary":                      CreateModeSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMode(input)
	return &out, nil
}

type CurrentBackupStorageRedundancy string

const (
	CurrentBackupStorageRedundancyGeo   CurrentBackupStorageRedundancy = "Geo"
	CurrentBackupStorageRedundancyLocal CurrentBackupStorageRedundancy = "Local"
	CurrentBackupStorageRedundancyZone  CurrentBackupStorageRedundancy = "Zone"
)

func PossibleValuesForCurrentBackupStorageRedundancy() []string {
	return []string{
		string(CurrentBackupStorageRedundancyGeo),
		string(CurrentBackupStorageRedundancyLocal),
		string(CurrentBackupStorageRedundancyZone),
	}
}

func (s *CurrentBackupStorageRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCurrentBackupStorageRedundancy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCurrentBackupStorageRedundancy(input string) (*CurrentBackupStorageRedundancy, error) {
	vals := map[string]CurrentBackupStorageRedundancy{
		"geo":   CurrentBackupStorageRedundancyGeo,
		"local": CurrentBackupStorageRedundancyLocal,
		"zone":  CurrentBackupStorageRedundancyZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CurrentBackupStorageRedundancy(input)
	return &out, nil
}

type DatabaseLicenseType string

const (
	DatabaseLicenseTypeBasePrice       DatabaseLicenseType = "BasePrice"
	DatabaseLicenseTypeLicenseIncluded DatabaseLicenseType = "LicenseIncluded"
)

func PossibleValuesForDatabaseLicenseType() []string {
	return []string{
		string(DatabaseLicenseTypeBasePrice),
		string(DatabaseLicenseTypeLicenseIncluded),
	}
}

func (s *DatabaseLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseLicenseType(input string) (*DatabaseLicenseType, error) {
	vals := map[string]DatabaseLicenseType{
		"baseprice":       DatabaseLicenseTypeBasePrice,
		"licenseincluded": DatabaseLicenseTypeLicenseIncluded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseLicenseType(input)
	return &out, nil
}

type DatabaseReadScale string

const (
	DatabaseReadScaleDisabled DatabaseReadScale = "Disabled"
	DatabaseReadScaleEnabled  DatabaseReadScale = "Enabled"
)

func PossibleValuesForDatabaseReadScale() []string {
	return []string{
		string(DatabaseReadScaleDisabled),
		string(DatabaseReadScaleEnabled),
	}
}

func (s *DatabaseReadScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseReadScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseReadScale(input string) (*DatabaseReadScale, error) {
	vals := map[string]DatabaseReadScale{
		"disabled": DatabaseReadScaleDisabled,
		"enabled":  DatabaseReadScaleEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseReadScale(input)
	return &out, nil
}

type DatabaseStatus string

const (
	DatabaseStatusAutoClosed                        DatabaseStatus = "AutoClosed"
	DatabaseStatusCopying                           DatabaseStatus = "Copying"
	DatabaseStatusCreating                          DatabaseStatus = "Creating"
	DatabaseStatusDisabled                          DatabaseStatus = "Disabled"
	DatabaseStatusEmergencyMode                     DatabaseStatus = "EmergencyMode"
	DatabaseStatusInaccessible                      DatabaseStatus = "Inaccessible"
	DatabaseStatusOffline                           DatabaseStatus = "Offline"
	DatabaseStatusOfflineChangingDwPerformanceTiers DatabaseStatus = "OfflineChangingDwPerformanceTiers"
	DatabaseStatusOfflineSecondary                  DatabaseStatus = "OfflineSecondary"
	DatabaseStatusOnline                            DatabaseStatus = "Online"
	DatabaseStatusOnlineChangingDwPerformanceTiers  DatabaseStatus = "OnlineChangingDwPerformanceTiers"
	DatabaseStatusPaused                            DatabaseStatus = "Paused"
	DatabaseStatusPausing                           DatabaseStatus = "Pausing"
	DatabaseStatusRecovering                        DatabaseStatus = "Recovering"
	DatabaseStatusRecoveryPending                   DatabaseStatus = "RecoveryPending"
	DatabaseStatusRestoring                         DatabaseStatus = "Restoring"
	DatabaseStatusResuming                          DatabaseStatus = "Resuming"
	DatabaseStatusScaling                           DatabaseStatus = "Scaling"
	DatabaseStatusShutdown                          DatabaseStatus = "Shutdown"
	DatabaseStatusStandby                           DatabaseStatus = "Standby"
	DatabaseStatusSuspect                           DatabaseStatus = "Suspect"
)

func PossibleValuesForDatabaseStatus() []string {
	return []string{
		string(DatabaseStatusAutoClosed),
		string(DatabaseStatusCopying),
		string(DatabaseStatusCreating),
		string(DatabaseStatusDisabled),
		string(DatabaseStatusEmergencyMode),
		string(DatabaseStatusInaccessible),
		string(DatabaseStatusOffline),
		string(DatabaseStatusOfflineChangingDwPerformanceTiers),
		string(DatabaseStatusOfflineSecondary),
		string(DatabaseStatusOnline),
		string(DatabaseStatusOnlineChangingDwPerformanceTiers),
		string(DatabaseStatusPaused),
		string(DatabaseStatusPausing),
		string(DatabaseStatusRecovering),
		string(DatabaseStatusRecoveryPending),
		string(DatabaseStatusRestoring),
		string(DatabaseStatusResuming),
		string(DatabaseStatusScaling),
		string(DatabaseStatusShutdown),
		string(DatabaseStatusStandby),
		string(DatabaseStatusSuspect),
	}
}

func (s *DatabaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseStatus(input string) (*DatabaseStatus, error) {
	vals := map[string]DatabaseStatus{
		"autoclosed":                        DatabaseStatusAutoClosed,
		"copying":                           DatabaseStatusCopying,
		"creating":                          DatabaseStatusCreating,
		"disabled":                          DatabaseStatusDisabled,
		"emergencymode":                     DatabaseStatusEmergencyMode,
		"inaccessible":                      DatabaseStatusInaccessible,
		"offline":                           DatabaseStatusOffline,
		"offlinechangingdwperformancetiers": DatabaseStatusOfflineChangingDwPerformanceTiers,
		"offlinesecondary":                  DatabaseStatusOfflineSecondary,
		"online":                            DatabaseStatusOnline,
		"onlinechangingdwperformancetiers":  DatabaseStatusOnlineChangingDwPerformanceTiers,
		"paused":                            DatabaseStatusPaused,
		"pausing":                           DatabaseStatusPausing,
		"recovering":                        DatabaseStatusRecovering,
		"recoverypending":                   DatabaseStatusRecoveryPending,
		"restoring":                         DatabaseStatusRestoring,
		"resuming":                          DatabaseStatusResuming,
		"scaling":                           DatabaseStatusScaling,
		"shutdown":                          DatabaseStatusShutdown,
		"standby":                           DatabaseStatusStandby,
		"suspect":                           DatabaseStatusSuspect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseStatus(input)
	return &out, nil
}

type ReplicaType string

const (
	ReplicaTypePrimary           ReplicaType = "Primary"
	ReplicaTypeReadableSecondary ReplicaType = "ReadableSecondary"
)

func PossibleValuesForReplicaType() []string {
	return []string{
		string(ReplicaTypePrimary),
		string(ReplicaTypeReadableSecondary),
	}
}

func (s *ReplicaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicaType(input string) (*ReplicaType, error) {
	vals := map[string]ReplicaType{
		"primary":           ReplicaTypePrimary,
		"readablesecondary": ReplicaTypeReadableSecondary,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicaType(input)
	return &out, nil
}

type RequestedBackupStorageRedundancy string

const (
	RequestedBackupStorageRedundancyGeo   RequestedBackupStorageRedundancy = "Geo"
	RequestedBackupStorageRedundancyLocal RequestedBackupStorageRedundancy = "Local"
	RequestedBackupStorageRedundancyZone  RequestedBackupStorageRedundancy = "Zone"
)

func PossibleValuesForRequestedBackupStorageRedundancy() []string {
	return []string{
		string(RequestedBackupStorageRedundancyGeo),
		string(RequestedBackupStorageRedundancyLocal),
		string(RequestedBackupStorageRedundancyZone),
	}
}

func (s *RequestedBackupStorageRedundancy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRequestedBackupStorageRedundancy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRequestedBackupStorageRedundancy(input string) (*RequestedBackupStorageRedundancy, error) {
	vals := map[string]RequestedBackupStorageRedundancy{
		"geo":   RequestedBackupStorageRedundancyGeo,
		"local": RequestedBackupStorageRedundancyLocal,
		"zone":  RequestedBackupStorageRedundancyZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RequestedBackupStorageRedundancy(input)
	return &out, nil
}

type SampleName string

const (
	SampleNameAdventureWorksLT       SampleName = "AdventureWorksLT"
	SampleNameWideWorldImportersFull SampleName = "WideWorldImportersFull"
	SampleNameWideWorldImportersStd  SampleName = "WideWorldImportersStd"
)

func PossibleValuesForSampleName() []string {
	return []string{
		string(SampleNameAdventureWorksLT),
		string(SampleNameWideWorldImportersFull),
		string(SampleNameWideWorldImportersStd),
	}
}

func (s *SampleName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSampleName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSampleName(input string) (*SampleName, error) {
	vals := map[string]SampleName{
		"adventureworkslt":       SampleNameAdventureWorksLT,
		"wideworldimportersfull": SampleNameWideWorldImportersFull,
		"wideworldimportersstd":  SampleNameWideWorldImportersStd,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SampleName(input)
	return &out, nil
}

type SecondaryType string

const (
	SecondaryTypeGeo   SecondaryType = "Geo"
	SecondaryTypeNamed SecondaryType = "Named"
)

func PossibleValuesForSecondaryType() []string {
	return []string{
		string(SecondaryTypeGeo),
		string(SecondaryTypeNamed),
	}
}

func (s *SecondaryType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecondaryType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecondaryType(input string) (*SecondaryType, error) {
	vals := map[string]SecondaryType{
		"geo":   SecondaryTypeGeo,
		"named": SecondaryTypeNamed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecondaryType(input)
	return &out, nil
}

type StorageKeyType string

const (
	StorageKeyTypeSharedAccessKey  StorageKeyType = "SharedAccessKey"
	StorageKeyTypeStorageAccessKey StorageKeyType = "StorageAccessKey"
)

func PossibleValuesForStorageKeyType() []string {
	return []string{
		string(StorageKeyTypeSharedAccessKey),
		string(StorageKeyTypeStorageAccessKey),
	}
}

func (s *StorageKeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageKeyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageKeyType(input string) (*StorageKeyType, error) {
	vals := map[string]StorageKeyType{
		"sharedaccesskey":  StorageKeyTypeSharedAccessKey,
		"storageaccesskey": StorageKeyTypeStorageAccessKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageKeyType(input)
	return &out, nil
}
