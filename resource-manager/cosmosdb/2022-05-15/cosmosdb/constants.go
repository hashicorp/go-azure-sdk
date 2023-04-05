package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalyticalStorageSchemaType string

const (
	AnalyticalStorageSchemaTypeFullFidelity AnalyticalStorageSchemaType = "FullFidelity"
	AnalyticalStorageSchemaTypeWellDefined  AnalyticalStorageSchemaType = "WellDefined"
)

func PossibleValuesForAnalyticalStorageSchemaType() []string {
	return []string{
		string(AnalyticalStorageSchemaTypeFullFidelity),
		string(AnalyticalStorageSchemaTypeWellDefined),
	}
}

type BackupPolicyMigrationStatus string

const (
	BackupPolicyMigrationStatusCompleted  BackupPolicyMigrationStatus = "Completed"
	BackupPolicyMigrationStatusFailed     BackupPolicyMigrationStatus = "Failed"
	BackupPolicyMigrationStatusInProgress BackupPolicyMigrationStatus = "InProgress"
	BackupPolicyMigrationStatusInvalid    BackupPolicyMigrationStatus = "Invalid"
)

func PossibleValuesForBackupPolicyMigrationStatus() []string {
	return []string{
		string(BackupPolicyMigrationStatusCompleted),
		string(BackupPolicyMigrationStatusFailed),
		string(BackupPolicyMigrationStatusInProgress),
		string(BackupPolicyMigrationStatusInvalid),
	}
}

type BackupPolicyType string

const (
	BackupPolicyTypeContinuous BackupPolicyType = "Continuous"
	BackupPolicyTypePeriodic   BackupPolicyType = "Periodic"
)

func PossibleValuesForBackupPolicyType() []string {
	return []string{
		string(BackupPolicyTypeContinuous),
		string(BackupPolicyTypePeriodic),
	}
}

type BackupStorageRedundancy string

const (
	BackupStorageRedundancyGeo   BackupStorageRedundancy = "Geo"
	BackupStorageRedundancyLocal BackupStorageRedundancy = "Local"
	BackupStorageRedundancyZone  BackupStorageRedundancy = "Zone"
)

func PossibleValuesForBackupStorageRedundancy() []string {
	return []string{
		string(BackupStorageRedundancyGeo),
		string(BackupStorageRedundancyLocal),
		string(BackupStorageRedundancyZone),
	}
}

type CompositePathSortOrder string

const (
	CompositePathSortOrderAscending  CompositePathSortOrder = "ascending"
	CompositePathSortOrderDescending CompositePathSortOrder = "descending"
)

func PossibleValuesForCompositePathSortOrder() []string {
	return []string{
		string(CompositePathSortOrderAscending),
		string(CompositePathSortOrderDescending),
	}
}

type ConflictResolutionMode string

const (
	ConflictResolutionModeCustom         ConflictResolutionMode = "Custom"
	ConflictResolutionModeLastWriterWins ConflictResolutionMode = "LastWriterWins"
)

func PossibleValuesForConflictResolutionMode() []string {
	return []string{
		string(ConflictResolutionModeCustom),
		string(ConflictResolutionModeLastWriterWins),
	}
}

type ConnectorOffer string

const (
	ConnectorOfferSmall ConnectorOffer = "Small"
)

func PossibleValuesForConnectorOffer() []string {
	return []string{
		string(ConnectorOfferSmall),
	}
}

type CreateMode string

const (
	CreateModeDefault CreateMode = "Default"
	CreateModeRestore CreateMode = "Restore"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeRestore),
	}
}

type DataType string

const (
	DataTypeLineString   DataType = "LineString"
	DataTypeMultiPolygon DataType = "MultiPolygon"
	DataTypeNumber       DataType = "Number"
	DataTypePoint        DataType = "Point"
	DataTypePolygon      DataType = "Polygon"
	DataTypeString       DataType = "String"
)

