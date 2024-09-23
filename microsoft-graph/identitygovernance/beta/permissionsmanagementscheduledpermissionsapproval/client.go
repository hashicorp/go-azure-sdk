package permissionsmanagementscheduledpermissionsapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsManagementScheduledPermissionsApprovalClient struct {
	Client *msgraph.Client
}

func NewPermissionsManagementScheduledPermissionsApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsManagementScheduledPermissionsApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsmanagementscheduledpermissionsapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsManagementScheduledPermissionsApprovalClient: %+v", err)
	}

	return &PermissionsManagementScheduledPermissionsApprovalClient{
		Client: client,
	}, nil
}
