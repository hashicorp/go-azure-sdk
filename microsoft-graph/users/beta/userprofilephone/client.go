package userprofilephone

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfilePhoneClient struct {
	Client *msgraph.Client
}

func NewUserProfilePhoneClientWithBaseURI(api sdkEnv.Api) (*UserProfilePhoneClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilephone", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfilePhoneClient: %+v", err)
	}

	return &UserProfilePhoneClient{
		Client: client,
	}, nil
}
