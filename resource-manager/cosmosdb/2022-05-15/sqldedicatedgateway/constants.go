package sqldedicatedgateway

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceSize string

const (
	ServiceSizeCosmosPointDEights  ServiceSize = "Cosmos.D8s"
	ServiceSizeCosmosPointDFours   ServiceSize = "Cosmos.D4s"
	ServiceSizeCosmosPointDOneSixs ServiceSize = "Cosmos.D16s"
)

func PossibleValuesForServiceSize() []string {
	return []string{
		string(ServiceSizeCosmosPointDEights),
		string(ServiceSizeCosmosPointDFours),
		string(ServiceSizeCosmosPointDOneSixs),
	}
}

type ServiceStatus string

const (
	ServiceStatusCreating ServiceStatus = "Creating"
	ServiceStatusDeleting ServiceStatus = "Deleting"
	ServiceStatusError    ServiceStatus = "Error"
	ServiceStatusRunning  ServiceStatus = "Running"
	ServiceStatusStopped  ServiceStatus = "Stopped"
	ServiceStatusUpdating ServiceStatus = "Updating"
)

func PossibleValuesForServiceStatus() []string {
	return []string{
		string(ServiceStatusCreating),
		string(ServiceStatusDeleting),
		string(ServiceStatusError),
		string(ServiceStatusRunning),
		string(ServiceStatusStopped),
		string(ServiceStatusUpdating),
	}
}

type ServiceType string

const (
	ServiceTypeDataTransfer             ServiceType = "DataTransfer"
	ServiceTypeGraphAPICompute          ServiceType = "GraphAPICompute"
	ServiceTypeMaterializedViewsBuilder ServiceType = "MaterializedViewsBuilder"
	ServiceTypeSqlDedicatedGateway      ServiceType = "SqlDedicatedGateway"
)

func PossibleValuesForServiceType() []string {
	return []string{
		string(ServiceTypeDataTransfer),
		string(ServiceTypeGraphAPICompute),
		string(ServiceTypeMaterializedViewsBuilder),
		string(ServiceTypeSqlDedicatedGateway),
	}
}
