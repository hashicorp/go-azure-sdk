package outlooktaskgrouptaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskGroupTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskGroupTaskFolderTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskGroupTaskFolderTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskgrouptaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskGroupTaskFolderTaskClient: %+v", err)
	}

	return &OutlookTaskGroupTaskFolderTaskClient{
		Client: client,
	}, nil
}
