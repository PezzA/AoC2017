package puzzles

import (
	"errors"
	"github.com/pezza/advent-of-code/2020/Day202020"
	"github.com/pezza/advent-of-code/2020/Day202021"
	"github.com/pezza/advent-of-code/2020/Day202022"
	"github.com/pezza/advent-of-code/2020/Day202023"
	"github.com/pezza/advent-of-code/2020/Day202024"
	"github.com/pezza/advent-of-code/2020/Day202025"

	"github.com/pezza/advent-of-code/2020/Day202018"
	"github.com/pezza/advent-of-code/2020/Day202019"

	"github.com/pezza/advent-of-code/2020/Day202017"

	"github.com/pezza/advent-of-code/2020/Day202016"

	"github.com/pezza/advent-of-code/2020/Day202015"

	"github.com/pezza/advent-of-code/2020/Day202014"

	"github.com/pezza/advent-of-code/2020/Day202013"

	"github.com/pezza/advent-of-code/2020/Day202012"

	"github.com/pezza/advent-of-code/2020/Day202011"

	"github.com/pezza/advent-of-code/2020/Day202010"

	"github.com/pezza/advent-of-code/2020/Day202009"

	"github.com/pezza/advent-of-code/2020/Day202007"
	"github.com/pezza/advent-of-code/2020/Day202008"

	"github.com/pezza/advent-of-code/2020/Day202005"
	"github.com/pezza/advent-of-code/2020/Day202006"

	"github.com/pezza/advent-of-code/2020/Day202004"

	"github.com/pezza/advent-of-code/2020/Day202001"
	"github.com/pezza/advent-of-code/2020/Day202002"
	"github.com/pezza/advent-of-code/2020/Day202003"

	"github.com/pezza/advent-of-code/2015/Day201501"
	"github.com/pezza/advent-of-code/2015/Day201502"
	"github.com/pezza/advent-of-code/2015/Day201503"
	"github.com/pezza/advent-of-code/2015/Day201504"
	"github.com/pezza/advent-of-code/2015/Day201505"
	"github.com/pezza/advent-of-code/2015/Day201506"
	"github.com/pezza/advent-of-code/2015/Day201507"
	"github.com/pezza/advent-of-code/2015/Day201508"
	"github.com/pezza/advent-of-code/2015/Day201509"
	"github.com/pezza/advent-of-code/2015/Day201510"
	"github.com/pezza/advent-of-code/2015/Day201511"
	"github.com/pezza/advent-of-code/2015/Day201512"
	"github.com/pezza/advent-of-code/2015/Day201513"
	"github.com/pezza/advent-of-code/2015/Day201514"
	"github.com/pezza/advent-of-code/2015/Day201515"
	"github.com/pezza/advent-of-code/2015/Day201516"
	"github.com/pezza/advent-of-code/2015/Day201517"
	"github.com/pezza/advent-of-code/2015/Day201518"
	"github.com/pezza/advent-of-code/2015/Day201519"
	"github.com/pezza/advent-of-code/2015/Day201520"
	"github.com/pezza/advent-of-code/2015/Day201521"
	"github.com/pezza/advent-of-code/2015/Day201522"
	"github.com/pezza/advent-of-code/2015/Day201523"
	"github.com/pezza/advent-of-code/2015/Day201524"
	"github.com/pezza/advent-of-code/2015/Day201525"
	"github.com/pezza/advent-of-code/2016/Day201601"
	"github.com/pezza/advent-of-code/2016/Day201602"
	"github.com/pezza/advent-of-code/2016/Day201603"
	"github.com/pezza/advent-of-code/2016/Day201604"
	"github.com/pezza/advent-of-code/2016/Day201611"
	"github.com/pezza/advent-of-code/2016/Day201617"
	"github.com/pezza/advent-of-code/2016/Day201619"
	"github.com/pezza/advent-of-code/2016/Day201622"
	"github.com/pezza/advent-of-code/2017/Day201701"
	"github.com/pezza/advent-of-code/2017/Day201702"
	"github.com/pezza/advent-of-code/2017/Day201703"
	"github.com/pezza/advent-of-code/2017/Day201704"
	"github.com/pezza/advent-of-code/2017/Day201705"
	"github.com/pezza/advent-of-code/2017/Day201706"
	"github.com/pezza/advent-of-code/2017/Day201707"
	"github.com/pezza/advent-of-code/2017/Day201708"
	"github.com/pezza/advent-of-code/2017/Day201709"
	"github.com/pezza/advent-of-code/2017/Day201710"
	"github.com/pezza/advent-of-code/2017/Day201711"
	"github.com/pezza/advent-of-code/2017/Day201712"
	"github.com/pezza/advent-of-code/2017/Day201713"
	"github.com/pezza/advent-of-code/2017/Day201714"
	"github.com/pezza/advent-of-code/2017/Day201715"
	"github.com/pezza/advent-of-code/2017/Day201716"
	"github.com/pezza/advent-of-code/2017/Day201717"
	"github.com/pezza/advent-of-code/2017/Day201718"
	"github.com/pezza/advent-of-code/2017/Day201719"
	"github.com/pezza/advent-of-code/2017/Day201720"
	"github.com/pezza/advent-of-code/2017/Day201721"
	"github.com/pezza/advent-of-code/2017/Day201722"
	"github.com/pezza/advent-of-code/2017/Day201723"
	"github.com/pezza/advent-of-code/2017/Day201724"
	"github.com/pezza/advent-of-code/2017/Day201725"
	"github.com/pezza/advent-of-code/2018/Day201801"
	"github.com/pezza/advent-of-code/2018/Day201802"
	"github.com/pezza/advent-of-code/2018/Day201803"
	"github.com/pezza/advent-of-code/2018/Day201804"
	"github.com/pezza/advent-of-code/2018/Day201805"
	"github.com/pezza/advent-of-code/2018/Day201806"
	"github.com/pezza/advent-of-code/2018/Day201807"
	"github.com/pezza/advent-of-code/2018/Day201808"
	"github.com/pezza/advent-of-code/2018/Day201809"
	"github.com/pezza/advent-of-code/2018/Day201810"
	"github.com/pezza/advent-of-code/2018/Day201811"
	"github.com/pezza/advent-of-code/2018/Day201812"
	"github.com/pezza/advent-of-code/2018/Day201813"
	"github.com/pezza/advent-of-code/2018/Day201814"
	"github.com/pezza/advent-of-code/2018/Day201815"
	"github.com/pezza/advent-of-code/2018/Day201816"
	"github.com/pezza/advent-of-code/2018/Day201817"
	"github.com/pezza/advent-of-code/2018/Day201818"
	"github.com/pezza/advent-of-code/2018/Day201819"
	"github.com/pezza/advent-of-code/2018/Day201820"
	"github.com/pezza/advent-of-code/2018/Day201823"
	"github.com/pezza/advent-of-code/2018/Day201825"
	"github.com/pezza/advent-of-code/2019/Day201901"
	"github.com/pezza/advent-of-code/2019/Day201902"
	"github.com/pezza/advent-of-code/2019/Day201903"
	"github.com/pezza/advent-of-code/2019/Day201904"
	"github.com/pezza/advent-of-code/2019/Day201905"
	"github.com/pezza/advent-of-code/2019/Day201906"
	"github.com/pezza/advent-of-code/2019/Day201907"
	"github.com/pezza/advent-of-code/2019/Day201908"
	"github.com/pezza/advent-of-code/2019/Day201909"
	"github.com/pezza/advent-of-code/2019/Day201910"
	"github.com/pezza/advent-of-code/2019/Day201911"
)

