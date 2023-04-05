package restorables

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiType string

const (
	ApiTypeCassandra   ApiType = "Cassandra"
	ApiTypeGremlin     ApiType = "Gremlin"
	ApiTypeGremlinVTwo ApiType = "GremlinV2"
	ApiTypeMongoDB     ApiType = "MongoDB"
	ApiTypeSql         ApiType = "Sql"
	ApiTypeTable       ApiType = "Table"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeCassandra),
		string(ApiTypeGremlin),
		string(ApiTypeGremlinVTwo),
		string(ApiTypeMongoDB),
		string(ApiTypeSql),
		string(ApiTypeTable),
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

type OperationType string

const (
	OperationTypeCreate          OperationType = "Create"
	OperationTypeDelete          OperationType = "Delete"
	OperationTypeReplace         OperationType = "Replace"
	OperationTypeSystemOperation OperationType = "SystemOperation"
)

func PossibleValuesForOperationType() []string {
	return []string{
		string(OperationTypeCreate),
		string(OperationTypeDelete),
		string(OperationTypeReplace),
		string(OperationTypeSystemOperation),
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
