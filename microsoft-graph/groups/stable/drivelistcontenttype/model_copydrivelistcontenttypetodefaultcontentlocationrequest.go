package drivelistcontenttype

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CopyDriveListContentTypeToDefaultContentLocationRequest struct {
	DestinationFileName nullable.Type[string] `json:"destinationFileName,omitempty"`
	SourceFile          *stable.ItemReference `json:"sourceFile,omitempty"`
}
