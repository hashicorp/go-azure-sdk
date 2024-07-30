package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppMetadata struct {
	Data      *[]AppMetadataEntry `json:"data,omitempty"`
	ODataType *string             `json:"@odata.type,omitempty"`
	Version   *int64              `json:"version,omitempty"`
}
