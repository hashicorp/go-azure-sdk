package datasecurityandgovernancesensitivitylabelright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSecurityAndGovernanceSensitivityLabelRightClient struct {
	Client *msgraph.Client
}

func NewDataSecurityAndGovernanceSensitivityLabelRightClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSecurityAndGovernanceSensitivityLabelRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "datasecurityandgovernancesensitivitylabelright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSecurityAndGovernanceSensitivityLabelRightClient: %+v", err)
	}

	return &DataSecurityAndGovernanceSensitivityLabelRightClient{
		Client: client,
	}, nil
}
