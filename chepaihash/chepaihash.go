package chepaihash

import "fmt"

const (
	PROVINCE     = "京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼"
	ALPHABET     = "ABCDEFGHJKLMNPQRSTUVWXYZ"
	ALPHANUMERIC = "ABCDEFGHJKLMNPQRSTUVWXYZ0123456789"
)

type LinearCongruentialRng struct {
	seed uint64
}

func NewLinearCongruentialRng(seed uint64) *LinearCongruentialRng {
	if seed == 0 {
		seed = 1 // 避免种子为0
	}
	return &LinearCongruentialRng{seed: seed}
}

func (rng *LinearCongruentialRng) Next() uint64 {
	const (
		LCG_A = 6364136223846793005
		LCG_C = 1442695040888963407
	)
	rng.seed = LCG_A*rng.seed + LCG_C
	return rng.seed
}

func Hash(value string) (string, error) {
	if value == "" {
		return "", fmt.Errorf("输入字符串不能为空")
	}

	seed := uint64(0)
	for _, c := range value {
		seed = seed*31 + uint64(c)
	}

	rng := NewLinearCongruentialRng(seed)

	chepai := make([]rune, 8)
	chepai[0] = getRandomChar(rng, PROVINCE)
	chepai[1] = getRandomChar(rng, ALPHABET)
	chepai[2] = '·'
	for i := 0; i < 5; i++ {
		chepai[3+i] = getRandomChar(rng, ALPHANUMERIC)
	}

	return string(chepai), nil
}

func getRandomChar(rng *LinearCongruentialRng, chars string) rune {
	return []rune(chars)[int(rng.Next()%uint64(len([]rune(chars))))]
}
