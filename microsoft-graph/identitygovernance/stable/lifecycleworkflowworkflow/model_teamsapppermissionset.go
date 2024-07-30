package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppPermissionSet struct {
	ODataType                   *string                               `json:"@odata.type,omitempty"`
	ResourceSpecificPermissions *[]TeamsAppResourceSpecificPermission `json:"resourceSpecificPermissions,omitempty"`
}
