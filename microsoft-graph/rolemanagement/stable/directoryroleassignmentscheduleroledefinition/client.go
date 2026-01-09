package directoryroleassignmentscheduleroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRoleDefinitionClient{
		Client: client,
	}, nil
}
