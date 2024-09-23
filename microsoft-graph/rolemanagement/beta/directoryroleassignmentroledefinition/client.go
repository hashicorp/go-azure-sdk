package directoryroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &DirectoryRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
