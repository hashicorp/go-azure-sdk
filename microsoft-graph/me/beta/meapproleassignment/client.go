package meapproleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAppRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewMeAppRoleAssignmentClientWithBaseURI(api sdkEnv.Api) (*MeAppRoleAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meapproleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAppRoleAssignmentClient: %+v", err)
	}

	return &MeAppRoleAssignmentClient{
		Client: client,
	}, nil
}
