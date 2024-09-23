package compliancecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceCategoryClient struct {
	Client *msgraph.Client
}

func NewComplianceCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ComplianceCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComplianceCategoryClient: %+v", err)
	}

	return &ComplianceCategoryClient{
		Client: client,
	}, nil
}
