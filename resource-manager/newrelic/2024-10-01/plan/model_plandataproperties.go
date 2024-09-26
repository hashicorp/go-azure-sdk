package plan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlanDataProperties struct {
	AccountCreationSource *AccountCreationSource `json:"accountCreationSource,omitempty"`
	OrgCreationSource     *OrgCreationSource     `json:"orgCreationSource,omitempty"`
	PlanData              *PlanData              `json:"planData,omitempty"`
}
