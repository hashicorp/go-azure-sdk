package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorInfo struct {
	AnalysisType     *[]string       `json:"analysisType,omitempty"`
	Author           *string         `json:"author,omitempty"`
	Category         *string         `json:"category,omitempty"`
	Description      *string         `json:"description,omitempty"`
	Id               *string         `json:"id,omitempty"`
	Name             *string         `json:"name,omitempty"`
	Score            *float64        `json:"score,omitempty"`
	SupportTopicList *[]SupportTopic `json:"supportTopicList,omitempty"`
	Type             *DetectorType   `json:"type,omitempty"`
}
