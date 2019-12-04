//nolint:deadcode,unused
package rest

import "github.com/cosmos/cosmos-sdk/x/mint/internal/types"

// Concrete Swagger types used to generate REST documentation. Note, these types
// are not actually used but since all queries return a generic JSON raw message,
// they enabled typed documentation.

type (
	// alias the params type for use in swagger auto gendoc
	MintParams types.Params
	mintParams struct {
		Height int64      `json:"height"`
		Result MintParams `json:"result"`
	}

	mintInflation struct {
		Height int64  `json:"height"`
		Result string `json:"result"`
	}
)