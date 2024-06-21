package databaseadvisors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorProperties struct {
	AdvisorStatus                  *AdvisorStatus                  `json:"advisorStatus,omitempty"`
	AutoExecuteStatus              AutoExecuteStatus               `json:"autoExecuteStatus"`
	AutoExecuteStatusInheritedFrom *AutoExecuteStatusInheritedFrom `json:"autoExecuteStatusInheritedFrom,omitempty"`
	LastChecked                    *string                         `json:"lastChecked,omitempty"`
	RecommendationsStatus          *string                         `json:"recommendationsStatus,omitempty"`
	RecommendedActions             *[]RecommendedAction            `json:"recommendedActions,omitempty"`
}
