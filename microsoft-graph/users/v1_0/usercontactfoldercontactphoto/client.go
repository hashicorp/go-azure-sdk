package usercontactfoldercontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactFolderContactPhotoClient struct {
	Client *msgraph.Client
}

func NewUserContactFolderContactPhotoClientWithBaseURI(api sdkEnv.Api) (*UserContactFolderContactPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactfoldercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactFolderContactPhotoClient: %+v", err)
	}

	return &UserContactFolderContactPhotoClient{
		Client: client,
	}, nil
}
