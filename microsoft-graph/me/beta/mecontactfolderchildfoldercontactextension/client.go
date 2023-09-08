package mecontactfolderchildfoldercontactextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderChildFolderContactExtensionClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderChildFolderContactExtensionClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderChildFolderContactExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfolderchildfoldercontactextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderChildFolderContactExtensionClient: %+v", err)
	}

	return &MeContactFolderChildFolderContactExtensionClient{
		Client: client,
	}, nil
}
