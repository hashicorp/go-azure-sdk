package outlooktaskgrouptaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskGroupTaskFolderClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskGroupTaskFolderClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskGroupTaskFolderClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskgrouptaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskGroupTaskFolderClient: %+v", err)
	}

	return &OutlookTaskGroupTaskFolderClient{
		Client: client,
	}, nil
}
