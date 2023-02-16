//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of customer address provided.
type TaxTransactionCustomerDetailsAddressSource string

// List of values that TaxTransactionCustomerDetailsAddressSource can take
const (
	TaxTransactionCustomerDetailsAddressSourceBilling  TaxTransactionCustomerDetailsAddressSource = "billing"
	TaxTransactionCustomerDetailsAddressSourceShipping TaxTransactionCustomerDetailsAddressSource = "shipping"
)

// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
type TaxTransactionCustomerDetailsTaxIDType string

// List of values that TaxTransactionCustomerDetailsTaxIDType can take
const (
	TaxTransactionCustomerDetailsTaxIDTypeAETRN    TaxTransactionCustomerDetailsTaxIDType = "ae_trn"
	TaxTransactionCustomerDetailsTaxIDTypeAUABN    TaxTransactionCustomerDetailsTaxIDType = "au_abn"
	TaxTransactionCustomerDetailsTaxIDTypeAUARN    TaxTransactionCustomerDetailsTaxIDType = "au_arn"
	TaxTransactionCustomerDetailsTaxIDTypeBGUIC    TaxTransactionCustomerDetailsTaxIDType = "bg_uic"
	TaxTransactionCustomerDetailsTaxIDTypeBRCNPJ   TaxTransactionCustomerDetailsTaxIDType = "br_cnpj"
	TaxTransactionCustomerDetailsTaxIDTypeBRCPF    TaxTransactionCustomerDetailsTaxIDType = "br_cpf"
	TaxTransactionCustomerDetailsTaxIDTypeCABN     TaxTransactionCustomerDetailsTaxIDType = "ca_bn"
	TaxTransactionCustomerDetailsTaxIDTypeCAGSTHST TaxTransactionCustomerDetailsTaxIDType = "ca_gst_hst"
	TaxTransactionCustomerDetailsTaxIDTypeCAPSTBC  TaxTransactionCustomerDetailsTaxIDType = "ca_pst_bc"
	TaxTransactionCustomerDetailsTaxIDTypeCAPSTMB  TaxTransactionCustomerDetailsTaxIDType = "ca_pst_mb"
	TaxTransactionCustomerDetailsTaxIDTypeCAPSTSK  TaxTransactionCustomerDetailsTaxIDType = "ca_pst_sk"
	TaxTransactionCustomerDetailsTaxIDTypeCAQST    TaxTransactionCustomerDetailsTaxIDType = "ca_qst"
	TaxTransactionCustomerDetailsTaxIDTypeCHVAT    TaxTransactionCustomerDetailsTaxIDType = "ch_vat"
	TaxTransactionCustomerDetailsTaxIDTypeCLTIN    TaxTransactionCustomerDetailsTaxIDType = "cl_tin"
	TaxTransactionCustomerDetailsTaxIDTypeEGTIN    TaxTransactionCustomerDetailsTaxIDType = "eg_tin"
	TaxTransactionCustomerDetailsTaxIDTypeESCIF    TaxTransactionCustomerDetailsTaxIDType = "es_cif"
	TaxTransactionCustomerDetailsTaxIDTypeEUOSSVAT TaxTransactionCustomerDetailsTaxIDType = "eu_oss_vat"
	TaxTransactionCustomerDetailsTaxIDTypeEUVAT    TaxTransactionCustomerDetailsTaxIDType = "eu_vat"
	TaxTransactionCustomerDetailsTaxIDTypeGBVAT    TaxTransactionCustomerDetailsTaxIDType = "gb_vat"
	TaxTransactionCustomerDetailsTaxIDTypeGEVAT    TaxTransactionCustomerDetailsTaxIDType = "ge_vat"
	TaxTransactionCustomerDetailsTaxIDTypeHKBR     TaxTransactionCustomerDetailsTaxIDType = "hk_br"
	TaxTransactionCustomerDetailsTaxIDTypeHUTIN    TaxTransactionCustomerDetailsTaxIDType = "hu_tin"
	TaxTransactionCustomerDetailsTaxIDTypeIDNPWP   TaxTransactionCustomerDetailsTaxIDType = "id_npwp"
	TaxTransactionCustomerDetailsTaxIDTypeILVAT    TaxTransactionCustomerDetailsTaxIDType = "il_vat"
	TaxTransactionCustomerDetailsTaxIDTypeINGST    TaxTransactionCustomerDetailsTaxIDType = "in_gst"
	TaxTransactionCustomerDetailsTaxIDTypeISVAT    TaxTransactionCustomerDetailsTaxIDType = "is_vat"
	TaxTransactionCustomerDetailsTaxIDTypeJPCN     TaxTransactionCustomerDetailsTaxIDType = "jp_cn"
	TaxTransactionCustomerDetailsTaxIDTypeJPRN     TaxTransactionCustomerDetailsTaxIDType = "jp_rn"
	TaxTransactionCustomerDetailsTaxIDTypeJPTRN    TaxTransactionCustomerDetailsTaxIDType = "jp_trn"
	TaxTransactionCustomerDetailsTaxIDTypeKEPIN    TaxTransactionCustomerDetailsTaxIDType = "ke_pin"
	TaxTransactionCustomerDetailsTaxIDTypeKRBRN    TaxTransactionCustomerDetailsTaxIDType = "kr_brn"
	TaxTransactionCustomerDetailsTaxIDTypeLIUID    TaxTransactionCustomerDetailsTaxIDType = "li_uid"
	TaxTransactionCustomerDetailsTaxIDTypeMXRFC    TaxTransactionCustomerDetailsTaxIDType = "mx_rfc"
	TaxTransactionCustomerDetailsTaxIDTypeMYFRP    TaxTransactionCustomerDetailsTaxIDType = "my_frp"
	TaxTransactionCustomerDetailsTaxIDTypeMYITN    TaxTransactionCustomerDetailsTaxIDType = "my_itn"
	TaxTransactionCustomerDetailsTaxIDTypeMYSST    TaxTransactionCustomerDetailsTaxIDType = "my_sst"
	TaxTransactionCustomerDetailsTaxIDTypeNOVAT    TaxTransactionCustomerDetailsTaxIDType = "no_vat"
	TaxTransactionCustomerDetailsTaxIDTypeNZGST    TaxTransactionCustomerDetailsTaxIDType = "nz_gst"
	TaxTransactionCustomerDetailsTaxIDTypePHTIN    TaxTransactionCustomerDetailsTaxIDType = "ph_tin"
	TaxTransactionCustomerDetailsTaxIDTypeRUINN    TaxTransactionCustomerDetailsTaxIDType = "ru_inn"
	TaxTransactionCustomerDetailsTaxIDTypeRUKPP    TaxTransactionCustomerDetailsTaxIDType = "ru_kpp"
	TaxTransactionCustomerDetailsTaxIDTypeSAVAT    TaxTransactionCustomerDetailsTaxIDType = "sa_vat"
	TaxTransactionCustomerDetailsTaxIDTypeSGGST    TaxTransactionCustomerDetailsTaxIDType = "sg_gst"
	TaxTransactionCustomerDetailsTaxIDTypeSGUEN    TaxTransactionCustomerDetailsTaxIDType = "sg_uen"
	TaxTransactionCustomerDetailsTaxIDTypeSITIN    TaxTransactionCustomerDetailsTaxIDType = "si_tin"
	TaxTransactionCustomerDetailsTaxIDTypeTHVAT    TaxTransactionCustomerDetailsTaxIDType = "th_vat"
	TaxTransactionCustomerDetailsTaxIDTypeTRTIN    TaxTransactionCustomerDetailsTaxIDType = "tr_tin"
	TaxTransactionCustomerDetailsTaxIDTypeTWVAT    TaxTransactionCustomerDetailsTaxIDType = "tw_vat"
	TaxTransactionCustomerDetailsTaxIDTypeUAVAT    TaxTransactionCustomerDetailsTaxIDType = "ua_vat"
	TaxTransactionCustomerDetailsTaxIDTypeUnknown  TaxTransactionCustomerDetailsTaxIDType = "unknown"
	TaxTransactionCustomerDetailsTaxIDTypeUSEIN    TaxTransactionCustomerDetailsTaxIDType = "us_ein"
	TaxTransactionCustomerDetailsTaxIDTypeZAVAT    TaxTransactionCustomerDetailsTaxIDType = "za_vat"
)

