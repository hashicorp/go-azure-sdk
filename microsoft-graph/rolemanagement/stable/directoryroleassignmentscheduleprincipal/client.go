package directoryroleassignmentscheduleprincipal

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryRoleAssignmentSchedulePrincipalClient struct {
	Client *msgraph.Client
}

func NewDirectoryRoleAssignmentSchedulePrincipalClientWithBaseURI(sdkApi sdkEnv.Api) (*DirectoryRoleAssignmentSchedulePrincipalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "directoryroleassignmentscheduleprincipal", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryRoleAssignmentSchedulePrincipalClient: %+v", err)
	}

	return &DirectoryRoleAssignmentSchedulePrincipalClient{
		Client: client,
	}, nil
}
