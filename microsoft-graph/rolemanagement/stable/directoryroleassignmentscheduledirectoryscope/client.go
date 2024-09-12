package directoryroleassignmentscheduledirectoryscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleDirectoryScopeClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleDirectoryScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleDirectoryScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduledirectoryscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleDirectoryScopeClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleDirectoryScopeClient{
		Client: client,
	}, nil
}
