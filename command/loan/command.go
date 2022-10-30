package loan

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evercyan/brick/xcli/xcolor"
	"github.com/evercyan/brick/xcli/xtable"
	"github.com/evercyan/mmmoney/config"
	"github.com/evercyan/mmmoney/internal"
	"github.com/spf13/cobra"
)

var (
	// Command ...
	Command = &cobra.Command{
		Use:     "loan",
		Aliases: []string{"l"},
		Short:   "计算房贷本金利息",
		Run: func(cmd *cobra.Command, args []string) {
			ans := &Answer{}
			err := survey.Ask(Question, ans)
			if err != nil {
				xcolor.Fail(config.Error, fmt.Sprintf("输入错误: %s", err.Error()))
				return
			}

			// 输出贷款基本信息
			items := [][]interface{}{
				{
					"金额", ans.Capital,
				},
				{
					"利率", fmt.Sprintf("%s %.2f%%", ans.RateType, ans.RateValue*100),
				},
				{
					"周期", fmt.Sprintf("%d %s", ans.PeriodValue, ans.PeriodType),
				},
				{
					"方式", ans.PayType,
				},
			}
			xcolor.Success(config.Separator)
			xtable.New(items).Style(xtable.Dashed).Border(true).Render()
			xcolor.Success(config.Separator)

			// 计算不同偿还方式下的还款信息
			switch ans.PayType {
			case PayTypeOnce:
				rate := ans.GetRateByPeroid()
				fee := ans.Capital * rate * float64(ans.PeriodValue)
				xcolor.Success(config.Success, fmt.Sprintf("利息: %s", internal.FormatPrice(fee)))
			case PayTypeMonthEqualInterest:
				rate, peroid := ans.GetMonthValue()
				monthList := CalMonthEqualInterest(ans.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本息")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			case PayTypeMonthEqualCapital:
				rate, peroid := ans.GetMonthValue()
				monthList := CalMonthEqualCapital(ans.Capital, rate, peroid)
				xcolor.Success(config.Success, "等额本金")
				xtable.New(monthList).Style(xtable.Dashed).Border(true).Render()
			}

			xcolor.Success(config.Separator)
		},
	}
)
