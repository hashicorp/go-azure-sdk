package conditionalaccestemplate

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

type ListConditionalAccesTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ConditionalAccessTemplate
}

type ListConditionalAccesTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ConditionalAccessTemplate
}

type ListConditionalAccesTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesTemplates ...
func (c ConditionalAccesTemplateClient) ListConditionalAccesTemplates(ctx context.Context) (result ListConditionalAccesTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesTemplatesCustomPager{},
		Path:       "/identity/conditionalAccess/templates",
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
		Values *[]stable.ConditionalAccessTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesTemplatesComplete retrieves all the results into a single object
func (c ConditionalAccesTemplateClient) ListConditionalAccesTemplatesComplete(ctx context.Context) (ListConditionalAccesTemplatesCompleteResult, error) {
	return c.ListConditionalAccesTemplatesCompleteMatchingPredicate(ctx, ConditionalAccessTemplateOperationPredicate{})
}

// ListConditionalAccesTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesTemplateClient) ListConditionalAccesTemplatesCompleteMatchingPredicate(ctx context.Context, predicate ConditionalAccessTemplateOperationPredicate) (result ListConditionalAccesTemplatesCompleteResult, err error) {
	items := make([]stable.ConditionalAccessTemplate, 0)

	resp, err := c.ListConditionalAccesTemplates(ctx)
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

	result = ListConditionalAccesTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
