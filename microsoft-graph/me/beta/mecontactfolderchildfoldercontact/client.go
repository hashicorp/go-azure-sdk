package mecontactfolderchildfoldercontact

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderChildFolderContactClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderChildFolderContactClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderChildFolderContactClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfolderchildfoldercontact", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderChildFolderContactClient: %+v", err)
	}

	return &MeContactFolderChildFolderContactClient{
		Client: client,
	}, nil
}
