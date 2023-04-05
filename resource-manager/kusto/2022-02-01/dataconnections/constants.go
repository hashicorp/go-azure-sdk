package dataconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobStorageEventType string

const (
	BlobStorageEventTypeMicrosoftPointStoragePointBlobCreated BlobStorageEventType = "Microsoft.Storage.BlobCreated"
	BlobStorageEventTypeMicrosoftPointStoragePointBlobRenamed BlobStorageEventType = "Microsoft.Storage.BlobRenamed"
)

func PossibleValuesForBlobStorageEventType() []string {
	return []string{
		string(BlobStorageEventTypeMicrosoftPointStoragePointBlobCreated),
		string(BlobStorageEventTypeMicrosoftPointStoragePointBlobRenamed),
	}
}

type Compression string

const (
	CompressionGZip Compression = "GZip"
	CompressionNone Compression = "None"
)

func PossibleValuesForCompression() []string {
	return []string{
		string(CompressionGZip),
		string(CompressionNone),
	}
}

type DataConnectionKind string

const (
	DataConnectionKindEventGrid DataConnectionKind = "EventGrid"
	DataConnectionKindEventHub  DataConnectionKind = "EventHub"
	DataConnectionKindIotHub    DataConnectionKind = "IotHub"
)

func PossibleValuesForDataConnectionKind() []string {
	return []string{
		string(DataConnectionKindEventGrid),
		string(DataConnectionKindEventHub),
		string(DataConnectionKindIotHub),
	}
}

type DataConnectionType string

const (
	DataConnectionTypeMicrosoftPointKustoClustersDatabasesDataConnections DataConnectionType = "Microsoft.Kusto/clusters/databases/dataConnections"
)

func PossibleValuesForDataConnectionType() []string {
	return []string{
		string(DataConnectionTypeMicrosoftPointKustoClustersDatabasesDataConnections),
	}
}

type DatabaseRouting string

const (
	DatabaseRoutingMulti  DatabaseRouting = "Multi"
	DatabaseRoutingSingle DatabaseRouting = "Single"
)

func PossibleValuesForDatabaseRouting() []string {
	return []string{
		string(DatabaseRoutingMulti),
		string(DatabaseRoutingSingle),
	}
}

type EventGridDataFormat string

const (
	EventGridDataFormatAPACHEAVRO     EventGridDataFormat = "APACHEAVRO"
	EventGridDataFormatAVRO           EventGridDataFormat = "AVRO"
	EventGridDataFormatCSV            EventGridDataFormat = "CSV"
	EventGridDataFormatJSON           EventGridDataFormat = "JSON"
	EventGridDataFormatMULTIJSON      EventGridDataFormat = "MULTIJSON"
	EventGridDataFormatORC            EventGridDataFormat = "ORC"
	EventGridDataFormatPARQUET        EventGridDataFormat = "PARQUET"
	EventGridDataFormatPSV            EventGridDataFormat = "PSV"
	EventGridDataFormatRAW            EventGridDataFormat = "RAW"
	EventGridDataFormatSCSV           EventGridDataFormat = "SCSV"
	EventGridDataFormatSINGLEJSON     EventGridDataFormat = "SINGLEJSON"
	EventGridDataFormatSOHSV          EventGridDataFormat = "SOHSV"
	EventGridDataFormatTSV            EventGridDataFormat = "TSV"
	EventGridDataFormatTSVE           EventGridDataFormat = "TSVE"
	EventGridDataFormatTXT            EventGridDataFormat = "TXT"
	EventGridDataFormatWThreeCLOGFILE EventGridDataFormat = "W3CLOGFILE"
)

func PossibleValuesForEventGridDataFormat() []string {
	return []string{
		string(EventGridDataFormatAPACHEAVRO),
		string(EventGridDataFormatAVRO),
		string(EventGridDataFormatCSV),
		string(EventGridDataFormatJSON),
		string(EventGridDataFormatMULTIJSON),
		string(EventGridDataFormatORC),
		string(EventGridDataFormatPARQUET),
		string(EventGridDataFormatPSV),
		string(EventGridDataFormatRAW),
		string(EventGridDataFormatSCSV),
		string(EventGridDataFormatSINGLEJSON),
		string(EventGridDataFormatSOHSV),
		string(EventGridDataFormatTSV),
		string(EventGridDataFormatTSVE),
		string(EventGridDataFormatTXT),
		string(EventGridDataFormatWThreeCLOGFILE),
	}
}

type EventHubDataFormat string

