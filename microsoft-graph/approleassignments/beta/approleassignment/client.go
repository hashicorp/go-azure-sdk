package approleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewAppRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*AppRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "approleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppRoleAssignmentClient: %+v", err)
	}

	return &AppRoleAssignmentClient{
		Client: client,
	}, nil
}
