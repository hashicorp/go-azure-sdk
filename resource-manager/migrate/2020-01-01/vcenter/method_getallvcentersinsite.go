package vcenter

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAllVCentersInSiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VCenter
}

type GetAllVCentersInSiteCompleteResult struct {
	Items []VCenter
}

type GetAllVCentersInSiteOperationOptions struct {
	Filter *string
}

func DefaultGetAllVCentersInSiteOperationOptions() GetAllVCentersInSiteOperationOptions {
	return GetAllVCentersInSiteOperationOptions{}
}

func (o GetAllVCentersInSiteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAllVCentersInSiteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o GetAllVCentersInSiteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// GetAllVCentersInSite ...
func (c VCenterClient) GetAllVCentersInSite(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions) (result GetAllVCentersInSiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/vCenters", id.ID()),
		OptionsObject: options,
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
		Values *[]VCenter `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAllVCentersInSiteComplete retrieves all the results into a single object
func (c VCenterClient) GetAllVCentersInSiteComplete(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions) (GetAllVCentersInSiteCompleteResult, error) {
	return c.GetAllVCentersInSiteCompleteMatchingPredicate(ctx, id, options, VCenterOperationPredicate{})
}

// GetAllVCentersInSiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VCenterClient) GetAllVCentersInSiteCompleteMatchingPredicate(ctx context.Context, id VMwareSiteId, options GetAllVCentersInSiteOperationOptions, predicate VCenterOperationPredicate) (result GetAllVCentersInSiteCompleteResult, err error) {
	items := make([]VCenter, 0)

	resp, err := c.GetAllVCentersInSite(ctx, id, options)
	if err != nil {
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

	result = GetAllVCentersInSiteCompleteResult{
		Items: items,
	}
	return
}
