package meonenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeOnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewMeOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(api sdkEnv.Api) (*MeOnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meonenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeOnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &MeOnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
