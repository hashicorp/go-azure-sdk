package sitecontenttype

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopySiteContentTypeToDefaultContentLocationRequest struct {
	DestinationFileName nullable.Type[string] `json:"destinationFileName,omitempty"`
	SourceFile          *beta.ItemReference   `json:"sourceFile,omitempty"`
}
