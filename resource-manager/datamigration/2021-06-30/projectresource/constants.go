package projectresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectProvisioningState string

const (
	ProjectProvisioningStateDeleting  ProjectProvisioningState = "Deleting"
	ProjectProvisioningStateSucceeded ProjectProvisioningState = "Succeeded"
)

func PossibleValuesForProjectProvisioningState() []string {
	return []string{
		string(ProjectProvisioningStateDeleting),
		string(ProjectProvisioningStateSucceeded),
	}
}

type ProjectSourcePlatform string

const (
	ProjectSourcePlatformMongoDb    ProjectSourcePlatform = "MongoDb"
	ProjectSourcePlatformMySQL      ProjectSourcePlatform = "MySQL"
	ProjectSourcePlatformPostgreSql ProjectSourcePlatform = "PostgreSql"
	ProjectSourcePlatformSQL        ProjectSourcePlatform = "SQL"
	ProjectSourcePlatformUnknown    ProjectSourcePlatform = "Unknown"
)

func PossibleValuesForProjectSourcePlatform() []string {
	return []string{
		string(ProjectSourcePlatformMongoDb),
		string(ProjectSourcePlatformMySQL),
		string(ProjectSourcePlatformPostgreSql),
		string(ProjectSourcePlatformSQL),
		string(ProjectSourcePlatformUnknown),
	}
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformAzureDbForMySql      ProjectTargetPlatform = "AzureDbForMySql"
	ProjectTargetPlatformAzureDbForPostgreSql ProjectTargetPlatform = "AzureDbForPostgreSql"
	ProjectTargetPlatformMongoDb              ProjectTargetPlatform = "MongoDb"
	ProjectTargetPlatformSQLDB                ProjectTargetPlatform = "SQLDB"
	ProjectTargetPlatformSQLMI                ProjectTargetPlatform = "SQLMI"
	ProjectTargetPlatformUnknown              ProjectTargetPlatform = "Unknown"
)

func PossibleValuesForProjectTargetPlatform() []string {
	return []string{
		string(ProjectTargetPlatformAzureDbForMySql),
		string(ProjectTargetPlatformAzureDbForPostgreSql),
		string(ProjectTargetPlatformMongoDb),
		string(ProjectTargetPlatformSQLDB),
		string(ProjectTargetPlatformSQLMI),
		string(ProjectTargetPlatformUnknown),
	}
}
