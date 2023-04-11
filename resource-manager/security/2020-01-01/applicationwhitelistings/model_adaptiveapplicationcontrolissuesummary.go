package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlIssueSummary struct {
	Issue       *Issue   `json:"issue,omitempty"`
	NumberOfVMs *float64 `json:"numberOfVms,omitempty"`
}
