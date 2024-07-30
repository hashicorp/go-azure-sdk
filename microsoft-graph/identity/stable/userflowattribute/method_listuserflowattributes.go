package userflowattribute

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

type ListUserFlowAttributesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.IdentityUserFlowAttribute
}

type ListUserFlowAttributesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.IdentityUserFlowAttribute
}

type ListUserFlowAttributesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListUserFlowAttributesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListUserFlowAttributes ...
func (c UserFlowAttributeClient) ListUserFlowAttributes(ctx context.Context) (result ListUserFlowAttributesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListUserFlowAttributesCustomPager{},
		Path:       "/identity/userFlowAttributes",
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
		Values *[]stable.IdentityUserFlowAttribute `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserFlowAttributesComplete retrieves all the results into a single object
func (c UserFlowAttributeClient) ListUserFlowAttributesComplete(ctx context.Context) (ListUserFlowAttributesCompleteResult, error) {
	return c.ListUserFlowAttributesCompleteMatchingPredicate(ctx, IdentityUserFlowAttributeOperationPredicate{})
}

// ListUserFlowAttributesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserFlowAttributeClient) ListUserFlowAttributesCompleteMatchingPredicate(ctx context.Context, predicate IdentityUserFlowAttributeOperationPredicate) (result ListUserFlowAttributesCompleteResult, err error) {
	items := make([]stable.IdentityUserFlowAttribute, 0)

	resp, err := c.ListUserFlowAttributes(ctx)
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

	result = ListUserFlowAttributesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
