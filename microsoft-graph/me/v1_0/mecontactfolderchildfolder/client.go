package mecontactfolderchildfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeContactFolderChildFolderClient struct {
	Client *msgraph.Client
}

func NewMeContactFolderChildFolderClientWithBaseURI(api sdkEnv.Api) (*MeContactFolderChildFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecontactfolderchildfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeContactFolderChildFolderClient: %+v", err)
	}

	return &MeContactFolderChildFolderClient{
		Client: client,
	}, nil
}
