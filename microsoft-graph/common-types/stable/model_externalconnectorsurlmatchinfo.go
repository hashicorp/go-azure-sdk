package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsUrlMatchInfo struct {
	BaseUrls   *[]string `json:"baseUrls,omitempty"`
	ODataType  *string   `json:"@odata.type,omitempty"`
	UrlPattern *string   `json:"urlPattern,omitempty"`
}
