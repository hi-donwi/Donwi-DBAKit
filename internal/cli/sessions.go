package cli

import (
	"context"

	"github.com/hi-donwi/Donwi-DBAKit/internal/collector"
	"github.com/hi-donwi/Donwi-DBAKit/internal/evaluator"
	"github.com/hi-donwi/Donwi-DBAKit/internal/model"
	"github.com/hi-donwi/Donwi-DBAKit/internal/report"
	"github.com/spf13/cobra"
)

var sessionsCmd = &cobra.Command{
	Use:     "sessions",
	Aliases: []string{"activity"},
	Short:   "List active PostgreSQL sessions",
	Long: `Lists sessions from pg_stat_activity (pid, user, database, state,
duration, wait event, and truncated query text) longest-running first. Read-only.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runReport(cmd, "sessions", sessionsRun)
	},
}

func sessionsRun(ctx context.Context, _ model.Connectivity, coll *collector.Collector, _ evaluator.Thresholds) (any, []model.Finding, error) {
	sessions, err := coll.Sessions(ctx, flagMaxQueryLen)
	if err != nil {
		return nil, nil, err
	}
	return report.SessionsData{Sessions: sessions}, nil, nil
}
