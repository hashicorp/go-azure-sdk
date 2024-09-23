package followedsite

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FollowedSiteClient struct {
	Client *msgraph.Client
}

func NewFollowedSiteClientWithBaseURI(sdkApi sdkEnv.Api) (*FollowedSiteClient, error) {
	client, err := msgraph.NewClient(sdkApi, "followedsite", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FollowedSiteClient: %+v", err)
	}

	return &FollowedSiteClient{
		Client: client,
	}, nil
}
