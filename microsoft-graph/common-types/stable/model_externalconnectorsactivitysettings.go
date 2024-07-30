package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsActivitySettings struct {
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	UrlToItemResolvers *[]ExternalConnectorsUrlToItemResolverBase `json:"urlToItemResolvers,omitempty"`
}
