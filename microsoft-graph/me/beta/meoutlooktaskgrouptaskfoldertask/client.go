package meoutlooktaskgrouptaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskGroupTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskGroupTaskFolderTaskClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskGroupTaskFolderTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskgrouptaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskGroupTaskFolderTaskClient: %+v", err)
	}

	return &MeOutlookTaskGroupTaskFolderTaskClient{
		Client: client,
	}, nil
}
