package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureActionPermissionsDefinitionAction struct {
	Actions   *[]string `json:"actions,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
}
