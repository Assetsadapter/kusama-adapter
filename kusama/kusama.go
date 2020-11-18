package kusama

import (
	"github.com/Assetsadapter/polkadot-adapter/polkadot"
	"github.com/blocktree/openwallet/v2/log"
)

const (
	Symbol      = "KSM"
	MasterKey   = "Kusama seed"
	GenesisHash = "b0a8d493285c2df73290dfb7e61f870f17b41801197a149ca93654499ea3dafe"
	SpecVersion = 1058
	AddrPrefix  = 0x02
)

type WalletManager struct {
	*polkadot.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = polkadot.NewWalletManager()
	wm.Config = polkadot.NewConfig(Symbol, MasterKey, GenesisHash, SpecVersion, AddrPrefix)
	//wm.LoadAssetsConfig( wm.Config )
	wm.Log = log.NewOWLogger(wm.Symbol())
	return &wm
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "Kusama"
}
