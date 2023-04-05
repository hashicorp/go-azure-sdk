package hybridrunbookworker

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *WorkerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForWorkerType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = WorkerType(decoded)
	return nil
}
