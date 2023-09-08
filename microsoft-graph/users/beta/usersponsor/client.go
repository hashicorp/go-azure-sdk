package usersponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSponsorClient struct {
	Client *msgraph.Client
}

func NewUserSponsorClientWithBaseURI(api sdkEnv.Api) (*UserSponsorClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usersponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSponsorClient: %+v", err)
	}

	return &UserSponsorClient{
		Client: client,
	}, nil
}
