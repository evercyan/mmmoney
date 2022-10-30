package loan

import (
	"math"

	"github.com/AlecAivazis/survey/v2"
	"github.com/evercyan/mmmoney/internal"
)

var (
	Question = []*survey.Question{
		{
			Name: "Capital",
			Prompt: &survey.Input{
				Message: "请输入贷款本金?",
				Default: "880000",
			},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "RateType",
			Prompt: &survey.Select{
				Message: "请选择利率类型:",
				Options: []string{
					RateTypeDay,
					RateTypeMonth,
					RateTypeYear,
				},
				Default: RateTypeYear,
			},
		},
		{
			Name: "RateValue",
			Prompt: &survey.Input{
				Message: "请输入利率?",
				Default: "0.0588",
			},
		},
		{
			Name: "PeriodType",
			Prompt: &survey.Select{
				Message: "请选择周期类型:",
				Options: []string{
					PeriodTypeDay,
					PeriodTypeMonth,
					PeriodTypeYear,
				},
				Default: PeriodTypeYear,
			},
		},
		{
			Name: "PeriodValue",
			Prompt: &survey.Input{
				Message: "请输入周期期数?",
				Default: "30",
			},
		},
		{
			Name: "PayType",
			Prompt: &survey.Select{
				Message: "请选择偿还方式:",
				Options: []string{
					PayTypeOnce,
					PayTypeMonthEqualInterest,
					PayTypeMonthEqualCapital,
				},
				Default: PayTypeMonthEqualInterest,
			},
		},
	}
)

// ----------------------------------------------------------------

// MonthPayment 月还款信息
type MonthPayment struct {
	PeroidNum     int    `table:"期数"`
	MonthTotal    string `table:"月还金额"`
	MonthCapital  string `table:"月还本金"`
	MonthInterest string `table:"月还利息"`
	Total         string `table:"已还金额"`
	TotalCapital  string `table:"已还本金"`
	TotalInterest string `table:"已还利息"`
	RemainCapital string `table:"剩余本金"`
}

// ----------------------------------------------------------------

// Answer ...
type Answer struct {
	Capital     float64 `survey:"Capital"`
	RateType    string  `survey:"RateType"`
	RateValue   float64 `survey:"RateValue"`
	PeriodType  string  `survey:"PeriodType"`
	PeriodValue int     `survey:"PeriodValue"`
	PayType     string  `survey:"PayType"`
}

// GetRateByPeroid 根据还款周期计算利率
func (t *Answer) GetRateByPeroid() float64 {
	if (t.PeriodType == PeriodTypeDay && t.RateType == RateTypeDay) ||
		(t.PeriodType == PeriodTypeMonth && t.RateType == RateTypeMonth) ||
		(t.PeriodType == PeriodTypeYear && t.RateType == RateTypeYear) {
		return t.RateValue
	}
	switch t.PeriodType {
	case PeriodTypeDay:
		if t.RateType == RateTypeMonth {
			return t.RateValue / 30
		}
		return t.RateValue / 365
	case PeriodTypeMonth:
		if t.RateType == RateTypeDay {
			return t.RateValue * 30
		}
		return t.RateValue / 12
	case PeriodTypeYear:
		if t.RateType == RateTypeDay {
			return t.RateValue * 365
		}
		return t.RateValue * 12
	}
	return 0
}

// GetMonthValue 获取按月还款参数
func (t *Answer) GetMonthValue() (float64, int) {
	var (
		rate   float64 = 0
		peroid         = 0
	)
	switch t.PeriodType {
	case PeriodTypeDay:
		peroid = t.PeriodValue / 30
	case PeriodTypeMonth:
		peroid = t.PeriodValue
	case PeriodTypeYear:
		peroid = t.PeriodValue * 12
	}
	switch t.RateType {
	case RateTypeDay:
		rate = t.RateValue * 30
	case RateTypeMonth:
		rate = t.RateValue
	case RateTypeYear:
		rate = t.RateValue / 12
	}
	return rate, peroid
}

// ----------------------------------------------------------------

// CalMonthEqualInterest 计算等额本息(总额, 月利率, 周期(月))
func CalMonthEqualInterest(
	amount float64,
	rate float64,
	peroid int,
) []*MonthPayment {
	list := make([]*MonthPayment, 0)
	monthTotal := amount * rate * math.Pow(1+rate, float64(peroid)) / (math.Pow(1+rate, float64(peroid)) - 1)
	remainCapital := amount
	total := float64(0)
	for i := 0; i < peroid; i++ {
		monthInterest := remainCapital * rate
		monthCapital := monthTotal - monthInterest
		total = total + monthTotal
		remainCapital = remainCapital - monthCapital
		if remainCapital <= 0 {
			remainCapital = 0
		}
		totalCapital := amount - remainCapital
		totalInterest := total - totalCapital
		list = append(list, &MonthPayment{
			PeroidNum:     i + 1,
			MonthTotal:    internal.FormatPrice(monthTotal),
			MonthCapital:  internal.FormatPrice(monthCapital),
			MonthInterest: internal.FormatPrice(monthInterest),
			Total:         internal.FormatPrice(total),
			TotalCapital:  internal.FormatPrice(totalCapital),
			TotalInterest: internal.FormatPrice(totalInterest),
			RemainCapital: internal.FormatPrice(remainCapital),
		})
	}
	return list
}

// CalMonthEqualCapital 计算等额本金(总额, 月利率, 周期(月))
func CalMonthEqualCapital(
	capital float64,
	rate float64,
	peroid int,
) []*MonthPayment {
	list := make([]*MonthPayment, 0)
	monthCapital := capital / float64(peroid)
	remainCapital := capital
	total := float64(0)
	for i := 0; i < peroid; i++ {
		monthInterest := remainCapital * rate
		monthTotal := monthCapital + monthInterest
		remainCapital = remainCapital - monthCapital
		total = total + monthTotal
		totalCapital := capital - remainCapital
		totalInterest := total - totalCapital
		list = append(list, &MonthPayment{
			PeroidNum:     i + 1,
			MonthTotal:    internal.FormatPrice(monthTotal),
			MonthCapital:  internal.FormatPrice(monthCapital),
			MonthInterest: internal.FormatPrice(monthInterest),
			Total:         internal.FormatPrice(total),
			TotalCapital:  internal.FormatPrice(totalCapital),
			TotalInterest: internal.FormatPrice(totalInterest),
			RemainCapital: internal.FormatPrice(remainCapital),
		})
	}
	return list
}
