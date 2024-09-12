package user

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

type GetPasswordSingleSignOnCredentialsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.PasswordSingleSignOnCredentialSet
}

type GetPasswordSingleSignOnCredentialsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.PasswordSingleSignOnCredentialSet
}

type GetPasswordSingleSignOnCredentialsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultGetPasswordSingleSignOnCredentialsOperationOptions() GetPasswordSingleSignOnCredentialsOperationOptions {
	return GetPasswordSingleSignOnCredentialsOperationOptions{}
}

func (o GetPasswordSingleSignOnCredentialsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetPasswordSingleSignOnCredentialsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o GetPasswordSingleSignOnCredentialsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type GetPasswordSingleSignOnCredentialsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *GetPasswordSingleSignOnCredentialsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetPasswordSingleSignOnCredentials - Invoke action getPasswordSingleSignOnCredentials. Get the list of password-based
// single sign-on credentials for a given user. This API returns the encrypted passwords as null or empty strings.
func (c UserClient) GetPasswordSingleSignOnCredentials(ctx context.Context, id beta.UserId, options GetPasswordSingleSignOnCredentialsOperationOptions) (result GetPasswordSingleSignOnCredentialsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &GetPasswordSingleSignOnCredentialsCustomPager{},
		Path:          fmt.Sprintf("%s/getPasswordSingleSignOnCredentials", id.ID()),
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
		Values *[]beta.PasswordSingleSignOnCredentialSet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetPasswordSingleSignOnCredentialsComplete retrieves all the results into a single object
func (c UserClient) GetPasswordSingleSignOnCredentialsComplete(ctx context.Context, id beta.UserId, options GetPasswordSingleSignOnCredentialsOperationOptions) (GetPasswordSingleSignOnCredentialsCompleteResult, error) {
	return c.GetPasswordSingleSignOnCredentialsCompleteMatchingPredicate(ctx, id, options, PasswordSingleSignOnCredentialSetOperationPredicate{})
}

// GetPasswordSingleSignOnCredentialsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserClient) GetPasswordSingleSignOnCredentialsCompleteMatchingPredicate(ctx context.Context, id beta.UserId, options GetPasswordSingleSignOnCredentialsOperationOptions, predicate PasswordSingleSignOnCredentialSetOperationPredicate) (result GetPasswordSingleSignOnCredentialsCompleteResult, err error) {
	items := make([]beta.PasswordSingleSignOnCredentialSet, 0)

	resp, err := c.GetPasswordSingleSignOnCredentials(ctx, id, options)
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

	result = GetPasswordSingleSignOnCredentialsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
