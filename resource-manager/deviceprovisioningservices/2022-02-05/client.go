package v2022_02_05

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/dpscertificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-02-05/iotdpsresource"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	DpsCertificate *dpscertificate.DpsCertificateClient
	IotDpsResource *iotdpsresource.IotDpsResourceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dpsCertificateClient := dpscertificate.NewDpsCertificateClientWithBaseURI(endpoint)
	configureAuthFunc(&dpsCertificateClient.Client)

	iotDpsResourceClient := iotdpsresource.NewIotDpsResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&iotDpsResourceClient.Client)

	return Client{
		DpsCertificate: &dpsCertificateClient,
		IotDpsResource: &iotDpsResourceClient,
	}
}
