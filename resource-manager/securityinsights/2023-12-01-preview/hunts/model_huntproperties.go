package hunts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HuntProperties struct {
	AttackTactics    *[]AttackTactic   `json:"attackTactics,omitempty"`
	AttackTechniques *[]string         `json:"attackTechniques,omitempty"`
	Description      string            `json:"description"`
	DisplayName      string            `json:"displayName"`
	HypothesisStatus *HypothesisStatus `json:"hypothesisStatus,omitempty"`
	Labels           *[]string         `json:"labels,omitempty"`
	Owner            *HuntOwner        `json:"owner,omitempty"`
	Status           *Status           `json:"status,omitempty"`
}
