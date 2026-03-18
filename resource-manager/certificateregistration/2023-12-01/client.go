package v2023_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/certificateregistration/2023-12-01/appservicecertificateorders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/certificateregistration/2023-12-01/certificateordersdiagnostics"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppServiceCertificateOrders  *appservicecertificateorders.AppServiceCertificateOrdersClient
	CertificateOrdersDiagnostics *certificateordersdiagnostics.CertificateOrdersDiagnosticsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appServiceCertificateOrdersClient, err := appservicecertificateorders.NewAppServiceCertificateOrdersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServiceCertificateOrders client: %+v", err)
	}
	configureFunc(appServiceCertificateOrdersClient.Client)

	certificateOrdersDiagnosticsClient, err := certificateordersdiagnostics.NewCertificateOrdersDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateOrdersDiagnostics client: %+v", err)
	}
	configureFunc(certificateOrdersDiagnosticsClient.Client)

	return &Client{
		AppServiceCertificateOrders:  appServiceCertificateOrdersClient,
		CertificateOrdersDiagnostics: certificateOrdersDiagnosticsClient,
	}, nil
}
