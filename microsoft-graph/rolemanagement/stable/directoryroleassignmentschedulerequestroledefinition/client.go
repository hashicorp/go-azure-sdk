package directoryroleassignmentschedulerequestroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentScheduleRequestRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentschedulerequestroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleAssignmentScheduleRequestRoleDefinitionClient{
		Client: client,
	}, nil
}
