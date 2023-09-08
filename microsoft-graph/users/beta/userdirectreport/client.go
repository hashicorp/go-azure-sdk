package userdirectreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserDirectReportClient struct {
	Client *msgraph.Client
}

func NewUserDirectReportClientWithBaseURI(api sdkEnv.Api) (*UserDirectReportClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userdirectreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserDirectReportClient: %+v", err)
	}

	return &UserDirectReportClient{
		Client: client,
	}, nil
}