const (
	EventHubDataFormatAPACHEAVRO     EventHubDataFormat = "APACHEAVRO"
	EventHubDataFormatAVRO           EventHubDataFormat = "AVRO"
	EventHubDataFormatCSV            EventHubDataFormat = "CSV"
	EventHubDataFormatJSON           EventHubDataFormat = "JSON"
	EventHubDataFormatMULTIJSON      EventHubDataFormat = "MULTIJSON"
	EventHubDataFormatORC            EventHubDataFormat = "ORC"
	EventHubDataFormatPARQUET        EventHubDataFormat = "PARQUET"
	EventHubDataFormatPSV            EventHubDataFormat = "PSV"
	EventHubDataFormatRAW            EventHubDataFormat = "RAW"
	EventHubDataFormatSCSV           EventHubDataFormat = "SCSV"
	EventHubDataFormatSINGLEJSON     EventHubDataFormat = "SINGLEJSON"
	EventHubDataFormatSOHSV          EventHubDataFormat = "SOHSV"
	EventHubDataFormatTSV            EventHubDataFormat = "TSV"
	EventHubDataFormatTSVE           EventHubDataFormat = "TSVE"
	EventHubDataFormatTXT            EventHubDataFormat = "TXT"
	EventHubDataFormatWThreeCLOGFILE EventHubDataFormat = "W3CLOGFILE"
)

func PossibleValuesForEventHubDataFormat() []string {
	return []string{
		string(EventHubDataFormatAPACHEAVRO),
		string(EventHubDataFormatAVRO),
		string(EventHubDataFormatCSV),
		string(EventHubDataFormatJSON),
		string(EventHubDataFormatMULTIJSON),
		string(EventHubDataFormatORC),
		string(EventHubDataFormatPARQUET),
		string(EventHubDataFormatPSV),
		string(EventHubDataFormatRAW),
		string(EventHubDataFormatSCSV),
		string(EventHubDataFormatSINGLEJSON),
		string(EventHubDataFormatSOHSV),
		string(EventHubDataFormatTSV),
		string(EventHubDataFormatTSVE),
		string(EventHubDataFormatTXT),
		string(EventHubDataFormatWThreeCLOGFILE),
	}
}

type IotHubDataFormat string

const (
	IotHubDataFormatAPACHEAVRO     IotHubDataFormat = "APACHEAVRO"
	IotHubDataFormatAVRO           IotHubDataFormat = "AVRO"
	IotHubDataFormatCSV            IotHubDataFormat = "CSV"
	IotHubDataFormatJSON           IotHubDataFormat = "JSON"
	IotHubDataFormatMULTIJSON      IotHubDataFormat = "MULTIJSON"
	IotHubDataFormatORC            IotHubDataFormat = "ORC"
	IotHubDataFormatPARQUET        IotHubDataFormat = "PARQUET"
	IotHubDataFormatPSV            IotHubDataFormat = "PSV"
	IotHubDataFormatRAW            IotHubDataFormat = "RAW"
	IotHubDataFormatSCSV           IotHubDataFormat = "SCSV"
	IotHubDataFormatSINGLEJSON     IotHubDataFormat = "SINGLEJSON"
	IotHubDataFormatSOHSV          IotHubDataFormat = "SOHSV"
	IotHubDataFormatTSV            IotHubDataFormat = "TSV"
	IotHubDataFormatTSVE           IotHubDataFormat = "TSVE"
	IotHubDataFormatTXT            IotHubDataFormat = "TXT"
	IotHubDataFormatWThreeCLOGFILE IotHubDataFormat = "W3CLOGFILE"
)

func PossibleValuesForIotHubDataFormat() []string {
	return []string{
		string(IotHubDataFormatAPACHEAVRO),
		string(IotHubDataFormatAVRO),
		string(IotHubDataFormatCSV),
		string(IotHubDataFormatJSON),
		string(IotHubDataFormatMULTIJSON),
		string(IotHubDataFormatORC),
		string(IotHubDataFormatPARQUET),
		string(IotHubDataFormatPSV),
		string(IotHubDataFormatRAW),
		string(IotHubDataFormatSCSV),
		string(IotHubDataFormatSINGLEJSON),
		string(IotHubDataFormatSOHSV),
		string(IotHubDataFormatTSV),
		string(IotHubDataFormatTSVE),
		string(IotHubDataFormatTXT),
		string(IotHubDataFormatWThreeCLOGFILE),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
	}
}

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
	ReasonInvalid       Reason = "Invalid"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
		string(ReasonInvalid),
	}
}
