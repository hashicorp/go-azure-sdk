package functionenvelopes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListFunctionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]FunctionEnvelope
}

type WebAppsListFunctionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []FunctionEnvelope
}

type WebAppsListFunctionsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListFunctionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListFunctions ...
func (c FunctionEnvelopesClient) WebAppsListFunctions(ctx context.Context, id commonids.AppServiceId) (result WebAppsListFunctionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListFunctionsCustomPager{},
		Path:       fmt.Sprintf("%s/functions", id.ID()),
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
		Values *[]FunctionEnvelope `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListFunctionsComplete retrieves all the results into a single object
func (c FunctionEnvelopesClient) WebAppsListFunctionsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListFunctionsCompleteResult, error) {
	return c.WebAppsListFunctionsCompleteMatchingPredicate(ctx, id, FunctionEnvelopeOperationPredicate{})
}

// WebAppsListFunctionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c FunctionEnvelopesClient) WebAppsListFunctionsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate FunctionEnvelopeOperationPredicate) (result WebAppsListFunctionsCompleteResult, err error) {
	items := make([]FunctionEnvelope, 0)

	resp, err := c.WebAppsListFunctions(ctx, id)
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

	result = WebAppsListFunctionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