func PossibleValuesForDataType() []string {
	return []string{
		string(DataTypeLineString),
		string(DataTypeMultiPolygon),
		string(DataTypeNumber),
		string(DataTypePoint),
		string(DataTypePolygon),
		string(DataTypeString),
	}
}

type DatabaseAccountKind string

const (
	DatabaseAccountKindGlobalDocumentDB DatabaseAccountKind = "GlobalDocumentDB"
	DatabaseAccountKindMongoDB          DatabaseAccountKind = "MongoDB"
	DatabaseAccountKindParse            DatabaseAccountKind = "Parse"
)

func PossibleValuesForDatabaseAccountKind() []string {
	return []string{
		string(DatabaseAccountKindGlobalDocumentDB),
		string(DatabaseAccountKindMongoDB),
		string(DatabaseAccountKindParse),
	}
}

type DatabaseAccountOfferType string

const (
	DatabaseAccountOfferTypeStandard DatabaseAccountOfferType = "Standard"
)

func PossibleValuesForDatabaseAccountOfferType() []string {
	return []string{
		string(DatabaseAccountOfferTypeStandard),
	}
}

type DefaultConsistencyLevel string

const (
	DefaultConsistencyLevelBoundedStaleness DefaultConsistencyLevel = "BoundedStaleness"
	DefaultConsistencyLevelConsistentPrefix DefaultConsistencyLevel = "ConsistentPrefix"
	DefaultConsistencyLevelEventual         DefaultConsistencyLevel = "Eventual"
	DefaultConsistencyLevelSession          DefaultConsistencyLevel = "Session"
	DefaultConsistencyLevelStrong           DefaultConsistencyLevel = "Strong"
)

func PossibleValuesForDefaultConsistencyLevel() []string {
	return []string{
		string(DefaultConsistencyLevelBoundedStaleness),
		string(DefaultConsistencyLevelConsistentPrefix),
		string(DefaultConsistencyLevelEventual),
		string(DefaultConsistencyLevelSession),
		string(DefaultConsistencyLevelStrong),
	}
}

type IndexKind string

const (
	IndexKindHash    IndexKind = "Hash"
	IndexKindRange   IndexKind = "Range"
	IndexKindSpatial IndexKind = "Spatial"
)

func PossibleValuesForIndexKind() []string {
	return []string{
		string(IndexKindHash),
		string(IndexKindRange),
		string(IndexKindSpatial),
	}
}

type IndexingMode string

const (
	IndexingModeConsistent IndexingMode = "consistent"
	IndexingModeLazy       IndexingMode = "lazy"
	IndexingModeNone       IndexingMode = "none"
)

func PossibleValuesForIndexingMode() []string {
	return []string{
		string(IndexingModeConsistent),
		string(IndexingModeLazy),
		string(IndexingModeNone),
	}
}

type KeyKind string

const (
	KeyKindPrimary           KeyKind = "primary"
	KeyKindPrimaryReadonly   KeyKind = "primaryReadonly"
	KeyKindSecondary         KeyKind = "secondary"
	KeyKindSecondaryReadonly KeyKind = "secondaryReadonly"
)

func PossibleValuesForKeyKind() []string {
	return []string{
		string(KeyKindPrimary),
		string(KeyKindPrimaryReadonly),
		string(KeyKindSecondary),
		string(KeyKindSecondaryReadonly),
	}
}

type NetworkAclBypass string

const (
	NetworkAclBypassAzureServices NetworkAclBypass = "AzureServices"
	NetworkAclBypassNone          NetworkAclBypass = "None"
)

func PossibleValuesForNetworkAclBypass() []string {
	return []string{
		string(NetworkAclBypassAzureServices),
		string(NetworkAclBypassNone),
	}
}

type PartitionKind string

const (
	PartitionKindHash      PartitionKind = "Hash"
	PartitionKindMultiHash PartitionKind = "MultiHash"
	PartitionKindRange     PartitionKind = "Range"
)

func PossibleValuesForPartitionKind() []string {
	return []string{
		string(PartitionKindHash),
		string(PartitionKindMultiHash),
		string(PartitionKindRange),
	}
}

