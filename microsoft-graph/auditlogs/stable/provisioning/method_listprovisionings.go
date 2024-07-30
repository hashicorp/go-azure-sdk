package provisioning

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListProvisioningsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ProvisioningObjectSummary
}

type ListProvisioningsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ProvisioningObjectSummary
}

type ListProvisioningsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProvisioningsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProvisionings ...
func (c ProvisioningClient) ListProvisionings(ctx context.Context) (result ListProvisioningsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProvisioningsCustomPager{},
		Path:       "/auditLogs/provisioning",
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
		Values *[]stable.ProvisioningObjectSummary `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProvisioningsComplete retrieves all the results into a single object
func (c ProvisioningClient) ListProvisioningsComplete(ctx context.Context) (ListProvisioningsCompleteResult, error) {
	return c.ListProvisioningsCompleteMatchingPredicate(ctx, ProvisioningObjectSummaryOperationPredicate{})
}

// ListProvisioningsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProvisioningClient) ListProvisioningsCompleteMatchingPredicate(ctx context.Context, predicate ProvisioningObjectSummaryOperationPredicate) (result ListProvisioningsCompleteResult, err error) {
	items := make([]stable.ProvisioningObjectSummary, 0)

	resp, err := c.ListProvisionings(ctx)
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

	result = ListProvisioningsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
