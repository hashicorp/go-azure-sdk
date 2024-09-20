package virtualendpointcloudpc

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateVirtualEndpointCloudPCsBulkResizesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCResizeValidationResult
}

type ValidateVirtualEndpointCloudPCsBulkResizesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCResizeValidationResult
}

type ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions struct {
	Metadata *odata.Metadata
	Skip     *int64
	Top      *int64
}

func DefaultValidateVirtualEndpointCloudPCsBulkResizesOperationOptions() ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions {
	return ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions{}
}

func (o ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ValidateVirtualEndpointCloudPCsBulkResizesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ValidateVirtualEndpointCloudPCsBulkResizesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ValidateVirtualEndpointCloudPCsBulkResizes - Invoke action validateBulkResize. Validate that a set of cloudPC devices
// meet the requirements to be bulk resized.
func (c VirtualEndpointCloudPCClient) ValidateVirtualEndpointCloudPCsBulkResizes(ctx context.Context, input ValidateVirtualEndpointCloudPCsBulkResizesRequest, options ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions) (result ValidateVirtualEndpointCloudPCsBulkResizesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ValidateVirtualEndpointCloudPCsBulkResizesCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/cloudPCs/validateBulkResize",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.CloudPCResizeValidationResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ValidateVirtualEndpointCloudPCsBulkResizesComplete retrieves all the results into a single object
func (c VirtualEndpointCloudPCClient) ValidateVirtualEndpointCloudPCsBulkResizesComplete(ctx context.Context, input ValidateVirtualEndpointCloudPCsBulkResizesRequest, options ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions) (ValidateVirtualEndpointCloudPCsBulkResizesCompleteResult, error) {
	return c.ValidateVirtualEndpointCloudPCsBulkResizesCompleteMatchingPredicate(ctx, input, options, CloudPCResizeValidationResultOperationPredicate{})
}

// ValidateVirtualEndpointCloudPCsBulkResizesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointCloudPCClient) ValidateVirtualEndpointCloudPCsBulkResizesCompleteMatchingPredicate(ctx context.Context, input ValidateVirtualEndpointCloudPCsBulkResizesRequest, options ValidateVirtualEndpointCloudPCsBulkResizesOperationOptions, predicate CloudPCResizeValidationResultOperationPredicate) (result ValidateVirtualEndpointCloudPCsBulkResizesCompleteResult, err error) {
	items := make([]beta.CloudPCResizeValidationResult, 0)

	resp, err := c.ValidateVirtualEndpointCloudPCsBulkResizes(ctx, input, options)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ValidateVirtualEndpointCloudPCsBulkResizesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
