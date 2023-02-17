package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration008ProposerCommit = &migrate.Migration{
	Id: "008-proposer-commit",
	Up: []string{`
		ALTER TABLE ` + vars.TableDeliveredPayload + ` ADD validated_at timestamp;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
