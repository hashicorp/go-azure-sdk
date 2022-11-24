package replicationmigrationitems

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskAccountType string

const (
	DiskAccountTypePremiumLRS     DiskAccountType = "Premium_LRS"
	DiskAccountTypeStandardLRS    DiskAccountType = "Standard_LRS"
	DiskAccountTypeStandardSSDLRS DiskAccountType = "StandardSSD_LRS"
)

func PossibleValuesForDiskAccountType() []string {
	return []string{
		string(DiskAccountTypePremiumLRS),
		string(DiskAccountTypeStandardLRS),
		string(DiskAccountTypeStandardSSDLRS),
	}
}

func parseDiskAccountType(input string) (*DiskAccountType, error) {
	vals := map[string]DiskAccountType{
		"premium_lrs":     DiskAccountTypePremiumLRS,
		"standard_lrs":    DiskAccountTypeStandardLRS,
		"standardssd_lrs": DiskAccountTypeStandardSSDLRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiskAccountType(input)
	return &out, nil
}

type EthernetAddressType string

const (
	EthernetAddressTypeDynamic EthernetAddressType = "Dynamic"
	EthernetAddressTypeStatic  EthernetAddressType = "Static"
)

func PossibleValuesForEthernetAddressType() []string {
	return []string{
		string(EthernetAddressTypeDynamic),
		string(EthernetAddressTypeStatic),
	}
}

func parseEthernetAddressType(input string) (*EthernetAddressType, error) {
	vals := map[string]EthernetAddressType{
		"dynamic": EthernetAddressTypeDynamic,
		"static":  EthernetAddressTypeStatic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EthernetAddressType(input)
	return &out, nil
}

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}

func parseHealthErrorCustomerResolvability(input string) (*HealthErrorCustomerResolvability, error) {
	vals := map[string]HealthErrorCustomerResolvability{
		"allowed":    HealthErrorCustomerResolvabilityAllowed,
		"notallowed": HealthErrorCustomerResolvabilityNotAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HealthErrorCustomerResolvability(input)
	return &out, nil
}

type LicenseType string

const (
	LicenseTypeNoLicenseType LicenseType = "NoLicenseType"
	LicenseTypeNotSpecified  LicenseType = "NotSpecified"
	LicenseTypeWindowsServer LicenseType = "WindowsServer"
)

func PossibleValuesForLicenseType() []string {
	return []string{
		string(LicenseTypeNoLicenseType),
		string(LicenseTypeNotSpecified),
		string(LicenseTypeWindowsServer),
	}
}

func parseLicenseType(input string) (*LicenseType, error) {
	vals := map[string]LicenseType{
		"nolicensetype": LicenseTypeNoLicenseType,
		"notspecified":  LicenseTypeNotSpecified,
		"windowsserver": LicenseTypeWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LicenseType(input)
	return &out, nil
}

type MigrationItemOperation string

const (
	MigrationItemOperationDisableMigration   MigrationItemOperation = "DisableMigration"
	MigrationItemOperationMigrate            MigrationItemOperation = "Migrate"
	MigrationItemOperationPauseReplication   MigrationItemOperation = "PauseReplication"
	MigrationItemOperationResumeReplication  MigrationItemOperation = "ResumeReplication"
	MigrationItemOperationStartResync        MigrationItemOperation = "StartResync"
	MigrationItemOperationTestMigrate        MigrationItemOperation = "TestMigrate"
	MigrationItemOperationTestMigrateCleanup MigrationItemOperation = "TestMigrateCleanup"
)

func PossibleValuesForMigrationItemOperation() []string {
	return []string{
		string(MigrationItemOperationDisableMigration),
		string(MigrationItemOperationMigrate),
		string(MigrationItemOperationPauseReplication),
		string(MigrationItemOperationResumeReplication),
		string(MigrationItemOperationStartResync),
		string(MigrationItemOperationTestMigrate),
		string(MigrationItemOperationTestMigrateCleanup),
	}
}

func parseMigrationItemOperation(input string) (*MigrationItemOperation, error) {
	vals := map[string]MigrationItemOperation{
		"disablemigration":   MigrationItemOperationDisableMigration,
		"migrate":            MigrationItemOperationMigrate,
		"pausereplication":   MigrationItemOperationPauseReplication,
		"resumereplication":  MigrationItemOperationResumeReplication,
		"startresync":        MigrationItemOperationStartResync,
		"testmigrate":        MigrationItemOperationTestMigrate,
		"testmigratecleanup": MigrationItemOperationTestMigrateCleanup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationItemOperation(input)
	return &out, nil
}

type MigrationState string

const (
	MigrationStateDisableMigrationFailed            MigrationState = "DisableMigrationFailed"
	MigrationStateDisableMigrationInProgress        MigrationState = "DisableMigrationInProgress"
	MigrationStateEnableMigrationFailed             MigrationState = "EnableMigrationFailed"
	MigrationStateEnableMigrationInProgress         MigrationState = "EnableMigrationInProgress"
	MigrationStateInitialSeedingFailed              MigrationState = "InitialSeedingFailed"
	MigrationStateInitialSeedingInProgress          MigrationState = "InitialSeedingInProgress"
	MigrationStateMigrationCompletedWithInformation MigrationState = "MigrationCompletedWithInformation"
	MigrationStateMigrationFailed                   MigrationState = "MigrationFailed"
	MigrationStateMigrationInProgress               MigrationState = "MigrationInProgress"
	MigrationStateMigrationPartiallySucceeded       MigrationState = "MigrationPartiallySucceeded"
	MigrationStateMigrationSucceeded                MigrationState = "MigrationSucceeded"
	MigrationStateNone                              MigrationState = "None"
	MigrationStateProtectionSuspended               MigrationState = "ProtectionSuspended"
	MigrationStateReplicating                       MigrationState = "Replicating"
	MigrationStateResumeInProgress                  MigrationState = "ResumeInProgress"
	MigrationStateResumeInitiated                   MigrationState = "ResumeInitiated"
	MigrationStateSuspendingProtection              MigrationState = "SuspendingProtection"
)

func PossibleValuesForMigrationState() []string {
	return []string{
		string(MigrationStateDisableMigrationFailed),
		string(MigrationStateDisableMigrationInProgress),
		string(MigrationStateEnableMigrationFailed),
		string(MigrationStateEnableMigrationInProgress),
		string(MigrationStateInitialSeedingFailed),
		string(MigrationStateInitialSeedingInProgress),
		string(MigrationStateMigrationCompletedWithInformation),
		string(MigrationStateMigrationFailed),
		string(MigrationStateMigrationInProgress),
		string(MigrationStateMigrationPartiallySucceeded),
		string(MigrationStateMigrationSucceeded),
		string(MigrationStateNone),
		string(MigrationStateProtectionSuspended),
		string(MigrationStateReplicating),
		string(MigrationStateResumeInProgress),
		string(MigrationStateResumeInitiated),
		string(MigrationStateSuspendingProtection),
	}
}

func parseMigrationState(input string) (*MigrationState, error) {
	vals := map[string]MigrationState{
		"disablemigrationfailed":            MigrationStateDisableMigrationFailed,
		"disablemigrationinprogress":        MigrationStateDisableMigrationInProgress,
		"enablemigrationfailed":             MigrationStateEnableMigrationFailed,
		"enablemigrationinprogress":         MigrationStateEnableMigrationInProgress,
		"initialseedingfailed":              MigrationStateInitialSeedingFailed,
		"initialseedinginprogress":          MigrationStateInitialSeedingInProgress,
		"migrationcompletedwithinformation": MigrationStateMigrationCompletedWithInformation,
		"migrationfailed":                   MigrationStateMigrationFailed,
		"migrationinprogress":               MigrationStateMigrationInProgress,
		"migrationpartiallysucceeded":       MigrationStateMigrationPartiallySucceeded,
		"migrationsucceeded":                MigrationStateMigrationSucceeded,
		"none":                              MigrationStateNone,
		"protectionsuspended":               MigrationStateProtectionSuspended,
		"replicating":                       MigrationStateReplicating,
		"resumeinprogress":                  MigrationStateResumeInProgress,
		"resumeinitiated":                   MigrationStateResumeInitiated,
		"suspendingprotection":              MigrationStateSuspendingProtection,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MigrationState(input)
	return &out, nil
}

type ProtectionHealth string

const (
	ProtectionHealthCritical ProtectionHealth = "Critical"
	ProtectionHealthNone     ProtectionHealth = "None"
	ProtectionHealthNormal   ProtectionHealth = "Normal"
	ProtectionHealthWarning  ProtectionHealth = "Warning"
)

func PossibleValuesForProtectionHealth() []string {
	return []string{
		string(ProtectionHealthCritical),
		string(ProtectionHealthNone),
		string(ProtectionHealthNormal),
		string(ProtectionHealthWarning),
	}
}

func parseProtectionHealth(input string) (*ProtectionHealth, error) {
	vals := map[string]ProtectionHealth{
		"critical": ProtectionHealthCritical,
		"none":     ProtectionHealthNone,
		"normal":   ProtectionHealthNormal,
		"warning":  ProtectionHealthWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionHealth(input)
	return &out, nil
}

type ResyncState string

const (
	ResyncStateNone                         ResyncState = "None"
	ResyncStatePreparedForResynchronization ResyncState = "PreparedForResynchronization"
	ResyncStateStartedResynchronization     ResyncState = "StartedResynchronization"
)

func PossibleValuesForResyncState() []string {
	return []string{
		string(ResyncStateNone),
		string(ResyncStatePreparedForResynchronization),
		string(ResyncStateStartedResynchronization),
	}
}

func parseResyncState(input string) (*ResyncState, error) {
	vals := map[string]ResyncState{
		"none":                         ResyncStateNone,
		"preparedforresynchronization": ResyncStatePreparedForResynchronization,
		"startedresynchronization":     ResyncStateStartedResynchronization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResyncState(input)
	return &out, nil
}

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeAHUB          SqlServerLicenseType = "AHUB"
	SqlServerLicenseTypeNoLicenseType SqlServerLicenseType = "NoLicenseType"
	SqlServerLicenseTypeNotSpecified  SqlServerLicenseType = "NotSpecified"
	SqlServerLicenseTypePAYG          SqlServerLicenseType = "PAYG"
)

func PossibleValuesForSqlServerLicenseType() []string {
	return []string{
		string(SqlServerLicenseTypeAHUB),
		string(SqlServerLicenseTypeNoLicenseType),
		string(SqlServerLicenseTypeNotSpecified),
		string(SqlServerLicenseTypePAYG),
	}
}

func parseSqlServerLicenseType(input string) (*SqlServerLicenseType, error) {
	vals := map[string]SqlServerLicenseType{
		"ahub":          SqlServerLicenseTypeAHUB,
		"nolicensetype": SqlServerLicenseTypeNoLicenseType,
		"notspecified":  SqlServerLicenseTypeNotSpecified,
		"payg":          SqlServerLicenseTypePAYG,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlServerLicenseType(input)
	return &out, nil
}

type TestMigrationState string

const (
	TestMigrationStateNone                                  TestMigrationState = "None"
	TestMigrationStateTestMigrationCleanupInProgress        TestMigrationState = "TestMigrationCleanupInProgress"
	TestMigrationStateTestMigrationCompletedWithInformation TestMigrationState = "TestMigrationCompletedWithInformation"
	TestMigrationStateTestMigrationFailed                   TestMigrationState = "TestMigrationFailed"
	TestMigrationStateTestMigrationInProgress               TestMigrationState = "TestMigrationInProgress"
	TestMigrationStateTestMigrationPartiallySucceeded       TestMigrationState = "TestMigrationPartiallySucceeded"
	TestMigrationStateTestMigrationSucceeded                TestMigrationState = "TestMigrationSucceeded"
)

func PossibleValuesForTestMigrationState() []string {
	return []string{
		string(TestMigrationStateNone),
		string(TestMigrationStateTestMigrationCleanupInProgress),
		string(TestMigrationStateTestMigrationCompletedWithInformation),
		string(TestMigrationStateTestMigrationFailed),
		string(TestMigrationStateTestMigrationInProgress),
		string(TestMigrationStateTestMigrationPartiallySucceeded),
		string(TestMigrationStateTestMigrationSucceeded),
	}
}

func parseTestMigrationState(input string) (*TestMigrationState, error) {
	vals := map[string]TestMigrationState{
		"none":                                  TestMigrationStateNone,
		"testmigrationcleanupinprogress":        TestMigrationStateTestMigrationCleanupInProgress,
		"testmigrationcompletedwithinformation": TestMigrationStateTestMigrationCompletedWithInformation,
		"testmigrationfailed":                   TestMigrationStateTestMigrationFailed,
		"testmigrationinprogress":               TestMigrationStateTestMigrationInProgress,
		"testmigrationpartiallysucceeded":       TestMigrationStateTestMigrationPartiallySucceeded,
		"testmigrationsucceeded":                TestMigrationStateTestMigrationSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TestMigrationState(input)
	return &out, nil
}
