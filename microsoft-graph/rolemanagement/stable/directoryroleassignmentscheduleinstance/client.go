package directoryroleassignmentscheduleinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstanceClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstanceClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstanceClient{
		Client: client,
	}, nil
}
