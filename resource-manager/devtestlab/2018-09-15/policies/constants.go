package policies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyEvaluatorType string

const (
	PolicyEvaluatorTypeAllowedValuesPolicy PolicyEvaluatorType = "AllowedValuesPolicy"
	PolicyEvaluatorTypeMaxValuePolicy      PolicyEvaluatorType = "MaxValuePolicy"
)

func PossibleValuesForPolicyEvaluatorType() []string {
	return []string{
		string(PolicyEvaluatorTypeAllowedValuesPolicy),
		string(PolicyEvaluatorTypeMaxValuePolicy),
	}
}

type PolicyFactName string

const (
	PolicyFactNameEnvironmentTemplate         PolicyFactName = "EnvironmentTemplate"
	PolicyFactNameGalleryImage                PolicyFactName = "GalleryImage"
	PolicyFactNameLabPremiumVMCount           PolicyFactName = "LabPremiumVmCount"
	PolicyFactNameLabTargetCost               PolicyFactName = "LabTargetCost"
	PolicyFactNameLabVMCount                  PolicyFactName = "LabVmCount"
	PolicyFactNameLabVMSize                   PolicyFactName = "LabVmSize"
	PolicyFactNameScheduleEditPermission      PolicyFactName = "ScheduleEditPermission"
	PolicyFactNameUserOwnedLabPremiumVMCount  PolicyFactName = "UserOwnedLabPremiumVmCount"
	PolicyFactNameUserOwnedLabVMCount         PolicyFactName = "UserOwnedLabVmCount"
	PolicyFactNameUserOwnedLabVMCountInSubnet PolicyFactName = "UserOwnedLabVmCountInSubnet"
)

func PossibleValuesForPolicyFactName() []string {
	return []string{
		string(PolicyFactNameEnvironmentTemplate),
		string(PolicyFactNameGalleryImage),
		string(PolicyFactNameLabPremiumVMCount),
		string(PolicyFactNameLabTargetCost),
		string(PolicyFactNameLabVMCount),
		string(PolicyFactNameLabVMSize),
		string(PolicyFactNameScheduleEditPermission),
		string(PolicyFactNameUserOwnedLabPremiumVMCount),
		string(PolicyFactNameUserOwnedLabVMCount),
		string(PolicyFactNameUserOwnedLabVMCountInSubnet),
	}
}

type PolicyStatus string

const (
	PolicyStatusDisabled PolicyStatus = "Disabled"
	PolicyStatusEnabled  PolicyStatus = "Enabled"
)

func PossibleValuesForPolicyStatus() []string {
	return []string{
		string(PolicyStatusDisabled),
		string(PolicyStatusEnabled),
	}
}
