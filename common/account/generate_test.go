package account_test

import (
	"strings"
	"testing"

	"github.com/fenghaojiang/ethereum-kits/common/account"
	"github.com/fenghaojiang/ethereum-kits/common/log"
	"go.uber.org/zap"
)

func TestOnGenerateNewAccount(t *testing.T) {
	privateKey, _account, err := account.GenerateAccount()
	for err == nil && !strings.HasPrefix(_account, "0x0000") {
		privateKey, _account, err = account.GenerateAccount()
	}

	log.Info("create account success", zap.String("private key", privateKey), zap.String("account", _account))
}
