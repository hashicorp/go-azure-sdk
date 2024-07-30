package b2xuserflowlanguagedefaultpage

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

type ListB2xUserFlowLanguageDefaultPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageDefaultPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageDefaultPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowLanguageDefaultPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowLanguageDefaultPages ...
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPages(ctx context.Context, id IdentityB2xUserFlowIdLanguageId) (result ListB2xUserFlowLanguageDefaultPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowLanguageDefaultPagesCustomPager{},
		Path:       fmt.Sprintf("%s/defaultPages", id.ID()),
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
		Values *[]stable.UserFlowLanguagePage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowLanguageDefaultPagesComplete retrieves all the results into a single object
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPagesComplete(ctx context.Context, id IdentityB2xUserFlowIdLanguageId) (ListB2xUserFlowLanguageDefaultPagesCompleteResult, error) {
	return c.ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate(ctx, id, UserFlowLanguagePageOperationPredicate{})
}

// ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowLanguageDefaultPageClient) ListB2xUserFlowLanguageDefaultPagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowIdLanguageId, predicate UserFlowLanguagePageOperationPredicate) (result ListB2xUserFlowLanguageDefaultPagesCompleteResult, err error) {
	items := make([]stable.UserFlowLanguagePage, 0)

	resp, err := c.ListB2xUserFlowLanguageDefaultPages(ctx, id)
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

	result = ListB2xUserFlowLanguageDefaultPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