func GetPuzzle(day int, year int) (DailyPuzzle, error) {
	dailyPuzzles := [...]DailyPuzzle{
		Day201501.Entry,
		Day201502.Entry,
		Day201503.Entry,
		Day201504.Entry,
		Day201505.Entry,
		Day201506.Entry,
		Day201507.Entry,
		Day201508.Entry,
		Day201509.Entry,
		Day201510.Entry,
		Day201511.Entry,
		Day201512.Entry,
		Day201513.Entry,
		Day201514.Entry,
		Day201515.Entry,
		Day201516.Entry,
		Day201517.Entry,
		Day201518.Entry,
		Day201519.Entry,
		Day201520.Entry,
		Day201521.Entry,
		Day201522.Entry,
		Day201523.Entry,
		Day201524.Entry,
		Day201525.Entry,
		Day201601.Entry,
		Day201602.Entry,
		Day201603.Entry,
		Day201604.Entry,
		Day201611.Entry,
		Day201617.Entry,
		Day201619.Entry,
		Day201622.Entry,
		Day201701.Entry,
		Day201702.Entry,
		Day201703.Entry,
		Day201704.Entry,
		Day201705.Entry,
		Day201706.Entry,
		Day201707.Entry,
		Day201708.Entry,
		Day201709.Entry,
		Day201710.Entry,
		Day201711.Entry,
		Day201712.Entry,
		Day201713.Entry,
		Day201714.Entry,
		Day201715.Entry,
		Day201716.Entry,
		Day201717.Entry,
		Day201718.Entry,
		Day201719.Entry,
		Day201720.Entry,
		Day201721.Entry,
		Day201722.Entry,
		Day201723.Entry,
		Day201724.Entry,
		Day201725.Entry,
		Day201801.Entry,
		Day201802.Entry,
		Day201803.Entry,
		Day201804.Entry,
		Day201805.Entry,
		Day201806.Entry,
		Day201807.Entry,
		Day201808.Entry,
		Day201809.Entry,
		Day201810.Entry,
		Day201811.Entry,
		Day201812.Entry,
		Day201813.Entry,
		Day201814.Entry,
		Day201815.Entry,
		Day201816.Entry,
		Day201817.Entry,
		Day201818.Entry,
		Day201819.Entry,
		Day201820.Entry,
		Day201823.Entry,
		Day201825.Entry,
		Day201901.Entry,
		Day201902.Entry,
		Day201903.Entry,
		Day201904.Entry,
		Day201905.Entry,
		Day201906.Entry,
		Day201907.Entry,
		Day201908.Entry,
		Day201909.Entry,
		Day201910.Entry,
		Day201911.Entry,
		Day202001.Entry,
		Day202002.Entry,
		Day202003.Entry,
		Day202004.Entry,
		Day202005.Entry,
		Day202006.Entry,
		Day202007.Entry,
		Day202008.Entry,
		Day202009.Entry,
		Day202010.Entry,
		Day202011.Entry,
		Day202012.Entry,
		Day202013.Entry,
		Day202014.Entry,
		Day202015.Entry,
		Day202016.Entry,
		Day202017.Entry,
		Day202018.Entry,
		Day202019.Entry,
		Day202020.Entry,
		Day202021.Entry,
		Day202022.Entry,
		Day202023.Entry,
		Day202024.Entry,
		Day202025.Entry,
	}

	for _, puzzle := range dailyPuzzles {
		puzzleYear, puzzleDay, _ := puzzle.Describe()
		if puzzleDay == day && puzzleYear == year {
			return puzzle, nil
		}
	}

	return nil, errors.New("day specified has not been fully implemented yet")
}
