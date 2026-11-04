package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NSPConfigAccessRule struct {
	Name       *string                        `json:"name,omitempty"`
	Properties *NSPConfigAccessRuleProperties `json:"properties,omitempty"`
}
