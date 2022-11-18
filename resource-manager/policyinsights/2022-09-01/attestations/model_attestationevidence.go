package attestations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationEvidence struct {
	Description *string `json:"description,omitempty"`
	SourceUri   *string `json:"sourceUri,omitempty"`
}
