package recommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Instructions struct {
	ActionsToBePerformed      string  `json:"actionsToBePerformed"`
	HowToPerformActionDetails *string `json:"howToPerformActionDetails,omitempty"`
	RecommendationImportance  string  `json:"recommendationImportance"`
}
