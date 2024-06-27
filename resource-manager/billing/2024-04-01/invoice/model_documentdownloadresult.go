package invoice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentDownloadResult struct {
	ExpiryTime *string `json:"expiryTime,omitempty"`
	Url        *string `json:"url,omitempty"`
}
