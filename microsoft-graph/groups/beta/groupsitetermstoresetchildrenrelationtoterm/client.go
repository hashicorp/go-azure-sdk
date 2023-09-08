package groupsitetermstoresetchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenRelationToTermClient{
		Client: client,
	}, nil
}
