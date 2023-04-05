package dataset

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetKind string

const (
	DataSetKindAdlsGenOneFile       DataSetKind = "AdlsGen1File"
	DataSetKindAdlsGenOneFolder     DataSetKind = "AdlsGen1Folder"
	DataSetKindAdlsGenTwoFile       DataSetKind = "AdlsGen2File"
	DataSetKindAdlsGenTwoFileSystem DataSetKind = "AdlsGen2FileSystem"
	DataSetKindAdlsGenTwoFolder     DataSetKind = "AdlsGen2Folder"
	DataSetKindBlob                 DataSetKind = "Blob"
	DataSetKindBlobFolder           DataSetKind = "BlobFolder"
	DataSetKindContainer            DataSetKind = "Container"
	DataSetKindKustoCluster         DataSetKind = "KustoCluster"
	DataSetKindKustoDatabase        DataSetKind = "KustoDatabase"
	DataSetKindSqlDBTable           DataSetKind = "SqlDBTable"
	DataSetKindSqlDWTable           DataSetKind = "SqlDWTable"
)

func PossibleValuesForDataSetKind() []string {
	return []string{
		string(DataSetKindAdlsGenOneFile),
		string(DataSetKindAdlsGenOneFolder),
		string(DataSetKindAdlsGenTwoFile),
		string(DataSetKindAdlsGenTwoFileSystem),
		string(DataSetKindAdlsGenTwoFolder),
		string(DataSetKindBlob),
		string(DataSetKindBlobFolder),
		string(DataSetKindContainer),
		string(DataSetKindKustoCluster),
		string(DataSetKindKustoDatabase),
		string(DataSetKindSqlDBTable),
		string(DataSetKindSqlDWTable),
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
