package usercontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactPhotoClient struct {
	Client *msgraph.Client
}

func NewUserContactPhotoClientWithBaseURI(api sdkEnv.Api) (*UserContactPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactPhotoClient: %+v", err)
	}

	return &UserContactPhotoClient{
		Client: client,
	}, nil
}
