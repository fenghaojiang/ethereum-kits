package account

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

func ValidateAccount(account *Account, address common.Address) bool {
	return strings.EqualFold(account.AccountAddress.String(), address.String())
}
