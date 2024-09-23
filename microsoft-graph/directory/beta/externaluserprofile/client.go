package externaluserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalUserProfileClient struct {
	Client *msgraph.Client
}

func NewExternalUserProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*ExternalUserProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "externaluserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExternalUserProfileClient: %+v", err)
	}

	return &ExternalUserProfileClient{
		Client: client,
	}, nil
}
