package remoteactionaudit

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

type ListRemoteActionAuditsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RemoteActionAudit
}

type ListRemoteActionAuditsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RemoteActionAudit
}

type ListRemoteActionAuditsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRemoteActionAuditsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRemoteActionAudits ...
func (c RemoteActionAuditClient) ListRemoteActionAudits(ctx context.Context) (result ListRemoteActionAuditsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRemoteActionAuditsCustomPager{},
		Path:       "/deviceManagement/remoteActionAudits",
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
		Values *[]beta.RemoteActionAudit `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRemoteActionAuditsComplete retrieves all the results into a single object
func (c RemoteActionAuditClient) ListRemoteActionAuditsComplete(ctx context.Context) (ListRemoteActionAuditsCompleteResult, error) {
	return c.ListRemoteActionAuditsCompleteMatchingPredicate(ctx, RemoteActionAuditOperationPredicate{})
}

// ListRemoteActionAuditsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RemoteActionAuditClient) ListRemoteActionAuditsCompleteMatchingPredicate(ctx context.Context, predicate RemoteActionAuditOperationPredicate) (result ListRemoteActionAuditsCompleteResult, err error) {
	items := make([]beta.RemoteActionAudit, 0)

	resp, err := c.ListRemoteActionAudits(ctx)
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

	result = ListRemoteActionAuditsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
