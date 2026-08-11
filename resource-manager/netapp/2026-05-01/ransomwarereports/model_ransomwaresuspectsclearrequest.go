package ransomwarereports

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RansomwareSuspectsClearRequest struct {
	Extensions []string                    `json:"extensions"`
	Resolution RansomwareSuspectResolution `json:"resolution"`
}
