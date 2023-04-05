package datacollectionruleassociations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KnownDataCollectionRuleAssociationProvisioningState string

const (
	KnownDataCollectionRuleAssociationProvisioningStateCanceled  KnownDataCollectionRuleAssociationProvisioningState = "Canceled"
	KnownDataCollectionRuleAssociationProvisioningStateCreating  KnownDataCollectionRuleAssociationProvisioningState = "Creating"
	KnownDataCollectionRuleAssociationProvisioningStateDeleting  KnownDataCollectionRuleAssociationProvisioningState = "Deleting"
	KnownDataCollectionRuleAssociationProvisioningStateFailed    KnownDataCollectionRuleAssociationProvisioningState = "Failed"
	KnownDataCollectionRuleAssociationProvisioningStateSucceeded KnownDataCollectionRuleAssociationProvisioningState = "Succeeded"
	KnownDataCollectionRuleAssociationProvisioningStateUpdating  KnownDataCollectionRuleAssociationProvisioningState = "Updating"
)

func PossibleValuesForKnownDataCollectionRuleAssociationProvisioningState() []string {
	return []string{
		string(KnownDataCollectionRuleAssociationProvisioningStateCanceled),
		string(KnownDataCollectionRuleAssociationProvisioningStateCreating),
		string(KnownDataCollectionRuleAssociationProvisioningStateDeleting),
		string(KnownDataCollectionRuleAssociationProvisioningStateFailed),
		string(KnownDataCollectionRuleAssociationProvisioningStateSucceeded),
		string(KnownDataCollectionRuleAssociationProvisioningStateUpdating),
	}
}
