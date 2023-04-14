package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CachedImages struct {
	Image  string `json:"image"`
	OsType string `json:"osType"`
}
