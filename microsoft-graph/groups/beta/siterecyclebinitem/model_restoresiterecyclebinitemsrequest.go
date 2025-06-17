package siterecyclebinitem

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestoreSiteRecycleBinItemsRequest struct {
	Ids *[]string `json:"ids,omitempty"`
}
