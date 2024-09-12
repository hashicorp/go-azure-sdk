package sitecontenttype

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopySiteContentTypeToDefaultContentLocationRequest struct {
	DestinationFileName nullable.Type[string] `json:"destinationFileName,omitempty"`
	SourceFile          *stable.ItemReference `json:"sourceFile,omitempty"`
}