// If `reversal`, this transaction reverses an earlier transaction.
type TaxTransactionType string

// List of values that TaxTransactionType can take
const (
	TaxTransactionTypeReversal    TaxTransactionType = "reversal"
	TaxTransactionTypeTransaction TaxTransactionType = "transaction"
)

// Retrieves a Tax Transaction object.
type TaxTransactionParams struct {
	Params `form:"*"`
	// Tax Calculation ID to be used as input when creating the transaction.
	FromCalculation *string `form:"from_calculation"`
	// A custom order or sale identifier, such as 'myOrder_123'.
	Reference *string `form:"reference"`
}

// Lists Tax Transaction objects.
type TaxTransactionListTransactionsParams struct {
	ListParams `form:"*"`
}

// The line item amounts to reverse.
type TaxTransactionCreateReversalLineItemParams struct {
	// The amount to reverse, in negative integer cents.
	Amount *int64 `form:"amount"`
	// The amount of tax to reverse, in negative integer cents.
	AmountTax *int64 `form:"amount_tax"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata"`
	// The `id` of the line item to reverse in the original transaction.
	OriginalLineItem *string `form:"original_line_item"`
	// The quantity reversed.
	Quantity *int64 `form:"quantity"`
	// A custom identifier for this line item in the reversal transaction, such as 'L1-refund'.
	Reference *string `form:"reference"`
}

