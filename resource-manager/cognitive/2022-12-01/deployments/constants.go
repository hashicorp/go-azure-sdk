package deployments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentProvisioningState string

const (
	DeploymentProvisioningStateAccepted  DeploymentProvisioningState = "Accepted"
	DeploymentProvisioningStateCreating  DeploymentProvisioningState = "Creating"
	DeploymentProvisioningStateDeleting  DeploymentProvisioningState = "Deleting"
	DeploymentProvisioningStateFailed    DeploymentProvisioningState = "Failed"
	DeploymentProvisioningStateMoving    DeploymentProvisioningState = "Moving"
	DeploymentProvisioningStateSucceeded DeploymentProvisioningState = "Succeeded"
)

func PossibleValuesForDeploymentProvisioningState() []string {
	return []string{
		string(DeploymentProvisioningStateAccepted),
		string(DeploymentProvisioningStateCreating),
		string(DeploymentProvisioningStateDeleting),
		string(DeploymentProvisioningStateFailed),
		string(DeploymentProvisioningStateMoving),
		string(DeploymentProvisioningStateSucceeded),
	}
}

type DeploymentScaleType string

const (
	DeploymentScaleTypeManual   DeploymentScaleType = "Manual"
	DeploymentScaleTypeStandard DeploymentScaleType = "Standard"
)

func PossibleValuesForDeploymentScaleType() []string {
	return []string{
		string(DeploymentScaleTypeManual),
		string(DeploymentScaleTypeStandard),
	}
}
