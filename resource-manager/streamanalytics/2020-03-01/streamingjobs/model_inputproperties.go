package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InputProperties struct {
	Compression   *Compression   `json:"compression,omitempty"`
	Diagnostics   *Diagnostics   `json:"diagnostics,omitempty"`
	Etag          *string        `json:"etag,omitempty"`
	PartitionKey  *string        `json:"partitionKey,omitempty"`
	Serialization *Serialization `json:"serialization,omitempty"`
	Type          string         `json:"type"`
}
