package groupsiteanalyticitemactivitystatactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteAnalyticItemActivityStatActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewGroupSiteAnalyticItemActivityStatActivityDriveItemContentClientWithBaseURI(api sdkEnv.Api) (*GroupSiteAnalyticItemActivityStatActivityDriveItemContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteanalyticitemactivitystatactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteAnalyticItemActivityStatActivityDriveItemContentClient: %+v", err)
	}

	return &GroupSiteAnalyticItemActivityStatActivityDriveItemContentClient{
		Client: client,
	}, nil
}
