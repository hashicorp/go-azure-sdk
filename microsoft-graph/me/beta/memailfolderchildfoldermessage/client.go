package memailfolderchildfoldermessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeMailFolderChildFolderMessageClient struct {
	Client *msgraph.Client
}

func NewMeMailFolderChildFolderMessageClientWithBaseURI(api sdkEnv.Api) (*MeMailFolderChildFolderMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "memailfolderchildfoldermessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeMailFolderChildFolderMessageClient: %+v", err)
	}

	return &MeMailFolderChildFolderMessageClient{
		Client: client,
	}, nil
}
