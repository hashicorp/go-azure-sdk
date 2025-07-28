package loganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RankingsResponseTablesInlined struct {
	Data    *[]RankingsResponseTablesInlinedDataInlined `json:"data,omitempty"`
	Ranking *string                                     `json:"ranking,omitempty"`
}
