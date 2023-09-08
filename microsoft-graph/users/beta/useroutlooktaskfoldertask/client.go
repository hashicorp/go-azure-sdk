package useroutlooktaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskFolderTaskClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskFolderTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskFolderTaskClient: %+v", err)
	}

	return &UserOutlookTaskFolderTaskClient{
		Client: client,
	}, nil
}
