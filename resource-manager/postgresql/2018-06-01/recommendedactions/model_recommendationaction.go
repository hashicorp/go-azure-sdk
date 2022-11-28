package recommendedactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationAction struct {
	Id         *string                         `json:"id,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties *RecommendationActionProperties `json:"properties"`
	Type       *string                         `json:"type,omitempty"`
}
