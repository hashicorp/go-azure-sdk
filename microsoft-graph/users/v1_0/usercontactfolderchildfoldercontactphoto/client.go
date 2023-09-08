package usercontactfolderchildfoldercontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactFolderChildFolderContactPhotoClient struct {
	Client *msgraph.Client
}

func NewUserContactFolderChildFolderContactPhotoClientWithBaseURI(api sdkEnv.Api) (*UserContactFolderChildFolderContactPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactfolderchildfoldercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactFolderChildFolderContactPhotoClient: %+v", err)
	}

	return &UserContactFolderChildFolderContactPhotoClient{
		Client: client,
	}, nil
}
