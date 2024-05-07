package sims

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimClone struct {
	Sims             *[]string           `json:"sims,omitempty"`
	TargetSimGroupId *SimGroupResourceId `json:"targetSimGroupId,omitempty"`
}
