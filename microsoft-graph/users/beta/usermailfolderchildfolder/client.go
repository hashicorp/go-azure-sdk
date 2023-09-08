package usermailfolderchildfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserMailFolderChildFolderClient struct {
	Client *msgraph.Client
}

func NewUserMailFolderChildFolderClientWithBaseURI(api sdkEnv.Api) (*UserMailFolderChildFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usermailfolderchildfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserMailFolderChildFolderClient: %+v", err)
	}

	return &UserMailFolderChildFolderClient{
		Client: client,
	}, nil
}
