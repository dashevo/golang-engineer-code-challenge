package metric

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
)

// RenderSummaryStats renders the summary statistics and writes to passed io.Writer
func RenderSummaryStats(w io.Writer, stats SummaryStats) error {
	var err error
	_, err = fmt.Fprintf(w, "Your total expenses: %.7f DASH\n\n", stats.TotalCost)
	if err != nil {
		return err
	}
	tw := tabwriter.NewWriter(w, 8, 8, 1, '\t', tabwriter.Debug|tabwriter.AlignRight)
	var (
		requestStats []Stats
		networkStats []Stats
	)
	for _, stats := range stats.GroupedStats.Slice() {
		switch stats.Type {
		case RequestStatsType:
			requestStats = append(requestStats, stats)
		case NetworkStatsType:
			networkStats = append(networkStats, stats)
		}
	}
	if len(requestStats) > 0 {
		_, err := io.WriteString(w, "Summary statistics for all performed requests\n")
		if err != nil {
			return err
		}
		err = writeTable(tw, []string{"Request URL", "Size/bytes", "Elapsed/ms", "Cost/dash"}, requestStats)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, "\n")
	if err != nil {
		return err
	}
	if len(networkStats) > 0 {
		_, err := io.WriteString(w, "Summary statistics for all used networks\n")
		if err != nil {
			return err
		}
		err = writeTable(tw, []string{"Network", "Size/bytes", "Elapsed/ms", "Cost/dash"}, networkStats)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeTable(tw *tabwriter.Writer, header []string, stats []Stats) error {
	_, err := fmt.Fprintf(tw, strings.Join(header, "\t")+"\n")
	if err != nil {
		return err
	}
	for _, s := range stats {
		_, err = fmt.Fprintf(tw, "%s\t%d\t%d\t%f\n", s.Name, s.Size, s.Elapsed, s.Cost)
		if err != nil {
			return err
		}
	}
	return tw.Flush()
}