// Partially or fully reverses a previously created Transaction.
type TaxTransactionCreateReversalParams struct {
	Params `form:"*"`
	// The line item amounts to reverse.
	LineItems []*TaxTransactionCreateReversalLineItemParams `form:"line_items"`
	// If `partial`, the provided line item amounts are reversed. If `full`, the original transaction is fully reversed.
	Mode *string `form:"mode"`
	// The ID of the Transaction to partially or fully reverse.
	OriginalTransaction *string `form:"original_transaction"`
	// A custom identifier for this reversal, such as 'myOrder_123-refund_1'. Must be unique across all transactions.
	Reference *string `form:"reference"`
}

// Retrieves the line items of a committed standalone transaction as a collection.
type TaxTransactionListLineItemsParams struct {
	ListParams  `form:"*"`
	Transaction *string `form:"-"` // Included in URL
}

// The customer's tax IDs (e.g., EU VAT numbers).
type TaxTransactionCustomerDetailsTaxID struct {
	// The type of the tax ID, one of `eu_vat`, `br_cnpj`, `br_cpf`, `eu_oss_vat`, `gb_vat`, `nz_gst`, `au_abn`, `au_arn`, `in_gst`, `no_vat`, `za_vat`, `ch_vat`, `mx_rfc`, `sg_uen`, `ru_inn`, `ru_kpp`, `ca_bn`, `hk_br`, `es_cif`, `tw_vat`, `th_vat`, `jp_cn`, `jp_rn`, `jp_trn`, `li_uid`, `my_itn`, `us_ein`, `kr_brn`, `ca_qst`, `ca_gst_hst`, `ca_pst_bc`, `ca_pst_mb`, `ca_pst_sk`, `my_sst`, `sg_gst`, `ae_trn`, `cl_tin`, `sa_vat`, `id_npwp`, `my_frp`, `il_vat`, `ge_vat`, `ua_vat`, `is_vat`, `bg_uic`, `hu_tin`, `si_tin`, `ke_pin`, `tr_tin`, `eg_tin`, `ph_tin`, or `unknown`
	Type TaxTransactionCustomerDetailsTaxIDType `json:"type"`
	// The value of the tax ID.
	Value string `json:"value"`
}
type TaxTransactionCustomerDetails struct {
	// The customer's postal address (e.g., home or business location).
	Address *Address `json:"address"`
	// The type of customer address provided.
	AddressSource TaxTransactionCustomerDetailsAddressSource `json:"address_source"`
	// The customer's IP address (IPv4 or IPv6).
	IPAddress string `json:"ip_address"`
	// The customer's tax IDs (e.g., EU VAT numbers).
	TaxIDs []*TaxTransactionCustomerDetailsTaxID `json:"tax_ids"`
}

// If `type=reversal`, contains information about what was reversed.
type TaxTransactionReversal struct {
	// The `id` of the reversed `Transaction` object.
	OriginalTransaction string `json:"original_transaction"`
}

// A Tax `Transaction` records the tax collected from or refunded to your customer.
type TaxTransaction struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of an existing [Customer](https://stripe.com/docs/api/customers/object) used for the resource.
	Customer        string                         `json:"customer"`
	CustomerDetails *TaxTransactionCustomerDetails `json:"customer_details"`
	// Unique identifier for the transaction.
	ID string `json:"id"`
	// The tax collected or refunded, by line item.
	LineItems *LineItemList `json:"line_items"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// A custom unique identifier, such as 'myOrder_123'.
	Reference string `json:"reference"`
	// If `type=reversal`, contains information about what was reversed.
	Reversal *TaxTransactionReversal `json:"reversal"`
	// Timestamp of date at which the tax rules and rates in effect applies for the calculation.
	TaxDate int64 `json:"tax_date"`
	// If `reversal`, this transaction reverses an earlier transaction.
	Type TaxTransactionType `json:"type"`
}

// TaxTransactionList is a list of Transactions as retrieved from a list endpoint.
type TaxTransactionList struct {
	APIResource
	ListMeta
	Data []*TaxTransaction `json:"data"`
}