type PrimaryAggregationType string

const (
	PrimaryAggregationTypeAverage PrimaryAggregationType = "Average"
	PrimaryAggregationTypeLast    PrimaryAggregationType = "Last"
	PrimaryAggregationTypeMaximum PrimaryAggregationType = "Maximum"
	PrimaryAggregationTypeMinimum PrimaryAggregationType = "Minimum"
	PrimaryAggregationTypeNone    PrimaryAggregationType = "None"
	PrimaryAggregationTypeTotal   PrimaryAggregationType = "Total"
)

func PossibleValuesForPrimaryAggregationType() []string {
	return []string{
		string(PrimaryAggregationTypeAverage),
		string(PrimaryAggregationTypeLast),
		string(PrimaryAggregationTypeMaximum),
		string(PrimaryAggregationTypeMinimum),
		string(PrimaryAggregationTypeNone),
		string(PrimaryAggregationTypeTotal),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type RestoreMode string

const (
	RestoreModePointInTime RestoreMode = "PointInTime"
)

func PossibleValuesForRestoreMode() []string {
	return []string{
		string(RestoreModePointInTime),
	}
}

type ServerVersion string

const (
	ServerVersionFourPointTwo  ServerVersion = "4.2"
	ServerVersionFourPointZero ServerVersion = "4.0"
	ServerVersionThreePointSix ServerVersion = "3.6"
	ServerVersionThreePointTwo ServerVersion = "3.2"
)

func PossibleValuesForServerVersion() []string {
	return []string{
		string(ServerVersionFourPointTwo),
		string(ServerVersionFourPointZero),
		string(ServerVersionThreePointSix),
		string(ServerVersionThreePointTwo),
	}
}

type SpatialType string

const (
	SpatialTypeLineString   SpatialType = "LineString"
	SpatialTypeMultiPolygon SpatialType = "MultiPolygon"
	SpatialTypePoint        SpatialType = "Point"
	SpatialTypePolygon      SpatialType = "Polygon"
)

func PossibleValuesForSpatialType() []string {
	return []string{
		string(SpatialTypeLineString),
		string(SpatialTypeMultiPolygon),
		string(SpatialTypePoint),
		string(SpatialTypePolygon),
	}
}

type TriggerOperation string

const (
	TriggerOperationAll     TriggerOperation = "All"
	TriggerOperationCreate  TriggerOperation = "Create"
	TriggerOperationDelete  TriggerOperation = "Delete"
	TriggerOperationReplace TriggerOperation = "Replace"
	TriggerOperationUpdate  TriggerOperation = "Update"
)

func PossibleValuesForTriggerOperation() []string {
	return []string{
		string(TriggerOperationAll),
		string(TriggerOperationCreate),
		string(TriggerOperationDelete),
		string(TriggerOperationReplace),
		string(TriggerOperationUpdate),
	}
}

type TriggerType string

const (
	TriggerTypePost TriggerType = "Post"
	TriggerTypePre  TriggerType = "Pre"
)

func PossibleValuesForTriggerType() []string {
	return []string{
		string(TriggerTypePost),
		string(TriggerTypePre),
	}
}

type UnitType string

const (
	UnitTypeBytes          UnitType = "Bytes"
	UnitTypeBytesPerSecond UnitType = "BytesPerSecond"
	UnitTypeCount          UnitType = "Count"
	UnitTypeCountPerSecond UnitType = "CountPerSecond"
	UnitTypeMilliseconds   UnitType = "Milliseconds"
	UnitTypePercent        UnitType = "Percent"
	UnitTypeSeconds        UnitType = "Seconds"
)

func PossibleValuesForUnitType() []string {
	return []string{
		string(UnitTypeBytes),
		string(UnitTypeBytesPerSecond),
		string(UnitTypeCount),
		string(UnitTypeCountPerSecond),
		string(UnitTypeMilliseconds),
		string(UnitTypePercent),
		string(UnitTypeSeconds),
	}
}
