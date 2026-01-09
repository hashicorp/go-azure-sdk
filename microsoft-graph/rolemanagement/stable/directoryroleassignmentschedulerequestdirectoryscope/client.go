package directoryroleassignmentschedulerequestdirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequestdirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestDirectoryScopeClient{
		Client: client,
	}, nil
}
