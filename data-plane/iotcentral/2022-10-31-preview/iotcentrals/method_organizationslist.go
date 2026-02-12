package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Organization
}

type OrganizationsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Organization
}

type OrganizationsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *OrganizationsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// OrganizationsList ...
func (c IotcentralsClient) OrganizationsList(ctx context.Context) (result OrganizationsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &OrganizationsListCustomPager{},
		Path:       "/organizations",
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
		Values *[]Organization `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// OrganizationsListComplete retrieves all the results into a single object
func (c IotcentralsClient) OrganizationsListComplete(ctx context.Context) (OrganizationsListCompleteResult, error) {
	return c.OrganizationsListCompleteMatchingPredicate(ctx, OrganizationOperationPredicate{})
}

// OrganizationsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) OrganizationsListCompleteMatchingPredicate(ctx context.Context, predicate OrganizationOperationPredicate) (result OrganizationsListCompleteResult, err error) {
	items := make([]Organization, 0)

	resp, err := c.OrganizationsList(ctx)
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

	result = OrganizationsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
