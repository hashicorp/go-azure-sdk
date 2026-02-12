package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestSgxEnclaveRequest struct {
	DraftPolicyForAttestation *string       `json:"draftPolicyForAttestation,omitempty"`
	InitTimeData              *InitTimeData `json:"initTimeData,omitempty"`
	Nonce                     *string       `json:"nonce,omitempty"`
	Quote                     *string       `json:"quote,omitempty"`
	RuntimeData               *RuntimeData  `json:"runtimeData,omitempty"`
}
