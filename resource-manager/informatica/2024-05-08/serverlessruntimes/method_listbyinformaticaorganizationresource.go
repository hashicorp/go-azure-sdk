package serverlessruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInformaticaOrganizationResourceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]InformaticaServerlessRuntimeResource
}

type ListByInformaticaOrganizationResourceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []InformaticaServerlessRuntimeResource
}

type ListByInformaticaOrganizationResourceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByInformaticaOrganizationResourceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByInformaticaOrganizationResource ...
func (c ServerlessRuntimesClient) ListByInformaticaOrganizationResource(ctx context.Context, id OrganizationId) (result ListByInformaticaOrganizationResourceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByInformaticaOrganizationResourceCustomPager{},
		Path:       fmt.Sprintf("%s/serverlessRuntimes", id.ID()),
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
		Values *[]InformaticaServerlessRuntimeResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInformaticaOrganizationResourceComplete retrieves all the results into a single object
func (c ServerlessRuntimesClient) ListByInformaticaOrganizationResourceComplete(ctx context.Context, id OrganizationId) (ListByInformaticaOrganizationResourceCompleteResult, error) {
	return c.ListByInformaticaOrganizationResourceCompleteMatchingPredicate(ctx, id, InformaticaServerlessRuntimeResourceOperationPredicate{})
}

// ListByInformaticaOrganizationResourceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServerlessRuntimesClient) ListByInformaticaOrganizationResourceCompleteMatchingPredicate(ctx context.Context, id OrganizationId, predicate InformaticaServerlessRuntimeResourceOperationPredicate) (result ListByInformaticaOrganizationResourceCompleteResult, err error) {
	items := make([]InformaticaServerlessRuntimeResource, 0)

	resp, err := c.ListByInformaticaOrganizationResource(ctx, id)
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

	result = ListByInformaticaOrganizationResourceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
