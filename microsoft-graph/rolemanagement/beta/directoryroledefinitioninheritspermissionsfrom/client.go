package directoryroledefinitioninheritspermissionsfrom

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleDefinitionInheritsPermissionsFromClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleDefinitionInheritsPermissionsFromClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleDefinitionInheritsPermissionsFromClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroledefinitioninheritspermissionsfrom", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleDefinitionInheritsPermissionsFromClient: %+v", err)
	}

	return &DirectoryRoleDefinitionInheritsPermissionsFromClient{
		Client: client,
	}, nil
}
