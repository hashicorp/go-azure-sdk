package useroutlooktaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskFolderClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskFolderClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskFolderClient: %+v", err)
	}

	return &UserOutlookTaskFolderClient{
		Client: client,
	}, nil
}
