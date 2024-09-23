package directoryroleassignmentprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentPrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentPrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentPrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentPrincipalClient: %+v", err)
	}

	return &DirectoryRoleAssignmentPrincipalClient{
		Client: client,
	}, nil
}
