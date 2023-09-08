package memailfolderchildfoldermessageextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderChildFolderMessageExtensionClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderChildFolderMessageExtensionClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderChildFolderMessageExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfolderchildfoldermessageextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderChildFolderMessageExtensionClient: %+v", err)
	}

	return &MeMailFolderChildFolderMessageExtensionClient{
		Client: client,
	}, nil
}
