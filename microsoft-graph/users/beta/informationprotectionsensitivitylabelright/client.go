package informationprotectionsensitivitylabelright

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionSensitivityLabelRightClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionSensitivityLabelRightClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionSensitivityLabelRightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionsensitivitylabelright", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionSensitivityLabelRightClient: %+v", err)
	}

	return &InformationProtectionSensitivityLabelRightClient{
		Client: client,
	}, nil
}
