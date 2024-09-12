package me

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportPersonalDataOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// ExportPersonalData - Invoke action exportPersonalData. Submit a data policy operation request from a company
// administrator or an application to export an organizational user's data. This data includes the user's data stored in
// OneDrive and their activity reports. For more information about exporting data while complying with regulations, see
// Data Subject Requests and the GDPR and CCPA.
func (c MeClient) ExportPersonalData(ctx context.Context, input ExportPersonalDataRequest) (result ExportPersonalDataOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/me/exportPersonalData",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
