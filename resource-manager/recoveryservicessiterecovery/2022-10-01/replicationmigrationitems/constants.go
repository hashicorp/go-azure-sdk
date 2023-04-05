package replicationmigrationitems

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
