package cosmosdb

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreParameters struct {
	DatabasesToRestore        *[]DatabaseRestoreResource        `json:"databasesToRestore,omitempty"`
	GremlinDatabasesToRestore *[]GremlinDatabaseRestoreResource `json:"gremlinDatabasesToRestore,omitempty"`
	RestoreMode               *RestoreMode                      `json:"restoreMode,omitempty"`
	RestoreSource             *string                           `json:"restoreSource,omitempty"`
	RestoreTimestampInUtc     *string                           `json:"restoreTimestampInUtc,omitempty"`
	TablesToRestore           *[]string                         `json:"tablesToRestore,omitempty"`
}
