package usercontactfolderchildfoldercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactFolderChildFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewUserContactFolderChildFolderContactExtensionClientWithBaseURI(api sdkEnv.Api) (*UserContactFolderChildFolderContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactfolderchildfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactFolderChildFolderContactExtensionClient: %+v", err)
	}

	return &UserContactFolderChildFolderContactExtensionClient{
		Client: client,
	}, nil
}
