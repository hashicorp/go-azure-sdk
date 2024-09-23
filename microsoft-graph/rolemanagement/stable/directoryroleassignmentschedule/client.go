package directoryroleassignmentschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleClient{
		Client: client,
	}, nil
}
