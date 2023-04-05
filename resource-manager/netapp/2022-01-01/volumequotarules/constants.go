package volumequotarules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStatePatching  ProvisioningState = "Patching"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStatePatching),
		string(ProvisioningStateSucceeded),
	}
}

type Type string

const (
	TypeDefaultGroupQuota    Type = "DefaultGroupQuota"
	TypeDefaultUserQuota     Type = "DefaultUserQuota"
	TypeIndividualGroupQuota Type = "IndividualGroupQuota"
	TypeIndividualUserQuota  Type = "IndividualUserQuota"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeDefaultGroupQuota),
		string(TypeDefaultUserQuota),
		string(TypeIndividualGroupQuota),
		string(TypeIndividualUserQuota),
	}
}
