package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration006ProfilingExt = &migrate.Migration{
	Id: "006-profiling-ext",
	Up: []string{`
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD decode_duration        bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD cache_read_duration    bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD randao_lock_1_duration bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD duties_lock_duration   bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD checks_duration        bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD randao_lock_2_duration bigint NOT NULL default 0;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
