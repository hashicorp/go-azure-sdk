package autonomousdatabases

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutonomousDatabaseBaseProperties = AutonomousDatabaseProperties{}

type AutonomousDatabaseProperties struct {

	// Fields inherited from AutonomousDatabaseBaseProperties

	ActualUsedDataStorageSizeInTbs           *float64                              `json:"actualUsedDataStorageSizeInTbs,omitempty"`
	AdminPassword                            *string                               `json:"adminPassword,omitempty"`
	AllocatedStorageSizeInTbs                *float64                              `json:"allocatedStorageSizeInTbs,omitempty"`
	ApexDetails                              *ApexDetailsType                      `json:"apexDetails,omitempty"`
	AutonomousDatabaseId                     *string                               `json:"autonomousDatabaseId,omitempty"`
	AutonomousMaintenanceScheduleType        *AutonomousMaintenanceScheduleType    `json:"autonomousMaintenanceScheduleType,omitempty"`
	AvailableUpgradeVersions                 *[]string                             `json:"availableUpgradeVersions,omitempty"`
	BackupRetentionPeriodInDays              *int64                                `json:"backupRetentionPeriodInDays,omitempty"`
	CharacterSet                             *string                               `json:"characterSet,omitempty"`
	ComputeCount                             *float64                              `json:"computeCount,omitempty"`
	ComputeModel                             *ComputeModel                         `json:"computeModel,omitempty"`
	ConnectionStrings                        *ConnectionStringType                 `json:"connectionStrings,omitempty"`
	ConnectionURLs                           *ConnectionURLType                    `json:"connectionUrls,omitempty"`
	CpuCoreCount                             *int64                                `json:"cpuCoreCount,omitempty"`
	CustomerContacts                         *[]CustomerContact                    `json:"customerContacts,omitempty"`
	DataBaseType                             DataBaseType                          `json:"dataBaseType"`
	DataSafeStatus                           *DataSafeStatusType                   `json:"dataSafeStatus,omitempty"`
	DataStorageSizeInGbs                     *int64                                `json:"dataStorageSizeInGbs,omitempty"`
	DataStorageSizeInTbs                     *int64                                `json:"dataStorageSizeInTbs,omitempty"`
	DatabaseEdition                          *DatabaseEditionType                  `json:"databaseEdition,omitempty"`
	DbVersion                                *string                               `json:"dbVersion,omitempty"`
	DbWorkload                               *WorkloadType                         `json:"dbWorkload,omitempty"`
	DisplayName                              *string                               `json:"displayName,omitempty"`
	FailedDataRecoveryInSeconds              *int64                                `json:"failedDataRecoveryInSeconds,omitempty"`
	InMemoryAreaInGbs                        *int64                                `json:"inMemoryAreaInGbs,omitempty"`
	IsAutoScalingEnabled                     *bool                                 `json:"isAutoScalingEnabled,omitempty"`
	IsAutoScalingForStorageEnabled           *bool                                 `json:"isAutoScalingForStorageEnabled,omitempty"`
	IsLocalDataGuardEnabled                  *bool                                 `json:"isLocalDataGuardEnabled,omitempty"`
	IsMtlsConnectionRequired                 *bool                                 `json:"isMtlsConnectionRequired,omitempty"`
	IsPreview                                *bool                                 `json:"isPreview,omitempty"`
	IsPreviewVersionWithServiceTermsAccepted *bool                                 `json:"isPreviewVersionWithServiceTermsAccepted,omitempty"`
	IsRemoteDataGuardEnabled                 *bool                                 `json:"isRemoteDataGuardEnabled,omitempty"`
	LicenseModel                             *LicenseModel                         `json:"licenseModel,omitempty"`
	LifecycleDetails                         *string                               `json:"lifecycleDetails,omitempty"`
	LifecycleState                           *AutonomousDatabaseLifecycleState     `json:"lifecycleState,omitempty"`
	LocalAdgAutoFailoverMaxDataLossLimit     *int64                                `json:"localAdgAutoFailoverMaxDataLossLimit,omitempty"`
	LocalDisasterRecoveryType                *DisasterRecoveryType                 `json:"localDisasterRecoveryType,omitempty"`
	LocalStandbyDb                           *AutonomousDatabaseStandbySummary     `json:"localStandbyDb,omitempty"`
	LongTermBackupSchedule                   *LongTermBackUpScheduleDetails        `json:"longTermBackupSchedule,omitempty"`
	MemoryPerOracleComputeUnitInGbs          *int64                                `json:"memoryPerOracleComputeUnitInGbs,omitempty"`
	NcharacterSet                            *string                               `json:"ncharacterSet,omitempty"`
	NextLongTermBackupTimeStamp              *string                               `json:"nextLongTermBackupTimeStamp,omitempty"`
	OciURL                                   *string                               `json:"ociUrl,omitempty"`
	Ocid                                     *string                               `json:"ocid,omitempty"`
	OpenMode                                 *OpenModeType                         `json:"openMode,omitempty"`
	OperationsInsightsStatus                 *OperationsInsightsStatusType         `json:"operationsInsightsStatus,omitempty"`
	PeerDbId                                 *string                               `json:"peerDbId,omitempty"`
	PeerDbIds                                *[]string                             `json:"peerDbIds,omitempty"`
	PermissionLevel                          *PermissionLevelType                  `json:"permissionLevel,omitempty"`
	PrivateEndpoint                          *string                               `json:"privateEndpoint,omitempty"`
	PrivateEndpointIP                        *string                               `json:"privateEndpointIp,omitempty"`
	PrivateEndpointLabel                     *string                               `json:"privateEndpointLabel,omitempty"`
	ProvisionableCPUs                        *[]int64                              `json:"provisionableCpus,omitempty"`
	ProvisioningState                        *AzureResourceProvisioningState       `json:"provisioningState,omitempty"`
	RemoteDisasterRecoveryConfiguration      *DisasterRecoveryConfigurationDetails `json:"remoteDisasterRecoveryConfiguration,omitempty"`
	Role                                     *RoleType                             `json:"role,omitempty"`
	ScheduledOperations                      *ScheduledOperationsType              `json:"scheduledOperations,omitempty"`
	ServiceConsoleURL                        *string                               `json:"serviceConsoleUrl,omitempty"`
	SqlWebDeveloperURL                       *string                               `json:"sqlWebDeveloperUrl,omitempty"`
	SubnetId                                 *string                               `json:"subnetId,omitempty"`
	SupportedRegionsToCloneTo                *[]string                             `json:"supportedRegionsToCloneTo,omitempty"`
	TimeCreated                              *string                               `json:"timeCreated,omitempty"`
	TimeDataGuardRoleChanged                 *string                               `json:"timeDataGuardRoleChanged,omitempty"`
	TimeDeletionOfFreeAutonomousDatabase     *string                               `json:"timeDeletionOfFreeAutonomousDatabase,omitempty"`
	TimeDisasterRecoveryRoleChanged          *string                               `json:"timeDisasterRecoveryRoleChanged,omitempty"`
	TimeLocalDataGuardEnabled                *string                               `json:"timeLocalDataGuardEnabled,omitempty"`
	TimeMaintenanceBegin                     *string                               `json:"timeMaintenanceBegin,omitempty"`
	TimeMaintenanceEnd                       *string                               `json:"timeMaintenanceEnd,omitempty"`
	TimeOfLastFailover                       *string                               `json:"timeOfLastFailover,omitempty"`
	TimeOfLastRefresh                        *string                               `json:"timeOfLastRefresh,omitempty"`
	TimeOfLastRefreshPoint                   *string                               `json:"timeOfLastRefreshPoint,omitempty"`
	TimeOfLastSwitchover                     *string                               `json:"timeOfLastSwitchover,omitempty"`
	TimeReclamationOfFreeAutonomousDatabase  *string                               `json:"timeReclamationOfFreeAutonomousDatabase,omitempty"`
	UsedDataStorageSizeInGbs                 *int64                                `json:"usedDataStorageSizeInGbs,omitempty"`
	UsedDataStorageSizeInTbs                 *int64                                `json:"usedDataStorageSizeInTbs,omitempty"`
	VnetId                                   *string                               `json:"vnetId,omitempty"`
	WhitelistedIPs                           *[]string                             `json:"whitelistedIps,omitempty"`
}

