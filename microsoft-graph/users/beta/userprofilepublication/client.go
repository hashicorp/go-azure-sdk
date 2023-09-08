package userprofilepublication

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProfilePublicationClient struct {
	Client *msgraph.Client
}

func NewUserProfilePublicationClientWithBaseURI(api sdkEnv.Api) (*UserProfilePublicationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userprofilepublication", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserProfilePublicationClient: %+v", err)
	}

	return &UserProfilePublicationClient{
		Client: client,
	}, nil
}
