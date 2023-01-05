package commitmenttiers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentCost struct {
	CommitmentMeterId *string `json:"commitmentMeterId,omitempty"`
	OverageMeterId    *string `json:"overageMeterId,omitempty"`
}
