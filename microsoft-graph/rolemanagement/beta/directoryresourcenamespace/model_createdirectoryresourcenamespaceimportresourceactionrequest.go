package directoryresourcenamespace

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDirectoryResourceNamespaceImportResourceActionRequest struct {
	Format                     *string `json:"format,omitempty"`
	OverwriteResourceNamespace *bool   `json:"overwriteResourceNamespace,omitempty"`
	Value                      *string `json:"value,omitempty"`
}
