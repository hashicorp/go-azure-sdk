package policydescription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyScopeContract string

const (
	PolicyScopeContractAll       PolicyScopeContract = "All"
	PolicyScopeContractApi       PolicyScopeContract = "Api"
	PolicyScopeContractOperation PolicyScopeContract = "Operation"
	PolicyScopeContractProduct   PolicyScopeContract = "Product"
	PolicyScopeContractTenant    PolicyScopeContract = "Tenant"
)

func PossibleValuesForPolicyScopeContract() []string {
	return []string{
		string(PolicyScopeContractAll),
		string(PolicyScopeContractApi),
		string(PolicyScopeContractOperation),
		string(PolicyScopeContractProduct),
		string(PolicyScopeContractTenant),
	}
}
