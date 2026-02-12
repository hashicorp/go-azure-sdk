package attestation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TpmAttestationResponse struct {
	Data *string `json:"data,omitempty"`
}
