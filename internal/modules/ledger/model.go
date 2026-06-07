package ledger

import "time"

// TransactionType constants
const (
	TxTypeImport  = "IMPORT"
	TxTypeSellMat = "SELL_MAT"
	TxTypeBuyProd = "BUY_PROD"
	TxTypeExport  = "EXPORT"
)

// Transaction records physical/virtual movement of items and related debt.
type Transaction struct {
	ID          string    `json:"id" db:"id"`
	Type        string    `json:"type" db:"type"`
	PartnerID   string    `json:"partner_id" db:"partner_id"`
	ItemID      string    `json:"item_id" db:"item_id"`
	Quantity    float64   `json:"quantity" db:"quantity"`
	UnitPrice   int64     `json:"unit_price" db:"unit_price"`
	TotalAmount int64     `json:"total_amount" db:"total_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// CashFlowType constants
const (
	CashIn  = "PAY_IN"
	CashOut = "PAY_OUT"
)

// CashFlow records actual money movements.
type CashFlow struct {
	ID        string    `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	PartnerID string    `json:"partner_id" db:"partner_id"`
	Amount    int64     `json:"amount" db:"amount"`
	Note      *string   `json:"note,omitempty" db:"note"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// LedgerEntry is an audit record created for every Transaction and CashFlow.
type LedgerEntry struct {
	ID          string    `json:"id" db:"id"`
	PartnerID   string    `json:"partner_id" db:"partner_id"`
	Amount      int64     `json:"amount" db:"amount"`           // positive or negative
	SourceType  string    `json:"source_type" db:"source_type"` // TRANSACTION or CASHFLOW
	SourceID    string    `json:"source_id" db:"source_id"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// LedgerBalance holds the current balance per partner (one row per partner).
type LedgerBalance struct {
	PartnerID string    `json:"partner_id" db:"partner_id"` // PK + FK to partners
	Balance   int64     `json:"balance" db:"balance"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (Transaction) TableName() string   { return "transactions" }
func (CashFlow) TableName() string      { return "cashflows" }
func (LedgerEntry) TableName() string   { return "ledger_entries" }
func (LedgerBalance) TableName() string { return "ledger_balances" }
