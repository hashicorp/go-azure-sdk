package permissionsmanagementscheduledpermissionsapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsManagementScheduledPermissionsApprovalStepClient struct {
	Client *msgraph.Client
}

func NewPermissionsManagementScheduledPermissionsApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsManagementScheduledPermissionsApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsmanagementscheduledpermissionsapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsManagementScheduledPermissionsApprovalStepClient: %+v", err)
	}

	return &PermissionsManagementScheduledPermissionsApprovalStepClient{
		Client: client,
	}, nil
}
