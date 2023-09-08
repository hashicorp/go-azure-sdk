package meonenotesectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupParentNotebookClient: %+v", err)
	}

	return &MeOnenoteSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
