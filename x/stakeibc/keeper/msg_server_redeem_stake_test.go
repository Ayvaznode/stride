package keeper_test

import (
	"fmt"

	epochtypes "github.com/Stride-Labs/stride/x/epochs/types"
	recordtypes "github.com/Stride-Labs/stride/x/records/types"
	stakeibc "github.com/Stride-Labs/stride/x/stakeibc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/stretchr/testify/suite"
)

type RedeemStakeTestCase struct {
	user         Account
	module       Account
	initialState State
	validMsg     stakeibc.MsgRedeemStake
}

func (suite *KeeperTestSuite) SetupRedeemStake() RedeemStakeTestCase {
	redeemAmount := int64(1_000_000)
	initialDepositAmount := int64(1_000_000)
	user := Account{
		acc:           suite.TestAccs[0],
		atomBalance:   sdk.NewInt64Coin("ibc/uatom", 10_000_000),
		stAtomBalance: sdk.NewInt64Coin("stuatom", 10_000_000),
	}
	suite.FundAccount(user.acc, user.atomBalance)
	suite.FundAccount(user.acc, user.stAtomBalance)

	module := Account{
		acc:           suite.App.AccountKeeper.GetModuleAddress(stakeibc.ModuleName),
		atomBalance:   sdk.NewInt64Coin("ibc/uatom", 10_000_000),
		stAtomBalance: sdk.NewInt64Coin("stuatom", 10_000_000),
	}
	suite.FundModuleAccount(stakeibc.ModuleName, module.atomBalance)
	suite.FundModuleAccount(stakeibc.ModuleName, module.stAtomBalance)

	// TODO define the host zone with stakedBal and validators with staked amounts
	hostZone := stakeibc.HostZone{
		ChainId:        "GAIA",
		HostDenom:      "uatom",
		IBCDenom:       "ibc/uatom",
		Bech32Prefix:   "cosmos",
		RedemptionRate: sdk.NewDec(1.0),
		StakedBal:      1234567890,
	}

	epochTrackerDay := stakeibc.EpochTracker{
		EpochIdentifier: epochtypes.DAY_EPOCH,
		EpochNumber:     1,
	}

	initialDepositRecord := recordtypes.DepositRecord{
		Id:                 1,
		DepositEpochNumber: 1,
		HostZoneId:         "GAIA",
		Amount:             initialDepositAmount,
	}

	epochUnbondingRecord := recordtypes.EpochUnbondingRecord{
		Id:                   1,
		UnbondingEpochNumber: 1,
		HostZoneUnbondings:   make(map[string]*recordtypes.HostZoneUnbonding),
	}

	epochUnbondingRecord.HostZoneUnbondings["GAIA"] = &recordtypes.HostZoneUnbonding{
		Amount:     uint64(0),
		Denom:      "uatom",
		HostZoneId: "GAIA",
		Status:     recordtypes.HostZoneUnbonding_BONDED,
	}

	suite.App.StakeibcKeeper.SetHostZone(suite.Ctx, hostZone)
	suite.App.StakeibcKeeper.SetEpochTracker(suite.Ctx, epochTrackerDay)
	suite.App.RecordsKeeper.SetDepositRecord(suite.Ctx, initialDepositRecord)
	suite.App.RecordsKeeper.SetEpochUnbondingRecord(suite.Ctx, epochUnbondingRecord)
	// TODO  why do we need to set this to 2 instead of 1? revisit
	suite.App.RecordsKeeper.SetEpochUnbondingRecordCount(suite.Ctx, 2)

	return RedeemStakeTestCase{
		user:   user,
		module: module,
		initialState: State{
			depositRecordAmount: initialDepositAmount,
			hostZone:            hostZone,
		},
		validMsg: stakeibc.MsgRedeemStake{
			Creator:  user.acc.String(),
			Amount:   redeemAmount,
			HostZone: "GAIA",
			// TODO set this dynamically through test helpers for host zone
			Receiver: "cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8uf",
		},
	}
}

