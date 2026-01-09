package directoryroleassignmentscheduleinstanceappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleInstanceAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleInstanceAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleInstanceAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleinstanceappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleInstanceAppScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleInstanceAppScopeClient{
		Client: client,
	}, nil
}
