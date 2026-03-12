package triggers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByImageTemplateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Trigger
}

type ListByImageTemplateCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Trigger
}

type ListByImageTemplateCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByImageTemplateCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByImageTemplate ...
func (c TriggersClient) ListByImageTemplate(ctx context.Context, id ImageTemplateId) (result ListByImageTemplateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListByImageTemplateCustomPager{},
		Path:       fmt.Sprintf("%s/triggers", id.ID()),
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
		Values *[]Trigger `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByImageTemplateComplete retrieves all the results into a single object
func (c TriggersClient) ListByImageTemplateComplete(ctx context.Context, id ImageTemplateId) (ListByImageTemplateCompleteResult, error) {
	return c.ListByImageTemplateCompleteMatchingPredicate(ctx, id, TriggerOperationPredicate{})
}

// ListByImageTemplateCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggersClient) ListByImageTemplateCompleteMatchingPredicate(ctx context.Context, id ImageTemplateId, predicate TriggerOperationPredicate) (result ListByImageTemplateCompleteResult, err error) {
	items := make([]Trigger, 0)

	resp, err := c.ListByImageTemplate(ctx, id)
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

	result = ListByImageTemplateCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
