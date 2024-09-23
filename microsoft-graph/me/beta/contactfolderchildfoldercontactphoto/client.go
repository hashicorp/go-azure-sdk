package contactfolderchildfoldercontactphoto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContactFolderChildFolderContactPhotoClient struct {
	Client *msgraph.Client
}

func NewContactFolderChildFolderContactPhotoClientWithBaseURI(sdkApi sdkEnv.Api) (*ContactFolderChildFolderContactPhotoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "contactfolderchildfoldercontactphoto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContactFolderChildFolderContactPhotoClient: %+v", err)
	}

	return &ContactFolderChildFolderContactPhotoClient{
		Client: client,
	}, nil
}
