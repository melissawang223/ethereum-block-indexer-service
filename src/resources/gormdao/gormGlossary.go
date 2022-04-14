package gormdao

// it's a grossary
const (
	StatusUnrevealed = "unrevealed"

	StatusEnabled    = "enabled"
	StatusSuppressed = "suppressed"
	StatusFreezed    = "freezed"
	StatusDisabled   = "disabled"
	StatusDiscarded  = "discarded"

	RoleCompany         = "company"
	RoleAgent           = "agent"
	RoleMember          = "member"
	RoleGod             = "auroratech_god"
	RoleCustomerService = "auroratech_customerService"
	RoleFinancial       = "auroratech_financial"

	SearchTypeAccount = "account"
	SearchTypeName    = "name"

	TradeTypeDeposit  = "deposit"
	TradeTypeWithdraw = "withdraw"

	WalletActionManualDeposit  = "manualDeposit"  // 上分
	WalletActionManualWithdraw = "manualWithdraw" // 減分

	TinyintMax  = 255
	SmallIntMax = 65535
	IntMax      = 4294967295
	BigIntMax   = 18446744073709551615

	//company Name
	AuroraTech = "auroratech"

	//deposit order
	Unsettled = "unsettled"
	Settled   = "settled"
	Cancelled = "cancelled"
	Failed    = "failed"

	//Announcement
	AnnouncementTypeFrontEnd = "frontend"
	AnnouncementTypeBackEnd  = "backend"
)
