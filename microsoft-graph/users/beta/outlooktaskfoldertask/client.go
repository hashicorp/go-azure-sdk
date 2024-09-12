package outlooktaskfoldertask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskFolderTaskClient struct {
	Client *msgraph.Client
}

func NewOutlookTaskFolderTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*OutlookTaskFolderTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outlooktaskfoldertask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutlookTaskFolderTaskClient: %+v", err)
	}

	return &OutlookTaskFolderTaskClient{
		Client: client,
	}, nil
}
