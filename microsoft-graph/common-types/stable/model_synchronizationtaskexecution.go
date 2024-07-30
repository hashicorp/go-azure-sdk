package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationTaskExecution struct {
	ActivityIdentifier           *string                            `json:"activityIdentifier,omitempty"`
	CountEntitled                *int64                             `json:"countEntitled,omitempty"`
	CountEntitledForProvisioning *int64                             `json:"countEntitledForProvisioning,omitempty"`
	CountEscrowed                *int64                             `json:"countEscrowed,omitempty"`
	CountEscrowedRaw             *int64                             `json:"countEscrowedRaw,omitempty"`
	CountExported                *int64                             `json:"countExported,omitempty"`
	CountExports                 *int64                             `json:"countExports,omitempty"`
	CountImported                *int64                             `json:"countImported,omitempty"`
	CountImportedDeltas          *int64                             `json:"countImportedDeltas,omitempty"`
	CountImportedReferenceDeltas *int64                             `json:"countImportedReferenceDeltas,omitempty"`
	Error                        *SynchronizationError              `json:"error,omitempty"`
	ODataType                    *string                            `json:"@odata.type,omitempty"`
	State                        *SynchronizationTaskExecutionState `json:"state,omitempty"`
	TimeBegan                    *string                            `json:"timeBegan,omitempty"`
	TimeEnded                    *string                            `json:"timeEnded,omitempty"`
}
