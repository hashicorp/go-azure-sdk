package cognitiveservicesaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInjections struct {
	Scenario                   *ScenarioType `json:"scenario,omitempty"`
	SubnetArmId                *string       `json:"subnetArmId,omitempty"`
	UseMicrosoftManagedNetwork *bool         `json:"useMicrosoftManagedNetwork,omitempty"`
}
