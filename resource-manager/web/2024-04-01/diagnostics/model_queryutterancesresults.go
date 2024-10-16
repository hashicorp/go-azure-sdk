package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryUtterancesResults struct {
	Query   *string                  `json:"query,omitempty"`
	Results *[]QueryUtterancesResult `json:"results,omitempty"`
}