func (s AutonomousDatabaseProperties) AutonomousDatabaseBaseProperties() BaseAutonomousDatabaseBasePropertiesImpl {
	return BaseAutonomousDatabaseBasePropertiesImpl{
		ActualUsedDataStorageSizeInTbs:           s.ActualUsedDataStorageSizeInTbs,
		AdminPassword:                            s.AdminPassword,
		AllocatedStorageSizeInTbs:                s.AllocatedStorageSizeInTbs,
		ApexDetails:                              s.ApexDetails,
		AutonomousDatabaseId:                     s.AutonomousDatabaseId,
		AutonomousMaintenanceScheduleType:        s.AutonomousMaintenanceScheduleType,
		AvailableUpgradeVersions:                 s.AvailableUpgradeVersions,
		BackupRetentionPeriodInDays:              s.BackupRetentionPeriodInDays,
		CharacterSet:                             s.CharacterSet,
		ComputeCount:                             s.ComputeCount,
		ComputeModel:                             s.ComputeModel,
		ConnectionStrings:                        s.ConnectionStrings,
		ConnectionURLs:                           s.ConnectionURLs,
		CpuCoreCount:                             s.CpuCoreCount,
		CustomerContacts:                         s.CustomerContacts,
		DataBaseType:                             s.DataBaseType,
		DataSafeStatus:                           s.DataSafeStatus,
		DataStorageSizeInGbs:                     s.DataStorageSizeInGbs,
		DataStorageSizeInTbs:                     s.DataStorageSizeInTbs,
		DatabaseEdition:                          s.DatabaseEdition,
		DbVersion:                                s.DbVersion,
		DbWorkload:                               s.DbWorkload,
		DisplayName:                              s.DisplayName,
		FailedDataRecoveryInSeconds:              s.FailedDataRecoveryInSeconds,
		InMemoryAreaInGbs:                        s.InMemoryAreaInGbs,
		IsAutoScalingEnabled:                     s.IsAutoScalingEnabled,
		IsAutoScalingForStorageEnabled:           s.IsAutoScalingForStorageEnabled,
		IsLocalDataGuardEnabled:                  s.IsLocalDataGuardEnabled,
		IsMtlsConnectionRequired:                 s.IsMtlsConnectionRequired,
		IsPreview:                                s.IsPreview,
		IsPreviewVersionWithServiceTermsAccepted: s.IsPreviewVersionWithServiceTermsAccepted,
		IsRemoteDataGuardEnabled:                 s.IsRemoteDataGuardEnabled,
		LicenseModel:                             s.LicenseModel,
		LifecycleDetails:                         s.LifecycleDetails,
		LifecycleState:                           s.LifecycleState,
		LocalAdgAutoFailoverMaxDataLossLimit:     s.LocalAdgAutoFailoverMaxDataLossLimit,
		LocalDisasterRecoveryType:                s.LocalDisasterRecoveryType,
		LocalStandbyDb:                           s.LocalStandbyDb,
		LongTermBackupSchedule:                   s.LongTermBackupSchedule,
		MemoryPerOracleComputeUnitInGbs:          s.MemoryPerOracleComputeUnitInGbs,
		NcharacterSet:                            s.NcharacterSet,
		NextLongTermBackupTimeStamp:              s.NextLongTermBackupTimeStamp,
		OciURL:                                   s.OciURL,
		Ocid:                                     s.Ocid,
		OpenMode:                                 s.OpenMode,
		OperationsInsightsStatus:                 s.OperationsInsightsStatus,
		PeerDbId:                                 s.PeerDbId,
		PeerDbIds:                                s.PeerDbIds,
		PermissionLevel:                          s.PermissionLevel,
		PrivateEndpoint:                          s.PrivateEndpoint,
		PrivateEndpointIP:                        s.PrivateEndpointIP,
		PrivateEndpointLabel:                     s.PrivateEndpointLabel,
		ProvisionableCPUs:                        s.ProvisionableCPUs,
		ProvisioningState:                        s.ProvisioningState,
		RemoteDisasterRecoveryConfiguration:      s.RemoteDisasterRecoveryConfiguration,
		Role:                                     s.Role,
		ScheduledOperations:                      s.ScheduledOperations,
		ServiceConsoleURL:                        s.ServiceConsoleURL,
		SqlWebDeveloperURL:                       s.SqlWebDeveloperURL,
		SubnetId:                                 s.SubnetId,
		SupportedRegionsToCloneTo:                s.SupportedRegionsToCloneTo,
		TimeCreated:                              s.TimeCreated,
		TimeDataGuardRoleChanged:                 s.TimeDataGuardRoleChanged,
		TimeDeletionOfFreeAutonomousDatabase:     s.TimeDeletionOfFreeAutonomousDatabase,
		TimeDisasterRecoveryRoleChanged:          s.TimeDisasterRecoveryRoleChanged,
		TimeLocalDataGuardEnabled:                s.TimeLocalDataGuardEnabled,
		TimeMaintenanceBegin:                     s.TimeMaintenanceBegin,
		TimeMaintenanceEnd:                       s.TimeMaintenanceEnd,
		TimeOfLastFailover:                       s.TimeOfLastFailover,
		TimeOfLastRefresh:                        s.TimeOfLastRefresh,
		TimeOfLastRefreshPoint:                   s.TimeOfLastRefreshPoint,
		TimeOfLastSwitchover:                     s.TimeOfLastSwitchover,
		TimeReclamationOfFreeAutonomousDatabase:  s.TimeReclamationOfFreeAutonomousDatabase,
		UsedDataStorageSizeInGbs:                 s.UsedDataStorageSizeInGbs,
		UsedDataStorageSizeInTbs:                 s.UsedDataStorageSizeInTbs,
		VnetId:                                   s.VnetId,
		WhitelistedIPs:                           s.WhitelistedIPs,
	}
}

