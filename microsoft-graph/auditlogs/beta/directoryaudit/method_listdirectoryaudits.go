package directoryaudit

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

type ListDirectoryAuditsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.DirectoryAudit
}

type ListDirectoryAuditsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.DirectoryAudit
}

type ListDirectoryAuditsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryAuditsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryAudits ...
func (c DirectoryAuditClient) ListDirectoryAudits(ctx context.Context) (result ListDirectoryAuditsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListDirectoryAuditsCustomPager{},
		Path:       "/auditLogs/directoryAudits",
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
		Values *[]beta.DirectoryAudit `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryAuditsComplete retrieves all the results into a single object
func (c DirectoryAuditClient) ListDirectoryAuditsComplete(ctx context.Context) (ListDirectoryAuditsCompleteResult, error) {
	return c.ListDirectoryAuditsCompleteMatchingPredicate(ctx, DirectoryAuditOperationPredicate{})
}

// ListDirectoryAuditsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryAuditClient) ListDirectoryAuditsCompleteMatchingPredicate(ctx context.Context, predicate DirectoryAuditOperationPredicate) (result ListDirectoryAuditsCompleteResult, err error) {
	items := make([]beta.DirectoryAudit, 0)

	resp, err := c.ListDirectoryAudits(ctx)
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

	result = ListDirectoryAuditsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
