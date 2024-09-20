package regulatorycompliance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ControlsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RegulatoryComplianceControl
}

type ControlsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RegulatoryComplianceControl
}

type ControlsListOperationOptions struct {
	Filter *string
}

func DefaultControlsListOperationOptions() ControlsListOperationOptions {
	return ControlsListOperationOptions{}
}

func (o ControlsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ControlsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ControlsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type ControlsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ControlsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ControlsList ...
func (c RegulatoryComplianceClient) ControlsList(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions) (result ControlsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ControlsListCustomPager{},
		Path:          fmt.Sprintf("%s/regulatoryComplianceControls", id.ID()),
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
		Values *[]RegulatoryComplianceControl `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ControlsListComplete retrieves all the results into a single object
func (c RegulatoryComplianceClient) ControlsListComplete(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions) (ControlsListCompleteResult, error) {
	return c.ControlsListCompleteMatchingPredicate(ctx, id, options, RegulatoryComplianceControlOperationPredicate{})
}

// ControlsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RegulatoryComplianceClient) ControlsListCompleteMatchingPredicate(ctx context.Context, id RegulatoryComplianceStandardId, options ControlsListOperationOptions, predicate RegulatoryComplianceControlOperationPredicate) (result ControlsListCompleteResult, err error) {
	items := make([]RegulatoryComplianceControl, 0)

	resp, err := c.ControlsList(ctx, id, options)
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

	result = ControlsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
