package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetMappingKind string

const (
	DataSetMappingKindAdlsGenTwoFile       DataSetMappingKind = "AdlsGen2File"
	DataSetMappingKindAdlsGenTwoFileSystem DataSetMappingKind = "AdlsGen2FileSystem"
	DataSetMappingKindAdlsGenTwoFolder     DataSetMappingKind = "AdlsGen2Folder"
	DataSetMappingKindBlob                 DataSetMappingKind = "Blob"
	DataSetMappingKindBlobFolder           DataSetMappingKind = "BlobFolder"
	DataSetMappingKindContainer            DataSetMappingKind = "Container"
	DataSetMappingKindKustoCluster         DataSetMappingKind = "KustoCluster"
	DataSetMappingKindKustoDatabase        DataSetMappingKind = "KustoDatabase"
	DataSetMappingKindSqlDBTable           DataSetMappingKind = "SqlDBTable"
	DataSetMappingKindSqlDWTable           DataSetMappingKind = "SqlDWTable"
)

func PossibleValuesForDataSetMappingKind() []string {
	return []string{
		string(DataSetMappingKindAdlsGenTwoFile),
		string(DataSetMappingKindAdlsGenTwoFileSystem),
		string(DataSetMappingKindAdlsGenTwoFolder),
		string(DataSetMappingKindBlob),
		string(DataSetMappingKindBlobFolder),
		string(DataSetMappingKindContainer),
		string(DataSetMappingKindKustoCluster),
		string(DataSetMappingKindKustoDatabase),
		string(DataSetMappingKindSqlDBTable),
		string(DataSetMappingKindSqlDWTable),
	}
}

type DataSetMappingStatus string

const (
	DataSetMappingStatusBroken DataSetMappingStatus = "Broken"
	DataSetMappingStatusOk     DataSetMappingStatus = "Ok"
)

func PossibleValuesForDataSetMappingStatus() []string {
	return []string{
		string(DataSetMappingStatusBroken),
		string(DataSetMappingStatusOk),
	}
}

type OutputType string

const (
	OutputTypeCsv     OutputType = "Csv"
	OutputTypeParquet OutputType = "Parquet"
)

func PossibleValuesForOutputType() []string {
	return []string{
		string(OutputTypeCsv),
		string(OutputTypeParquet),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}
