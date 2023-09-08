package useroutlooktaskgrouptaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskGroupTaskFolderClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskGroupTaskFolderClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskGroupTaskFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskgrouptaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskGroupTaskFolderClient: %+v", err)
	}

	return &UserOutlookTaskGroupTaskFolderClient{
		Client: client,
	}, nil
}
