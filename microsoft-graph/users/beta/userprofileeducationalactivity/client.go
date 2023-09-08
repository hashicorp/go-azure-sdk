package userprofileeducationalactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileEducationalActivityClient struct {
	Client *msgraph.Client
}

func NewUserProfileEducationalActivityClientWithBaseURI(api sdkEnv.Api) (*UserProfileEducationalActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileeducationalactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileEducationalActivityClient: %+v", err)
	}

	return &UserProfileEducationalActivityClient{
		Client: client,
	}, nil
}
