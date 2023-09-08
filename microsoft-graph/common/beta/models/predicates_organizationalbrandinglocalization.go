package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationalBrandingLocalizationOperationPredicate struct {
	BackgroundColor                   *string
	BackgroundImage                   *string
	BackgroundImageRelativeUrl        *string
	BannerLogo                        *string
	BannerLogoRelativeUrl             *string
	CustomAccountResetCredentialsUrl  *string
	CustomCSS                         *string
	CustomCSSRelativeUrl              *string
	CustomCannotAccessYourAccountText *string
	CustomCannotAccessYourAccountUrl  *string
	CustomForgotMyPasswordText        *string
	CustomPrivacyAndCookiesText       *string
	CustomPrivacyAndCookiesUrl        *string
	CustomResetItNowText              *string
	CustomTermsOfUseText              *string
	CustomTermsOfUseUrl               *string
	Favicon                           *string
	FaviconRelativeUrl                *string
	HeaderBackgroundColor             *string
	HeaderLogo                        *string
	HeaderLogoRelativeUrl             *string
	Id                                *string
	ODataType                         *string
	SignInPageText                    *string
	SquareLogo                        *string
	SquareLogoDark                    *string
	SquareLogoDarkRelativeUrl         *string
	SquareLogoRelativeUrl             *string
	UsernameHintText                  *string
}

func (p OrganizationalBrandingLocalizationOperationPredicate) Matches(input OrganizationalBrandingLocalization) bool {

	if p.BackgroundColor != nil && (input.BackgroundColor == nil || *p.BackgroundColor != *input.BackgroundColor) {
		return false
	}

	if p.BackgroundImage != nil && (input.BackgroundImage == nil || *p.BackgroundImage != *input.BackgroundImage) {
		return false
	}

	if p.BackgroundImageRelativeUrl != nil && (input.BackgroundImageRelativeUrl == nil || *p.BackgroundImageRelativeUrl != *input.BackgroundImageRelativeUrl) {
		return false
	}

	if p.BannerLogo != nil && (input.BannerLogo == nil || *p.BannerLogo != *input.BannerLogo) {
		return false
	}

	if p.BannerLogoRelativeUrl != nil && (input.BannerLogoRelativeUrl == nil || *p.BannerLogoRelativeUrl != *input.BannerLogoRelativeUrl) {
		return false
	}

	if p.CustomAccountResetCredentialsUrl != nil && (input.CustomAccountResetCredentialsUrl == nil || *p.CustomAccountResetCredentialsUrl != *input.CustomAccountResetCredentialsUrl) {
		return false
	}

	if p.CustomCSS != nil && (input.CustomCSS == nil || *p.CustomCSS != *input.CustomCSS) {
		return false
	}

	if p.CustomCSSRelativeUrl != nil && (input.CustomCSSRelativeUrl == nil || *p.CustomCSSRelativeUrl != *input.CustomCSSRelativeUrl) {
		return false
	}

	if p.CustomCannotAccessYourAccountText != nil && (input.CustomCannotAccessYourAccountText == nil || *p.CustomCannotAccessYourAccountText != *input.CustomCannotAccessYourAccountText) {
		return false
	}

	if p.CustomCannotAccessYourAccountUrl != nil && (input.CustomCannotAccessYourAccountUrl == nil || *p.CustomCannotAccessYourAccountUrl != *input.CustomCannotAccessYourAccountUrl) {
		return false
	}

	if p.CustomForgotMyPasswordText != nil && (input.CustomForgotMyPasswordText == nil || *p.CustomForgotMyPasswordText != *input.CustomForgotMyPasswordText) {
		return false
	}

	if p.CustomPrivacyAndCookiesText != nil && (input.CustomPrivacyAndCookiesText == nil || *p.CustomPrivacyAndCookiesText != *input.CustomPrivacyAndCookiesText) {
		return false
	}

	if p.CustomPrivacyAndCookiesUrl != nil && (input.CustomPrivacyAndCookiesUrl == nil || *p.CustomPrivacyAndCookiesUrl != *input.CustomPrivacyAndCookiesUrl) {
		return false
	}

	if p.CustomResetItNowText != nil && (input.CustomResetItNowText == nil || *p.CustomResetItNowText != *input.CustomResetItNowText) {
		return false
	}

	if p.CustomTermsOfUseText != nil && (input.CustomTermsOfUseText == nil || *p.CustomTermsOfUseText != *input.CustomTermsOfUseText) {
		return false
	}

	if p.CustomTermsOfUseUrl != nil && (input.CustomTermsOfUseUrl == nil || *p.CustomTermsOfUseUrl != *input.CustomTermsOfUseUrl) {
		return false
	}

	if p.Favicon != nil && (input.Favicon == nil || *p.Favicon != *input.Favicon) {
		return false
	}

	if p.FaviconRelativeUrl != nil && (input.FaviconRelativeUrl == nil || *p.FaviconRelativeUrl != *input.FaviconRelativeUrl) {
		return false
	}

	if p.HeaderBackgroundColor != nil && (input.HeaderBackgroundColor == nil || *p.HeaderBackgroundColor != *input.HeaderBackgroundColor) {
		return false
	}

	if p.HeaderLogo != nil && (input.HeaderLogo == nil || *p.HeaderLogo != *input.HeaderLogo) {
		return false
	}

	if p.HeaderLogoRelativeUrl != nil && (input.HeaderLogoRelativeUrl == nil || *p.HeaderLogoRelativeUrl != *input.HeaderLogoRelativeUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SignInPageText != nil && (input.SignInPageText == nil || *p.SignInPageText != *input.SignInPageText) {
		return false
	}

	if p.SquareLogo != nil && (input.SquareLogo == nil || *p.SquareLogo != *input.SquareLogo) {
		return false
	}

	if p.SquareLogoDark != nil && (input.SquareLogoDark == nil || *p.SquareLogoDark != *input.SquareLogoDark) {
		return false
	}

	if p.SquareLogoDarkRelativeUrl != nil && (input.SquareLogoDarkRelativeUrl == nil || *p.SquareLogoDarkRelativeUrl != *input.SquareLogoDarkRelativeUrl) {
		return false
	}

	if p.SquareLogoRelativeUrl != nil && (input.SquareLogoRelativeUrl == nil || *p.SquareLogoRelativeUrl != *input.SquareLogoRelativeUrl) {
		return false
	}

	if p.UsernameHintText != nil && (input.UsernameHintText == nil || *p.UsernameHintText != *input.UsernameHintText) {
		return false
	}

	return true
}
