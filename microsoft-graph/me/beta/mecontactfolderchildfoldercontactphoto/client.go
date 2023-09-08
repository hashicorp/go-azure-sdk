package mecontactfolderchildfoldercontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderChildFolderContactPhotoClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderChildFolderContactPhotoClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderChildFolderContactPhotoClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfolderchildfoldercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderChildFolderContactPhotoClient: %+v", err)
	}

	return &MeContactFolderChildFolderContactPhotoClient{
		Client: client,
	}, nil
}
