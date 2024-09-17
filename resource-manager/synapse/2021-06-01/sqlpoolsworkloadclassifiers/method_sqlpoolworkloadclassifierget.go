package sqlpoolsworkloadclassifiers

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolWorkloadClassifierGetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *WorkloadClassifier
}

// SqlPoolWorkloadClassifierGet ...
func (c SqlPoolsWorkloadClassifiersClient) SqlPoolWorkloadClassifierGet(ctx context.Context, id WorkloadClassifierId) (result SqlPoolWorkloadClassifierGetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model WorkloadClassifier
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
