package directoryroleassignmentscheduleappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleAppScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleAppScopeClient{
		Client: client,
	}, nil
}
