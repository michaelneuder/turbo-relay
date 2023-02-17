package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration009DemotionRefactor = &migrate.Migration{
	Id: "008-demotion-refactor",
	Up: []string{`
		ALTER TABLE ` + vars.TableBuilderDemotions + ` DROP COLUMN get_payload_sim_error;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
