package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDataLakeAnalyticsAccountProperties struct {
	ComputePolicies              *[]CreateComputePolicyWithAccountParameters `json:"computePolicies,omitempty"`
	DataLakeStoreAccounts        []AddDataLakeStoreWithAccountParameters     `json:"dataLakeStoreAccounts"`
	DefaultDataLakeStoreAccount  string                                      `json:"defaultDataLakeStoreAccount"`
	FirewallAllowAzureIPs        *FirewallAllowAzureIPsState                 `json:"firewallAllowAzureIps,omitempty"`
	FirewallRules                *[]CreateFirewallRuleWithAccountParameters  `json:"firewallRules,omitempty"`
	FirewallState                *FirewallState                              `json:"firewallState,omitempty"`
	MaxDegreeOfParallelism       *int64                                      `json:"maxDegreeOfParallelism,omitempty"`
	MaxDegreeOfParallelismPerJob *int64                                      `json:"maxDegreeOfParallelismPerJob,omitempty"`
	MaxJobCount                  *int64                                      `json:"maxJobCount,omitempty"`
	MinPriorityPerJob            *int64                                      `json:"minPriorityPerJob,omitempty"`
	NewTier                      *TierType                                   `json:"newTier,omitempty"`
	QueryStoreRetention          *int64                                      `json:"queryStoreRetention,omitempty"`
	StorageAccounts              *[]AddStorageAccountWithAccountParameters   `json:"storageAccounts,omitempty"`
}
