package synchronizationtemplate

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

type ListSynchronizationTemplatesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SynchronizationTemplate
}

type ListSynchronizationTemplatesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SynchronizationTemplate
}

type ListSynchronizationTemplatesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationTemplatesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationTemplates ...
func (c SynchronizationTemplateClient) ListSynchronizationTemplates(ctx context.Context, id ServicePrincipalId) (result ListSynchronizationTemplatesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListSynchronizationTemplatesCustomPager{},
		Path:       fmt.Sprintf("%s/synchronization/templates", id.ID()),
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
		Values *[]beta.SynchronizationTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationTemplatesComplete retrieves all the results into a single object
func (c SynchronizationTemplateClient) ListSynchronizationTemplatesComplete(ctx context.Context, id ServicePrincipalId) (ListSynchronizationTemplatesCompleteResult, error) {
	return c.ListSynchronizationTemplatesCompleteMatchingPredicate(ctx, id, SynchronizationTemplateOperationPredicate{})
}

// ListSynchronizationTemplatesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationTemplateClient) ListSynchronizationTemplatesCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate SynchronizationTemplateOperationPredicate) (result ListSynchronizationTemplatesCompleteResult, err error) {
	items := make([]beta.SynchronizationTemplate, 0)

	resp, err := c.ListSynchronizationTemplates(ctx, id)
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

	result = ListSynchronizationTemplatesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
