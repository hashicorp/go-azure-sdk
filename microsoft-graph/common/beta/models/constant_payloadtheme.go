package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadTheme string

const (
	PayloadThemeaccountActivation     PayloadTheme = "AccountActivation"
	PayloadThemeaccountVerification   PayloadTheme = "AccountVerification"
	PayloadThemeadvertisement         PayloadTheme = "Advertisement"
	PayloadThemebilling               PayloadTheme = "Billing"
	PayloadThemecleanUpMail           PayloadTheme = "CleanUpMail"
	PayloadThemecontroversial         PayloadTheme = "Controversial"
	PayloadThemedocumentReceived      PayloadTheme = "DocumentReceived"
	PayloadThemeemployeeEngagement    PayloadTheme = "EmployeeEngagement"
	PayloadThemeexpense               PayloadTheme = "Expense"
	PayloadThemefax                   PayloadTheme = "Fax"
	PayloadThemefinanceReport         PayloadTheme = "FinanceReport"
	PayloadThemeincomingMessages      PayloadTheme = "IncomingMessages"
	PayloadThemeinvoice               PayloadTheme = "Invoice"
	PayloadThemeitemReceived          PayloadTheme = "ItemReceived"
	PayloadThemeloginAlert            PayloadTheme = "LoginAlert"
	PayloadThememailReceived          PayloadTheme = "MailReceived"
	PayloadThemeother                 PayloadTheme = "Other"
	PayloadThemepassword              PayloadTheme = "Password"
	PayloadThemepayment               PayloadTheme = "Payment"
	PayloadThemepayroll               PayloadTheme = "Payroll"
	PayloadThemepersonalizedOffer     PayloadTheme = "PersonalizedOffer"
	PayloadThemequarantine            PayloadTheme = "Quarantine"
	PayloadThemeremoteWork            PayloadTheme = "RemoteWork"
	PayloadThemereviewMessage         PayloadTheme = "ReviewMessage"
	PayloadThemesecurityUpdate        PayloadTheme = "SecurityUpdate"
	PayloadThemeserviceSuspended      PayloadTheme = "ServiceSuspended"
	PayloadThemesignatureRequired     PayloadTheme = "SignatureRequired"
	PayloadThemeunknown               PayloadTheme = "Unknown"
	PayloadThemeupgradeMailboxStorage PayloadTheme = "UpgradeMailboxStorage"
	PayloadThemeverifyMailbox         PayloadTheme = "VerifyMailbox"
	PayloadThemevoicemail             PayloadTheme = "Voicemail"
)

func PossibleValuesForPayloadTheme() []string {
	return []string{
		string(PayloadThemeaccountActivation),
		string(PayloadThemeaccountVerification),
		string(PayloadThemeadvertisement),
		string(PayloadThemebilling),
		string(PayloadThemecleanUpMail),
		string(PayloadThemecontroversial),
		string(PayloadThemedocumentReceived),
		string(PayloadThemeemployeeEngagement),
		string(PayloadThemeexpense),
		string(PayloadThemefax),
		string(PayloadThemefinanceReport),
		string(PayloadThemeincomingMessages),
		string(PayloadThemeinvoice),
		string(PayloadThemeitemReceived),
		string(PayloadThemeloginAlert),
		string(PayloadThememailReceived),
		string(PayloadThemeother),
		string(PayloadThemepassword),
		string(PayloadThemepayment),
		string(PayloadThemepayroll),
		string(PayloadThemepersonalizedOffer),
		string(PayloadThemequarantine),
		string(PayloadThemeremoteWork),
		string(PayloadThemereviewMessage),
		string(PayloadThemesecurityUpdate),
		string(PayloadThemeserviceSuspended),
		string(PayloadThemesignatureRequired),
		string(PayloadThemeunknown),
		string(PayloadThemeupgradeMailboxStorage),
		string(PayloadThemeverifyMailbox),
		string(PayloadThemevoicemail),
	}
}

func parsePayloadTheme(input string) (*PayloadTheme, error) {
	vals := map[string]PayloadTheme{
		"accountactivation":     PayloadThemeaccountActivation,
		"accountverification":   PayloadThemeaccountVerification,
		"advertisement":         PayloadThemeadvertisement,
		"billing":               PayloadThemebilling,
		"cleanupmail":           PayloadThemecleanUpMail,
		"controversial":         PayloadThemecontroversial,
		"documentreceived":      PayloadThemedocumentReceived,
		"employeeengagement":    PayloadThemeemployeeEngagement,
		"expense":               PayloadThemeexpense,
		"fax":                   PayloadThemefax,
		"financereport":         PayloadThemefinanceReport,
		"incomingmessages":      PayloadThemeincomingMessages,
		"invoice":               PayloadThemeinvoice,
		"itemreceived":          PayloadThemeitemReceived,
		"loginalert":            PayloadThemeloginAlert,
		"mailreceived":          PayloadThememailReceived,
		"other":                 PayloadThemeother,
		"password":              PayloadThemepassword,
		"payment":               PayloadThemepayment,
		"payroll":               PayloadThemepayroll,
		"personalizedoffer":     PayloadThemepersonalizedOffer,
		"quarantine":            PayloadThemequarantine,
		"remotework":            PayloadThemeremoteWork,
		"reviewmessage":         PayloadThemereviewMessage,
		"securityupdate":        PayloadThemesecurityUpdate,
		"servicesuspended":      PayloadThemeserviceSuspended,
		"signaturerequired":     PayloadThemesignatureRequired,
		"unknown":               PayloadThemeunknown,
		"upgrademailboxstorage": PayloadThemeupgradeMailboxStorage,
		"verifymailbox":         PayloadThemeverifyMailbox,
		"voicemail":             PayloadThemevoicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadTheme(input)
	return &out, nil
}
