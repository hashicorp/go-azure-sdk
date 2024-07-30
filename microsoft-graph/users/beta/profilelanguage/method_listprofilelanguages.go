package profilelanguage

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

type ListProfileLanguagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.LanguageProficiency
}

type ListProfileLanguagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.LanguageProficiency
}

type ListProfileLanguagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListProfileLanguagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListProfileLanguages ...
func (c ProfileLanguageClient) ListProfileLanguages(ctx context.Context, id UserId) (result ListProfileLanguagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListProfileLanguagesCustomPager{},
		Path:       fmt.Sprintf("%s/profile/languages", id.ID()),
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
		Values *[]beta.LanguageProficiency `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListProfileLanguagesComplete retrieves all the results into a single object
func (c ProfileLanguageClient) ListProfileLanguagesComplete(ctx context.Context, id UserId) (ListProfileLanguagesCompleteResult, error) {
	return c.ListProfileLanguagesCompleteMatchingPredicate(ctx, id, LanguageProficiencyOperationPredicate{})
}

// ListProfileLanguagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProfileLanguageClient) ListProfileLanguagesCompleteMatchingPredicate(ctx context.Context, id UserId, predicate LanguageProficiencyOperationPredicate) (result ListProfileLanguagesCompleteResult, err error) {
	items := make([]beta.LanguageProficiency, 0)

	resp, err := c.ListProfileLanguages(ctx, id)
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

	result = ListProfileLanguagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
