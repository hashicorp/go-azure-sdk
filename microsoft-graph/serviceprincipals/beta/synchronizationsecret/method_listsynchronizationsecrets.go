package synchronizationsecret

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

type ListSynchronizationSecretsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.SynchronizationSecretKeyStringValuePair
}

type ListSynchronizationSecretsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.SynchronizationSecretKeyStringValuePair
}

type ListSynchronizationSecretsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListSynchronizationSecretsOperationOptions() ListSynchronizationSecretsOperationOptions {
	return ListSynchronizationSecretsOperationOptions{}
}

func (o ListSynchronizationSecretsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListSynchronizationSecretsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListSynchronizationSecretsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSynchronizationSecretsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSynchronizationSecretsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSynchronizationSecrets - Retrieve synchronization secrets.
func (c SynchronizationSecretClient) ListSynchronizationSecrets(ctx context.Context, id beta.ServicePrincipalId, options ListSynchronizationSecretsOperationOptions) (result ListSynchronizationSecretsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSynchronizationSecretsCustomPager{},
		Path:          fmt.Sprintf("%s/synchronization/secrets", id.ID()),
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
		Values *[]beta.SynchronizationSecretKeyStringValuePair `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSynchronizationSecretsComplete retrieves all the results into a single object
func (c SynchronizationSecretClient) ListSynchronizationSecretsComplete(ctx context.Context, id beta.ServicePrincipalId, options ListSynchronizationSecretsOperationOptions) (ListSynchronizationSecretsCompleteResult, error) {
	return c.ListSynchronizationSecretsCompleteMatchingPredicate(ctx, id, options, SynchronizationSecretKeyStringValuePairOperationPredicate{})
}

// ListSynchronizationSecretsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SynchronizationSecretClient) ListSynchronizationSecretsCompleteMatchingPredicate(ctx context.Context, id beta.ServicePrincipalId, options ListSynchronizationSecretsOperationOptions, predicate SynchronizationSecretKeyStringValuePairOperationPredicate) (result ListSynchronizationSecretsCompleteResult, err error) {
	items := make([]beta.SynchronizationSecretKeyStringValuePair, 0)

	resp, err := c.ListSynchronizationSecrets(ctx, id, options)
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

	result = ListSynchronizationSecretsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
