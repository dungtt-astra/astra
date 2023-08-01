package app

import (
	v1 "github.com/AstraProtocol/astra/v3/app/upgrades/v1"
	v3 "github.com/AstraProtocol/astra/v3/app/upgrades/v3"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// BeginBlockForks executes any necessary fork logic based upon the current block height.
func BeginBlockForks(ctx sdk.Context, app *Astra) {
	switch ctx.BlockHeight() {
	case v1.UpgradeHeight:
		// NOTE: only run for testnet
		if strings.HasPrefix(ctx.ChainID(), MainnetChainID) {
			return
		}

		upgradePlan := upgradetypes.Plan{
			Name:   v1.UpgradeName,
			Info:   v1.UpgradeInfo,
			Height: v1.UpgradeHeight,
		}

		err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan)
		if err != nil {
			panic(err)
		}
	case v3.MainnetUpgradeHeight:
		upgradePlan := upgradetypes.Plan{
			Name:   v3.UpgradeName,
			Info:   v3.UpgradeInfo,
			Height: v3.MainnetUpgradeHeight,
		}

		err := app.UpgradeKeeper.ScheduleUpgrade(ctx, upgradePlan)
		if err != nil {
			panic(err)
		}
	default:
		// do nothing
		return
	}
}
