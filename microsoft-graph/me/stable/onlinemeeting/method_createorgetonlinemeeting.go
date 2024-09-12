package onlinemeeting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrGetOnlineMeetingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.OnlineMeeting
}

// CreateOrGetOnlineMeeting - Invoke action createOrGet. Create an onlineMeeting object with a custom specified external
// ID. If the external ID already exists, this API will return the onlineMeeting object with that external ID.
func (c OnlineMeetingClient) CreateOrGetOnlineMeeting(ctx context.Context, input CreateOrGetOnlineMeetingRequest) (result CreateOrGetOnlineMeetingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/me/onlineMeetings/createOrGet",
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

	var model stable.OnlineMeeting
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
