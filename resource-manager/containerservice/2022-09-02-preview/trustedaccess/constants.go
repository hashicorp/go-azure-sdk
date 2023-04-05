package trustedaccess

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustedAccessRoleBindingProvisioningState string

const (
	TrustedAccessRoleBindingProvisioningStateDeleting  TrustedAccessRoleBindingProvisioningState = "Deleting"
	TrustedAccessRoleBindingProvisioningStateFailed    TrustedAccessRoleBindingProvisioningState = "Failed"
	TrustedAccessRoleBindingProvisioningStateSucceeded TrustedAccessRoleBindingProvisioningState = "Succeeded"
	TrustedAccessRoleBindingProvisioningStateUpdating  TrustedAccessRoleBindingProvisioningState = "Updating"
)

func PossibleValuesForTrustedAccessRoleBindingProvisioningState() []string {
	return []string{
		string(TrustedAccessRoleBindingProvisioningStateDeleting),
		string(TrustedAccessRoleBindingProvisioningStateFailed),
		string(TrustedAccessRoleBindingProvisioningStateSucceeded),
		string(TrustedAccessRoleBindingProvisioningStateUpdating),
	}
}
