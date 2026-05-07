package openapis

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

type AttestationsListForResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Attestation
}

type AttestationsListForResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Attestation
}

type AttestationsListForResourceGroupOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultAttestationsListForResourceGroupOperationOptions() AttestationsListForResourceGroupOperationOptions {
	return AttestationsListForResourceGroupOperationOptions{}
}

func (o AttestationsListForResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AttestationsListForResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AttestationsListForResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type AttestationsListForResourceGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *AttestationsListForResourceGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// AttestationsListForResourceGroup ...
func (c OpenapisClient) AttestationsListForResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options AttestationsListForResourceGroupOperationOptions) (result AttestationsListForResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &AttestationsListForResourceGroupCustomPager{},
		Path:          fmt.Sprintf("%s/providers/Microsoft.PolicyInsights/attestations", id.ID()),
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
		Values *[]Attestation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// AttestationsListForResourceGroupComplete retrieves all the results into a single object
func (c OpenapisClient) AttestationsListForResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options AttestationsListForResourceGroupOperationOptions) (AttestationsListForResourceGroupCompleteResult, error) {
	return c.AttestationsListForResourceGroupCompleteMatchingPredicate(ctx, id, options, AttestationOperationPredicate{})
}

// AttestationsListForResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) AttestationsListForResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options AttestationsListForResourceGroupOperationOptions, predicate AttestationOperationPredicate) (result AttestationsListForResourceGroupCompleteResult, err error) {
	items := make([]Attestation, 0)

	resp, err := c.AttestationsListForResourceGroup(ctx, id, options)
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

	result = AttestationsListForResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
