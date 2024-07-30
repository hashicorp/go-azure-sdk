package directoryroleassignmentschedulerequestprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestPrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequestprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestPrincipalClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestPrincipalClient{
		Client: client,
	}, nil
}
