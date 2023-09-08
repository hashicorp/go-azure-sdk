package useroutlooktaskgrouptaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserOutlookTaskGroupTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewUserOutlookTaskGroupTaskFolderTaskClientWithBaseURI(api sdkEnv.Api) (*UserOutlookTaskGroupTaskFolderTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "useroutlooktaskgrouptaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserOutlookTaskGroupTaskFolderTaskClient: %+v", err)
	}

	return &UserOutlookTaskGroupTaskFolderTaskClient{
		Client: client,
	}, nil
}
