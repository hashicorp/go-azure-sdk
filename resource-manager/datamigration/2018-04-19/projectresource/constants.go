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
	ProjectSourcePlatformSQL     ProjectSourcePlatform = "SQL"
	ProjectSourcePlatformUnknown ProjectSourcePlatform = "Unknown"
)

func PossibleValuesForProjectSourcePlatform() []string {
	return []string{
		string(ProjectSourcePlatformSQL),
		string(ProjectSourcePlatformUnknown),
	}
}

type ProjectTargetPlatform string

const (
	ProjectTargetPlatformSQLDB   ProjectTargetPlatform = "SQLDB"
	ProjectTargetPlatformUnknown ProjectTargetPlatform = "Unknown"
)

func PossibleValuesForProjectTargetPlatform() []string {
	return []string{
		string(ProjectTargetPlatformSQLDB),
		string(ProjectTargetPlatformUnknown),
	}
}