func (o *AutonomousDatabaseProperties) GetNextLongTermBackupTimeStampAsTime() (*time.Time, error) {
	if o.NextLongTermBackupTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NextLongTermBackupTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *AutonomousDatabaseProperties) SetNextLongTermBackupTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NextLongTermBackupTimeStamp = &formatted
}

func (o *AutonomousDatabaseProperties) GetTimeCreatedAsTime() (*time.Time, error) {
	if o.TimeCreated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeCreated, "2006-01-02T15:04:05Z07:00")
}

func (o *AutonomousDatabaseProperties) SetTimeCreatedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeCreated = &formatted
}

func (o *AutonomousDatabaseProperties) GetTimeDisasterRecoveryRoleChangedAsTime() (*time.Time, error) {
	if o.TimeDisasterRecoveryRoleChanged == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeDisasterRecoveryRoleChanged, "2006-01-02T15:04:05Z07:00")
}

func (o *AutonomousDatabaseProperties) SetTimeDisasterRecoveryRoleChangedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeDisasterRecoveryRoleChanged = &formatted
}

func (o *AutonomousDatabaseProperties) GetTimeMaintenanceBeginAsTime() (*time.Time, error) {
	if o.TimeMaintenanceBegin == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeMaintenanceBegin, "2006-01-02T15:04:05Z07:00")
}

func (o *AutonomousDatabaseProperties) SetTimeMaintenanceBeginAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeMaintenanceBegin = &formatted
}

func (o *AutonomousDatabaseProperties) GetTimeMaintenanceEndAsTime() (*time.Time, error) {
	if o.TimeMaintenanceEnd == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeMaintenanceEnd, "2006-01-02T15:04:05Z07:00")
}

func (o *AutonomousDatabaseProperties) SetTimeMaintenanceEndAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeMaintenanceEnd = &formatted
}

var _ json.Marshaler = AutonomousDatabaseProperties{}

func (s AutonomousDatabaseProperties) MarshalJSON() ([]byte, error) {
	type wrapper AutonomousDatabaseProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AutonomousDatabaseProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AutonomousDatabaseProperties: %+v", err)
	}

	decoded["dataBaseType"] = "Regular"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AutonomousDatabaseProperties: %+v", err)
	}

	return encoded, nil
}
