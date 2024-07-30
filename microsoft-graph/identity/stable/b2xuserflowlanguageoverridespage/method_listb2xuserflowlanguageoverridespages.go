package b2xuserflowlanguageoverridespage

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

type ListB2xUserFlowLanguageOverridesPagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageOverridesPagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserFlowLanguagePage
}

type ListB2xUserFlowLanguageOverridesPagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowLanguageOverridesPagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowLanguageOverridesPages ...
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPages(ctx context.Context, id IdentityB2xUserFlowIdLanguageId) (result ListB2xUserFlowLanguageOverridesPagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowLanguageOverridesPagesCustomPager{},
		Path:       fmt.Sprintf("%s/overridesPages", id.ID()),
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

// ListB2xUserFlowLanguageOverridesPagesComplete retrieves all the results into a single object
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPagesComplete(ctx context.Context, id IdentityB2xUserFlowIdLanguageId) (ListB2xUserFlowLanguageOverridesPagesCompleteResult, error) {
	return c.ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate(ctx, id, UserFlowLanguagePageOperationPredicate{})
}

// ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowLanguageOverridesPageClient) ListB2xUserFlowLanguageOverridesPagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowIdLanguageId, predicate UserFlowLanguagePageOperationPredicate) (result ListB2xUserFlowLanguageOverridesPagesCompleteResult, err error) {
	items := make([]stable.UserFlowLanguagePage, 0)

	resp, err := c.ListB2xUserFlowLanguageOverridesPages(ctx, id)
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

	result = ListB2xUserFlowLanguageOverridesPagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
