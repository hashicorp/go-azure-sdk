package directoryroleassignmentschedulerequestappscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestAppScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestAppScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestAppScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequestappscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestAppScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestAppScopeClient{
		Client: client,
	}, nil
}
