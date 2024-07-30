package directoryroleassignmentscheduleinstanceprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstancePrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstancePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstancePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstanceprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstancePrincipalClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstancePrincipalClient{
		Client: client,
	}, nil
}
