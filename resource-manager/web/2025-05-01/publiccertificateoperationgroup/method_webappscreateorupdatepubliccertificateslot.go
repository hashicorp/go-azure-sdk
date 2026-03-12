package publiccertificateoperationgroup

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsCreateOrUpdatePublicCertificateSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *PublicCertificate
}

// WebAppsCreateOrUpdatePublicCertificateSlot ...
func (c PublicCertificateOperationGroupClient) WebAppsCreateOrUpdatePublicCertificateSlot(ctx context.Context, id SlotPublicCertificateId, input PublicCertificate) (result WebAppsCreateOrUpdatePublicCertificateSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

	var model PublicCertificate
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
