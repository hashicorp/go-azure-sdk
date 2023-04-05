package sqlvirtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssessmentDayOfWeek string

const (
	AssessmentDayOfWeekFriday    AssessmentDayOfWeek = "Friday"
	AssessmentDayOfWeekMonday    AssessmentDayOfWeek = "Monday"
	AssessmentDayOfWeekSaturday  AssessmentDayOfWeek = "Saturday"
	AssessmentDayOfWeekSunday    AssessmentDayOfWeek = "Sunday"
	AssessmentDayOfWeekThursday  AssessmentDayOfWeek = "Thursday"
	AssessmentDayOfWeekTuesday   AssessmentDayOfWeek = "Tuesday"
	AssessmentDayOfWeekWednesday AssessmentDayOfWeek = "Wednesday"
)

func PossibleValuesForAssessmentDayOfWeek() []string {
	return []string{
		string(AssessmentDayOfWeekFriday),
		string(AssessmentDayOfWeekMonday),
		string(AssessmentDayOfWeekSaturday),
		string(AssessmentDayOfWeekSunday),
		string(AssessmentDayOfWeekThursday),
		string(AssessmentDayOfWeekTuesday),
		string(AssessmentDayOfWeekWednesday),
	}
}

type AutoBackupDaysOfWeek string

const (
	AutoBackupDaysOfWeekFriday    AutoBackupDaysOfWeek = "Friday"
	AutoBackupDaysOfWeekMonday    AutoBackupDaysOfWeek = "Monday"
	AutoBackupDaysOfWeekSaturday  AutoBackupDaysOfWeek = "Saturday"
	AutoBackupDaysOfWeekSunday    AutoBackupDaysOfWeek = "Sunday"
	AutoBackupDaysOfWeekThursday  AutoBackupDaysOfWeek = "Thursday"
	AutoBackupDaysOfWeekTuesday   AutoBackupDaysOfWeek = "Tuesday"
	AutoBackupDaysOfWeekWednesday AutoBackupDaysOfWeek = "Wednesday"
)

func PossibleValuesForAutoBackupDaysOfWeek() []string {
	return []string{
		string(AutoBackupDaysOfWeekFriday),
		string(AutoBackupDaysOfWeekMonday),
		string(AutoBackupDaysOfWeekSaturday),
		string(AutoBackupDaysOfWeekSunday),
		string(AutoBackupDaysOfWeekThursday),
		string(AutoBackupDaysOfWeekTuesday),
		string(AutoBackupDaysOfWeekWednesday),
	}
}

type BackupScheduleType string

const (
	BackupScheduleTypeAutomated BackupScheduleType = "Automated"
	BackupScheduleTypeManual    BackupScheduleType = "Manual"
)

func PossibleValuesForBackupScheduleType() []string {
	return []string{
		string(BackupScheduleTypeAutomated),
		string(BackupScheduleTypeManual),
	}
}

type ConnectivityType string

const (
	ConnectivityTypeLOCAL   ConnectivityType = "LOCAL"
	ConnectivityTypePRIVATE ConnectivityType = "PRIVATE"
	ConnectivityTypePUBLIC  ConnectivityType = "PUBLIC"
)

func PossibleValuesForConnectivityType() []string {
	return []string{
		string(ConnectivityTypeLOCAL),
		string(ConnectivityTypePRIVATE),
		string(ConnectivityTypePUBLIC),
	}
}

type DayOfWeek string

const (
	DayOfWeekEveryday  DayOfWeek = "Everyday"
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekEveryday),
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

type DiskConfigurationType string

const (
	DiskConfigurationTypeADD    DiskConfigurationType = "ADD"
	DiskConfigurationTypeEXTEND DiskConfigurationType = "EXTEND"
	DiskConfigurationTypeNEW    DiskConfigurationType = "NEW"
)

func PossibleValuesForDiskConfigurationType() []string {
	return []string{
		string(DiskConfigurationTypeADD),
		string(DiskConfigurationTypeEXTEND),
		string(DiskConfigurationTypeNEW),
	}
}

type FullBackupFrequencyType string

const (
	FullBackupFrequencyTypeDaily  FullBackupFrequencyType = "Daily"
	FullBackupFrequencyTypeWeekly FullBackupFrequencyType = "Weekly"
)

func PossibleValuesForFullBackupFrequencyType() []string {
	return []string{
		string(FullBackupFrequencyTypeDaily),
		string(FullBackupFrequencyTypeWeekly),
	}
}

type SqlImageSku string

const (
	SqlImageSkuDeveloper  SqlImageSku = "Developer"
	SqlImageSkuEnterprise SqlImageSku = "Enterprise"
	SqlImageSkuExpress    SqlImageSku = "Express"
	SqlImageSkuStandard   SqlImageSku = "Standard"
	SqlImageSkuWeb        SqlImageSku = "Web"
)

func PossibleValuesForSqlImageSku() []string {
	return []string{
		string(SqlImageSkuDeveloper),
		string(SqlImageSkuEnterprise),
		string(SqlImageSkuExpress),
		string(SqlImageSkuStandard),
		string(SqlImageSkuWeb),
	}
}

type SqlManagementMode string

const (
	SqlManagementModeFull        SqlManagementMode = "Full"
	SqlManagementModeLightWeight SqlManagementMode = "LightWeight"
	SqlManagementModeNoAgent     SqlManagementMode = "NoAgent"
)

func PossibleValuesForSqlManagementMode() []string {
	return []string{
		string(SqlManagementModeFull),
		string(SqlManagementModeLightWeight),
		string(SqlManagementModeNoAgent),
	}
}

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeAHUB SqlServerLicenseType = "AHUB"
	SqlServerLicenseTypeDR   SqlServerLicenseType = "DR"
	SqlServerLicenseTypePAYG SqlServerLicenseType = "PAYG"
)

func PossibleValuesForSqlServerLicenseType() []string {
	return []string{
		string(SqlServerLicenseTypeAHUB),
		string(SqlServerLicenseTypeDR),
		string(SqlServerLicenseTypePAYG),
	}
}

type SqlWorkloadType string

const (
	SqlWorkloadTypeDW      SqlWorkloadType = "DW"
	SqlWorkloadTypeGENERAL SqlWorkloadType = "GENERAL"
	SqlWorkloadTypeOLTP    SqlWorkloadType = "OLTP"
)

func PossibleValuesForSqlWorkloadType() []string {
	return []string{
		string(SqlWorkloadTypeDW),
		string(SqlWorkloadTypeGENERAL),
		string(SqlWorkloadTypeOLTP),
	}
}

type StorageWorkloadType string

const (
	StorageWorkloadTypeDW      StorageWorkloadType = "DW"
	StorageWorkloadTypeGENERAL StorageWorkloadType = "GENERAL"
	StorageWorkloadTypeOLTP    StorageWorkloadType = "OLTP"
)

func PossibleValuesForStorageWorkloadType() []string {
	return []string{
		string(StorageWorkloadTypeDW),
		string(StorageWorkloadTypeGENERAL),
		string(StorageWorkloadTypeOLTP),
	}
}
