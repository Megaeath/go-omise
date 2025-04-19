package summary

import (
	"fmt"
	"sort"

	"go-cli/internal/model"
)

func PrintSummary(results []model.DonationResult) {
	var successSum, failSum, total int
	successMap := map[string]int{}
	var failCount int

	for _, res := range results {
		amt := res.Row.AmountSubunits
		total += amt
		if res.Success {
			successSum += amt
			successMap[res.Row.Name] += amt
		} else {
			failSum += amt
			failCount++
		}
	}

	// sort by donation
	type kv struct {
		Name  string
		Total int
	}
	var top []kv
	for name, amt := range successMap {
		top = append(top, kv{name, amt})
	}
	sort.Slice(top, func(i, j int) bool {
		return top[i].Total > top[j].Total
	})

	fmt.Println()
	fmt.Println("        total received:", formatTHB(total))
	fmt.Println("  successfully donated:", formatTHB(successSum))
	fmt.Println("       faulty donation:", formatTHB(failSum))
	if len(results) > 0 {
		fmt.Println("    average per person:", formatTHB(successSum/int(len(results))))
	}
	fmt.Println("            top donors:")
	for i := 0; i < len(top) && i < 3; i++ {
		fmt.Println("                        " + top[i].Name)
	}
}

func formatTHB(subunits int) string {
	return fmt.Sprintf("THB %10.2f", float64(subunits)/100)
}
