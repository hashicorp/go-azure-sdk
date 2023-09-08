package meoutlooktaskgrouptaskfolder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskGroupTaskFolderClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskGroupTaskFolderClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskGroupTaskFolderClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskgrouptaskfolder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskGroupTaskFolderClient: %+v", err)
	}

	return &MeOutlookTaskGroupTaskFolderClient{
		Client: client,
	}, nil
}
