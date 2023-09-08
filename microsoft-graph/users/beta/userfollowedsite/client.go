package userfollowedsite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFollowedSiteClient struct {
	Client *msgraph.Client
}

func NewUserFollowedSiteClientWithBaseURI(api sdkEnv.Api) (*UserFollowedSiteClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userfollowedsite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserFollowedSiteClient: %+v", err)
	}

	return &UserFollowedSiteClient{
		Client: client,
	}, nil
}
