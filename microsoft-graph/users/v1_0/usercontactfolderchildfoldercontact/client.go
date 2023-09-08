package usercontactfolderchildfoldercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContactFolderChildFolderContactClient struct {
	Client *msgraph.Client
}

func NewUserContactFolderChildFolderContactClientWithBaseURI(api sdkEnv.Api) (*UserContactFolderChildFolderContactClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercontactfolderchildfoldercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserContactFolderChildFolderContactClient: %+v", err)
	}

	return &UserContactFolderChildFolderContactClient{
		Client: client,
	}, nil
}
