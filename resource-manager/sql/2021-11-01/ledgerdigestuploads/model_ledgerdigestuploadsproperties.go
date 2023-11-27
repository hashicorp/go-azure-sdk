package ledgerdigestuploads

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LedgerDigestUploadsProperties struct {
	DigestStorageEndpoint *string                   `json:"digestStorageEndpoint,omitempty"`
	State                 *LedgerDigestUploadsState `json:"state,omitempty"`
}
