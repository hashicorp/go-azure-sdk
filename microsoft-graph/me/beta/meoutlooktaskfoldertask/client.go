package meoutlooktaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOutlookTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewMeOutlookTaskFolderTaskClientWithBaseURI(api sdkEnv.Api) (*MeOutlookTaskFolderTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meoutlooktaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOutlookTaskFolderTaskClient: %+v", err)
	}

	return &MeOutlookTaskFolderTaskClient{
		Client: client,
	}, nil
}
