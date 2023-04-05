package hybridrunbookworker

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkerType string

const (
	WorkerTypeHybridVOne WorkerType = "HybridV1"
	WorkerTypeHybridVTwo WorkerType = "HybridV2"
)

func PossibleValuesForWorkerType() []string {
	return []string{
		string(WorkerTypeHybridVOne),
		string(WorkerTypeHybridVTwo),
	}
}
