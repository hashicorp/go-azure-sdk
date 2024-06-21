package streamingpoliciesandstreaminglocators

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamingPolicyProperties struct {
	CommonEncryptionCbcs        *CommonEncryptionCbcs `json:"commonEncryptionCbcs,omitempty"`
	CommonEncryptionCenc        *CommonEncryptionCenc `json:"commonEncryptionCenc,omitempty"`
	Created                     *string               `json:"created,omitempty"`
	DefaultContentKeyPolicyName *string               `json:"defaultContentKeyPolicyName,omitempty"`
	EnvelopeEncryption          *EnvelopeEncryption   `json:"envelopeEncryption,omitempty"`
	NoEncryption                *NoEncryption         `json:"noEncryption,omitempty"`
}
