package medirectreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeDirectReportClient struct {
	Client *msgraph.Client
}

func NewMeDirectReportClientWithBaseURI(api sdkEnv.Api) (*MeDirectReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "medirectreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeDirectReportClient: %+v", err)
	}

	return &MeDirectReportClient{
		Client: client,
	}, nil
}
