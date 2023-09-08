package usersecurity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityClient struct {
	Client *msgraph.Client
}

func NewUserSecurityClientWithBaseURI(api sdkEnv.Api) (*UserSecurityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersecurity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSecurityClient: %+v", err)
	}

	return &UserSecurityClient{
		Client: client,
	}, nil
}
