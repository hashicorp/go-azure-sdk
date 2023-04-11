package securescore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreItemProperties struct {
	DisplayName *string       `json:"displayName,omitempty"`
	Score       *ScoreDetails `json:"score,omitempty"`
	Weight      *int64        `json:"weight,omitempty"`
}
