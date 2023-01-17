package keeper_test

import (
	"time"

	"github.com/tessornetwork/fury/x/epochs/types"
)

func (suite *KeeperTestSuite) TestEpochLifeCycle() {
	suite.SetupTest()

	epochInfo := types.EpochInfo{
		Identifier:            "month",
		StartTime:             time.Time{},
		Duration:              time.Hour * 24 * 30,
		CurrentEpoch:          0,
		CurrentEpochStartTime: time.Time{},
		EpochCountingStarted:  false,
		CurrentEpochEnded:     true,
	}
	suite.app.EpochsKeeper.SetEpochInfo(suite.ctx, epochInfo)
	epochInfoSaved := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, "month")
	suite.Require().Equal(epochInfo, epochInfoSaved)

	allEpochs := suite.app.EpochsKeeper.AllEpochInfos(suite.ctx)
	suite.Require().Len(allEpochs, 5)
	suite.Require().Equal(allEpochs[0].Identifier, "day") // alphabetical order
	suite.Require().Equal(allEpochs[1].Identifier, "hour")
	suite.Require().Equal(allEpochs[2].Identifier, "minute")
	suite.Require().Equal(allEpochs[3].Identifier, "month")
	suite.Require().Equal(allEpochs[4].Identifier, "week")

	suite.app.EpochsKeeper.DeleteEpochInfo(suite.ctx, "month")
	epochInfomonth := suite.app.EpochsKeeper.GetEpochInfo(suite.ctx, "month")
	suite.Require().Equal(types.EpochInfo{}, epochInfomonth)
}
