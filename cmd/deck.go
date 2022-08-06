package main

import (
	cbn "github.com/DesistDaydream/yu-gi-oh/pkg/combination"
	"github.com/DesistDaydream/yu-gi-oh/pkg/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var (
	z    = "黑魔导+黑魔导女孩+马哈德"
	c    = "黑之魔导阵"
	e    = "永远之魂"
	s    = "魔术师双魂"
	Deck = []string{
		z, z, z, z, z, "黑魔导", "黑魔导", "黑魔导",
		"b", "b", "b", "b", "b", "b", "b", "b", "b",
		c, c, c, e, e, s, s, s,
		"TheEyeOfTimaeus", "TheEyeOfTimaeus",
		"ApprenticeIllusionMagician", "ApprenticeIllusionMagician",
		"e", "f", "g", "h", "i", "g", "k", "l", "m", "n", "o",
	}
	Hand = []string{z}
	// Hand = []string{"DarkMagicalCircle", "EternalSoul", "MagiciansSouls", "TheEyeOfTimaeus", "ApprenticeIllusionMagician"}
)

type DarkMagician struct {
	DarkMagician               string `json:"黑魔导"`
	DarkMagicalCircle          string `json:"黑之魔导阵"`
	EternalSoul                string `json:"永远之魂"`
	MagiciansSouls             string `json:"魔术师双魂"`
	TheEyeOfTimaeus            string `json:"蒂迈欧之眼"`
	ApprenticeIllusionMagician string `json:"幻想见习魔导师"`

	// a string `json:"黑魔术师+幻想之见习魔导师+魔术师双魂"`
	// b string `json:"魔术师的救出+永远之魂+魔术师的导门阵+黑魔术之杖"`
}

func main() {
	// 设置命令行标志
	ygoFlags := &log.YGOLogFlags{}
	ygoFlags.AddYuqueExportFlags()
	pflag.Parse()

	// 初始化日志
	if err := log.LogInit(ygoFlags.LogLevel, ygoFlags.LogFile, ygoFlags.LogFormat); err != nil {
		logrus.Fatal("初始化日志失败", err)
	}

	// 设置变量
	var (
		DeckCount         int = len(Deck) // 牌组总数
		HandCount         int = 5         // 起始手牌数
		TargetCombination int = 0         // 满足条件的手牌组合数
	)

	logrus.Debugf("牌组总数: %v", DeckCount)

	// 遍历牌组，获取牌组中所有组合种类的列表
	combinations := cbn.TraversalDeckCombination(Deck, cbn.CombinationIndexs(DeckCount, int(HandCount)))

	logrus.Debugf("原始组合总数: %v", len(combinations))
	// fmt.Println("牌组中所有组合列表:", combinations)
	cbn.CheckResult(DeckCount, HandCount, combinations)

	// 获取牌组中指定组合的总数
	for _, combination := range combinations {
		if cbn.ConditionCount(combination, Hand) {
			TargetCombination++
		}
	}

	logrus.WithFields(logrus.Fields{
		"总数": TargetCombination,
		"概率": float64(TargetCombination) / float64(len(combinations)),
	}).Debugf("取到 %v 组合的信息", Hand)
}
