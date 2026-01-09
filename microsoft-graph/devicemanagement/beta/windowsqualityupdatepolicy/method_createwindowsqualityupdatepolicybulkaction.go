package windowsqualityupdatepolicy

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsQualityUpdatePolicyBulkActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.BulkCatalogItemActionResult
}

type CreateWindowsQualityUpdatePolicyBulkActionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsQualityUpdatePolicyBulkActionOperationOptions() CreateWindowsQualityUpdatePolicyBulkActionOperationOptions {
	return CreateWindowsQualityUpdatePolicyBulkActionOperationOptions{}
}

func (o CreateWindowsQualityUpdatePolicyBulkActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsQualityUpdatePolicyBulkActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsQualityUpdatePolicyBulkActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsQualityUpdatePolicyBulkAction - Invoke action bulkAction
func (c WindowsQualityUpdatePolicyClient) CreateWindowsQualityUpdatePolicyBulkAction(ctx context.Context, id beta.DeviceManagementWindowsQualityUpdatePolicyId, input CreateWindowsQualityUpdatePolicyBulkActionRequest, options CreateWindowsQualityUpdatePolicyBulkActionOperationOptions) (result CreateWindowsQualityUpdatePolicyBulkActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/bulkAction", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model beta.BulkCatalogItemActionResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
