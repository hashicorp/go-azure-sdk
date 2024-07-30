package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerCollectionAssignmentTarget struct {
	CollectionId *string `json:"collectionId,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
