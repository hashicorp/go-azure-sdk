package userphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPhotoClient struct {
	Client *msgraph.Client
}

func NewUserPhotoClientWithBaseURI(api sdkEnv.Api) (*UserPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPhotoClient: %+v", err)
	}

	return &UserPhotoClient{
		Client: client,
	}, nil
}
