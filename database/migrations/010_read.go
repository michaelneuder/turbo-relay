package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration010Read = &migrate.Migration{
	Id: "010-read",
	Up: []string{`
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD read_header_duration bigint NOT NULL default 0;
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD read_duration bigint NOT NULL default 0;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
