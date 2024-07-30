package b2xuserflowlanguage

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

type ListB2xUserFlowLanguagesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UserFlowLanguageConfiguration
}

type ListB2xUserFlowLanguagesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UserFlowLanguageConfiguration
}

type ListB2xUserFlowLanguagesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListB2xUserFlowLanguagesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListB2xUserFlowLanguages ...
func (c B2xUserFlowLanguageClient) ListB2xUserFlowLanguages(ctx context.Context, id IdentityB2xUserFlowId) (result ListB2xUserFlowLanguagesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListB2xUserFlowLanguagesCustomPager{},
		Path:       fmt.Sprintf("%s/languages", id.ID()),
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
		Values *[]stable.UserFlowLanguageConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListB2xUserFlowLanguagesComplete retrieves all the results into a single object
func (c B2xUserFlowLanguageClient) ListB2xUserFlowLanguagesComplete(ctx context.Context, id IdentityB2xUserFlowId) (ListB2xUserFlowLanguagesCompleteResult, error) {
	return c.ListB2xUserFlowLanguagesCompleteMatchingPredicate(ctx, id, UserFlowLanguageConfigurationOperationPredicate{})
}

// ListB2xUserFlowLanguagesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c B2xUserFlowLanguageClient) ListB2xUserFlowLanguagesCompleteMatchingPredicate(ctx context.Context, id IdentityB2xUserFlowId, predicate UserFlowLanguageConfigurationOperationPredicate) (result ListB2xUserFlowLanguagesCompleteResult, err error) {
	items := make([]stable.UserFlowLanguageConfiguration, 0)

	resp, err := c.ListB2xUserFlowLanguages(ctx, id)
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

	result = ListB2xUserFlowLanguagesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
