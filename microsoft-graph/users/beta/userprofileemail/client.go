package userprofileemail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfileEmailClient struct {
	Client *msgraph.Client
}

func NewUserProfileEmailClientWithBaseURI(api sdkEnv.Api) (*UserProfileEmailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofileemail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfileEmailClient: %+v", err)
	}

	return &UserProfileEmailClient{
		Client: client,
	}, nil
}
