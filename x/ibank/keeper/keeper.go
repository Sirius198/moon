package keeper

import (
	"errors"
	"fmt"

	"moon/x/ibank/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramstore:    ps,
		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateModuleAccount(ctx sdk.Context) {
	macc := authtypes.NewEmptyModuleAccount(types.ModuleName, authtypes.Minter)
	k.accountKeeper.SetModuleAccount(ctx, macc)
}

// This function send amt Coin from `from` account to a module account
func (k Keeper) SendCoin(ctx sdk.Context, from, to sdk.AccAddress, amt sdk.Coin) error {
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, from, types.ModuleName, sdk.Coins{amt})
	if err != nil {
		return err
	}

	k.AppendTransaction(ctx, types.Transaction{
		Sender:     from.String(),
		Receiver:   to.String(),
		Coins:      sdk.Coins{amt},
		SentAt:     ctx.BlockTime(),
		ReceivedAt: ctx.BlockTime(),
		Status:     types.TXN_PENDING,
	})

	return nil
}

func (k Keeper) ReceiveCoin(ctx sdk.Context, receiver sdk.AccAddress, transactionId int64) error {
	txn, found := k.GetTransaction(ctx, uint64(transactionId))
	if !found {
		return errors.New("transaction not found")
	}

	// Check if transaction is performed
	if txn.ReceivedAt != txn.SentAt {
		return errors.New("you already performed this transaction")
	}

	// Check if transaction is expired
	params := k.GetParams(ctx)
	duration := ctx.BlockTime().Sub(txn.SentAt)
	if duration > params.DurationOfExpiration {
		return errors.New("transaction has expired")
	}

	to, err := sdk.AccAddressFromBech32(txn.Receiver)
	if err != nil || !to.Equals(receiver) {
		return errors.New("receiver not found")
	}

	txn.ReceivedAt = ctx.BlockTime()
	txn.Status = types.TXN_SENT
	k.SetTransaction(ctx, txn)

	return k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, txn.Coins)
}
