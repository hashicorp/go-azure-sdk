package linkedstorageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSourceType string

const (
	DataSourceTypeAlerts      DataSourceType = "Alerts"
	DataSourceTypeAzureWatson DataSourceType = "AzureWatson"
	DataSourceTypeCustomLogs  DataSourceType = "CustomLogs"
	DataSourceTypeIngestion   DataSourceType = "Ingestion"
	DataSourceTypeQuery       DataSourceType = "Query"
)

func PossibleValuesForDataSourceType() []string {
	return []string{
		string(DataSourceTypeAlerts),
		string(DataSourceTypeAzureWatson),
		string(DataSourceTypeCustomLogs),
		string(DataSourceTypeIngestion),
		string(DataSourceTypeQuery),
	}
}
