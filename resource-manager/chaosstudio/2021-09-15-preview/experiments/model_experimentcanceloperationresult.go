package experiments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentCancelOperationResult struct {
	Name      *string `json:"name,omitempty"`
	StatusUrl *string `json:"statusUrl,omitempty"`
}
