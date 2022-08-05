package querytexts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryTextProperties struct {
	QueryId   *string `json:"queryId,omitempty"`
	QueryText *string `json:"queryText,omitempty"`
}
