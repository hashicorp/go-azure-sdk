package groupsiteanalyticitemactivitystatactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteAnalyticItemActivityStatActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteAnalyticItemActivityStatActivityDriveItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteAnalyticItemActivityStatActivityDriveItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteanalyticitemactivitystatactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteAnalyticItemActivityStatActivityDriveItemClient: %+v", err)
	}

	return &GroupSiteAnalyticItemActivityStatActivityDriveItemClient{
		Client: client,
	}, nil
}