func (suite *KeeperTestSuite) TestRedeemStakeSuccessful() {
	tc := suite.SetupRedeemStake()
	msg := tc.validMsg
	user := tc.user
	redeemAmount := sdk.NewInt(msg.Amount)

	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &msg)
	suite.Require().NoError(err)

	// User STUATOM balance should have DECREASED by the amount to be redeemed
	expectedUserStAtomBalance := user.stAtomBalance.SubAmount(redeemAmount)
	actualUserStAtomBalance := suite.App.BankKeeper.GetBalance(suite.Ctx, user.acc, "stuatom")
	suite.CompareCoins(expectedUserStAtomBalance, actualUserStAtomBalance, "user stuatom balance")

	// Gaia's hostZoneUnbonding amount should have INCREASED from 0 to be amount redeemed multiplied by the redemption rate
	// TODO how can we check the INCREASE rather than the absolute amount here?
	epochUnbondingRecord, found := suite.App.RecordsKeeper.GetLatestEpochUnbondingRecord(suite.Ctx)
	suite.Require().True(found)
	hostZoneUnbondingGaia, found := epochUnbondingRecord.HostZoneUnbondings["GAIA"]
	suite.Require().True(found)
	actualHostZoneUnbondingGaiaAmount := int64(hostZoneUnbondingGaia.Amount)
	hostZone, _ := suite.App.StakeibcKeeper.GetHostZone(suite.Ctx, msg.HostZone)
	expectedHostZoneUnbondingGaiaAmount := redeemAmount.Int64() * hostZone.RedemptionRate.TruncateInt().Int64()
	suite.Require().Equal(expectedHostZoneUnbondingGaiaAmount, actualHostZoneUnbondingGaiaAmount, "host zone unbonding amount")

	// UserRedemptionRecord should have been created with correct amount, sender, receiver, host zone, isClaimable
	userRedemptionRecords := hostZoneUnbondingGaia.UserRedemptionRecords
	suite.Require().Equal(len(userRedemptionRecords), 1)
	userRedemptionRecordId := userRedemptionRecords[0]
	userRedemptionRecord, found := suite.App.RecordsKeeper.GetUserRedemptionRecord(suite.Ctx, userRedemptionRecordId)
	suite.Require().True(found)
	// check amount
	suite.Require().Equal(int64(userRedemptionRecord.Amount), expectedHostZoneUnbondingGaiaAmount)
	// check sender
	suite.Require().Equal(userRedemptionRecord.Sender, msg.Creator)
	// check receiver
	suite.Require().Equal(userRedemptionRecord.Receiver, msg.Receiver)
	// check host zone
	suite.Require().Equal(userRedemptionRecord.HostZoneId, msg.HostZone)
	// check isClaimable
	suite.Require().Equal(userRedemptionRecord.IsClaimable, false)

	// make sure stTokens that were transfered to the module account were burned (stAsset supply should decrease by redeemAmount)
	expectedStAssetSupply := tc.module.stAtomBalance.Amount.Int64() + tc.user.stAtomBalance.Amount.Int64() - redeemAmount.Int64()
	actualStAssetSupply := suite.App.BankKeeper.GetSupply(suite.Ctx, "stuatom")
	suite.Require().Equal(expectedStAssetSupply, actualStAssetSupply.Amount.Int64())
}

func (suite *KeeperTestSuite) TestInvalidCreatorAddress() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg

	// cosmos instead of stride address
	invalidMsg.Creator = "cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8uf"
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, fmt.Sprintf("creator address is invalid: %s: invalid address", invalidMsg.Creator))

	// invalid stride address
	invalidMsg.Creator = "stride1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8uf"
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, fmt.Sprintf("creator address is invalid: %s: invalid address", invalidMsg.Creator))

	// empty address
	invalidMsg.Creator = ""
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, fmt.Sprintf("creator address is invalid: %s: invalid address", invalidMsg.Creator))

	// wrong len address
	invalidMsg.Creator = "stride1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8ufabc"
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, fmt.Sprintf("creator address is invalid: %s: invalid address", invalidMsg.Creator))
}

