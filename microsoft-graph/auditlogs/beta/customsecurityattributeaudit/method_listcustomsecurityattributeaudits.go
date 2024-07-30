package customsecurityattributeaudit

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

type ListCustomSecurityAttributeAuditsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CustomSecurityAttributeAudit
}

type ListCustomSecurityAttributeAuditsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CustomSecurityAttributeAudit
}

type ListCustomSecurityAttributeAuditsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCustomSecurityAttributeAuditsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCustomSecurityAttributeAudits ...
func (c CustomSecurityAttributeAuditClient) ListCustomSecurityAttributeAudits(ctx context.Context) (result ListCustomSecurityAttributeAuditsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListCustomSecurityAttributeAuditsCustomPager{},
		Path:       "/auditLogs/customSecurityAttributeAudits",
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
		Values *[]beta.CustomSecurityAttributeAudit `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListCustomSecurityAttributeAuditsComplete retrieves all the results into a single object
func (c CustomSecurityAttributeAuditClient) ListCustomSecurityAttributeAuditsComplete(ctx context.Context) (ListCustomSecurityAttributeAuditsCompleteResult, error) {
	return c.ListCustomSecurityAttributeAuditsCompleteMatchingPredicate(ctx, CustomSecurityAttributeAuditOperationPredicate{})
}

// ListCustomSecurityAttributeAuditsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CustomSecurityAttributeAuditClient) ListCustomSecurityAttributeAuditsCompleteMatchingPredicate(ctx context.Context, predicate CustomSecurityAttributeAuditOperationPredicate) (result ListCustomSecurityAttributeAuditsCompleteResult, err error) {
	items := make([]beta.CustomSecurityAttributeAudit, 0)

	resp, err := c.ListCustomSecurityAttributeAudits(ctx)
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

	result = ListCustomSecurityAttributeAuditsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
