package groupsitetermstoresetchildrenchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenChildrenRelationToTermClient{
		Client: client,
	}, nil
}