func (suite *KeeperTestSuite) TestRedeemStakeHostZoneNotFound() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	invalidMsg.HostZone = "fake_host_zone"
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, "host zone is invalid: fake_host_zone: host zone not registered")
}

func (suite *KeeperTestSuite) TestInvalidReceiverAddress() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg

	// stride instead of cosmos address
	invalidMsg.Receiver = "stride159atdlc3ksl50g0659w5tq42wwer334ajl7xnq"
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, "invalid receiver address (invalid Bech32 prefix; expected cosmos, got stride): invalid address")

	// invalid cosmos address
	invalidMsg.Receiver = "cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8ua"
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, "invalid receiver address (decoding bech32 failed: invalid checksum (expected yxp8uf got yxp8ua)): invalid address")

	// empty address
	invalidMsg.Receiver = ""
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, "invalid receiver address (empty address string is not allowed): invalid address")

	// wrong len address
	invalidMsg.Receiver = "cosmos1g6qdx6kdhpf000afvvpte7hp0vnpzapuyxp8ufa"
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	suite.Require().EqualError(err, "invalid receiver address (decoding bech32 failed: invalid checksum (expected xp8ugp got xp8ufa)): invalid address")
}

func (suite *KeeperTestSuite) TestRedeemStakeRedeemMoreThanStaked() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	invalidMsg.Amount = int64(1_000_000_000_000_000)
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, fmt.Sprintf("cannot unstake an amount g.t. staked balance on host zone: %d: invalid amount", invalidMsg.Amount))
}

func (suite *KeeperTestSuite) TestRedeemStakeRedeemUnableToParseCoin() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	invalidMsg.Amount = int64(-1_000_000)
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, fmt.Sprintf("could not parse inCoin: %d%v: invalid coins", invalidMsg.Amount, "stuatom"))
}

func (suite *KeeperTestSuite) TestRedeemStakeRedeemNoEpochTrackerDay() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	suite.App.RecordsKeeper.SetEpochUnbondingRecord(suite.Ctx, recordtypes.EpochUnbondingRecord{})
	suite.App.RecordsKeeper.SetEpochUnbondingRecordCount(suite.Ctx, 0)
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, fmt.Sprintf("latest epoch unbonding record not found: epoch unbonding record not found"))
}

func (suite *KeeperTestSuite) TestRedeemStakeRedeemUserAlreadyRedeemedThisEpoch() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)
	_, err = suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, fmt.Sprintf("user already redeemed this epoch: GAIA.1.%s: redemption record already exists", suite.TestAccs[0]))
}

func (suite *KeeperTestSuite) TestRedeemStakeRedeemHostZoneNotUnbondings() {
	tc := suite.SetupRedeemStake()
	// Update message with invalid denom
	invalidMsg := tc.validMsg
	epochUnbondingRecord := recordtypes.EpochUnbondingRecord{
		Id:                   1,
		UnbondingEpochNumber: 1,
		HostZoneUnbondings:   make(map[string]*recordtypes.HostZoneUnbonding),
	}
	epochUnbondingRecord.HostZoneUnbondings["NOT_GAIA"] = &recordtypes.HostZoneUnbonding{
		Amount:     uint64(0),
		Denom:      "uatom",
		HostZoneId: "GAIA",
		Status:     recordtypes.HostZoneUnbonding_BONDED,
	}
	suite.App.RecordsKeeper.SetEpochUnbondingRecord(suite.Ctx, epochUnbondingRecord)
	_, err := suite.msgServer.RedeemStake(sdk.WrapSDKContext(suite.Ctx), &invalidMsg)

	suite.Require().EqualError(err, "host zone not found in unbondings: GAIA: host zone not registered")
}