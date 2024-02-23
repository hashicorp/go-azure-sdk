package listqnamakerendpointkeys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QnAMakerEndpointKeysGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *QnAMakerEndpointKeysResponse
}

// QnAMakerEndpointKeysGet ...
func (c ListQnAMakerEndpointKeysClient) QnAMakerEndpointKeysGet(ctx context.Context, id commonids.SubscriptionId, input QnAMakerEndpointKeysRequestBody) (result QnAMakerEndpointKeysGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/providers/Microsoft.BotService/listQnAMakerEndpointKeys", id.ID()),
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

	var model QnAMakerEndpointKeysResponse
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
