package usertransitivereport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTransitiveReportClient struct {
	Client *msgraph.Client
}

func NewUserTransitiveReportClientWithBaseURI(api sdkEnv.Api) (*UserTransitiveReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usertransitivereport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserTransitiveReportClient: %+v", err)
	}

	return &UserTransitiveReportClient{
		Client: client,
	}, nil
}
