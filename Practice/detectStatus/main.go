package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MachineIO struct { //MachineIO的結構
	ID         int64
	Timestamp  int64
	Di0        int
	Di1        int
	Di2        int
	Di3        int
	Di4        int
	Di5        int
	Di6        int
	Di7        int
	Analyzed   int
	CreateTime int64
}
type Status struct { //Status的結構
	ID        int64
	CycleTime float64 //s
	Status    int
	Timestamp int64 //ms

}

var signalOne = MachineIO{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}    //給出不同的fourdig資訊	0000*	 300010
var signalTwo = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}    //0000*							300290
var signalThree = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0}  //1100							303700
var signalFour = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0}   //1010							346000
var signalFive = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}   //1100							250000
var signalSix = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0}    //1010							500
var signalSeven = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}  //1100							400
var signalEight = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0}  //1010							4100
var signalNine = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}   //1100							5000
var signalNinee = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}  //1010							310000
var signalEightt = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0} //1100							310000
var signalSevenn = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0} //1010							100
var signalSixx = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}   //1100							400
var signalFivee = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0}  //1010							300
var signalFourr = MachineIO{0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}  //1101*							200
var signalThreee = MachineIO{0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0} //1011*							4000
var signalTwoo = MachineIO{0, 0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0}   //1101*							1000
var signalOnee = MachineIO{0, 0, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0}   //1011*

func main() {
	//testdata := []MachineIO{signalOne, signalTwo, signalThree, signalFour, signalFive, signalSix, signalSeven, signalEight, signalNine, signalNinee, signalEightt, signalSevenn, signalSixx, signalFivee, signalFourr, signalThreee, signalTwoo, signalOnee}

	var sNewNewData []MachineIO //打亂後的順序
	//var statusInfo []Status
	//statusInfo = make([]Status, 50)
	fmt.Print(numberOfDataYouWant(50, sNewNewData))
	fmt.Println(buildStatus(numberOfDataYouWant(50, sNewNewData)))
	//fmt.Print(testdata)
	//fmt.Print(buildStatus(testdata))
	//fmt.Println(statusInfo)
	//Output(sNewNewData)

}

func numberOfDataYouWant(t int, sNewNewData []MachineIO) []MachineIO {

	var signalTime []int64 //int64的訊號時間

	t1 := time.Now()
	rand.Seed(time.Now().Unix())
	data := []MachineIO{signalOne, signalSeven, signalThree, signalFive, signalFour, signalSix, signalTwo, signalNine, signalEight, signalOnee, signalTwoo, signalThreee, signalFourr, signalFivee, signalSixx, signalSevenn, signalEightt, signalNinee}

	for i := 0; i < t; i++ {

		sNewNewData = append(sNewNewData, data[rand.Intn(17)])
		sNewNewData[i].ID = int64(i)
		elapsed := time.Since(t1)
		var int64_time int64 = elapsed.Milliseconds()
		signalTime = append(signalTime, int64_time)
		time.Sleep(1 * time.Millisecond)
		sNewNewData[i].Timestamp = signalTime[i]

		//fmt.Println("000", signalTime[i])
		//signalTime[i] = newNewData[i].Timestamp
		//fmt.Println(signalTime[i])
		//fmt.Println(newNewData[i])

	}
	//var endTime = int64(time.Since(t1) / 1000000) //- sNewNewData[t].Timestamp
	//sNewNewData = append(sNewNewData, MachineIO{0, endTime, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	//fmt.Println(sNewNewData)
	return sNewNewData
}

func buildStatus(sNewNewData []MachineIO) []Status {

	const fiveMinute int64 = 300000
	var seriesCycleTime int64

	//var statusInfo = make([]Status, len(sNewNewData))

	var statusInfo []Status

	for i := 0; i < len(sNewNewData)-2; i++ {
		if sNewNewData[i].Di0 == 0 {
			statusInfo = append(statusInfo, Status{int64(i), 0, 2, sNewNewData[i].Timestamp})

		} else if sNewNewData[i].Di3 == 1 {
			statusInfo = append(statusInfo, Status{int64(i), 0, 4, sNewNewData[i].Timestamp})

		} else if sNewNewData[i+1].Timestamp-sNewNewData[i].Timestamp > fiveMinute {
			statusInfo = append(statusInfo, Status{int64(i), 0, 3, sNewNewData[i].Timestamp})
		} else {
			if sNewNewData[i].Di1 == 0 && sNewNewData[i].Di2 == 1 { //查看第一組是不是1010,如果是的話
				if sNewNewData[i+1].Di1 == 1 && sNewNewData[i+1].Di2 == 0 { //查看第二組是不是1100,如果是的話
					if sNewNewData[i+2].Di1 == 0 && sNewNewData[i+2].Di2 == 1 { //查看第三組是不是1010,如果是的話

						if i < len(sNewNewData)-3 { //如果i不是最後第4筆資料 則

							if sNewNewData[i+3].Timestamp-sNewNewData[i+2].Timestamp < fiveMinute && sNewNewData[i+2].Timestamp-sNewNewData[i+1].Timestamp < fiveMinute && sNewNewData[i+1].Timestamp-sNewNewData[i].Timestamp < fiveMinute { //查看這三組是不是小於5分鐘,如果是的話
								seriesCycleTime = sNewNewData[i+2].Timestamp - sNewNewData[i].Timestamp
								statusInfo = append(statusInfo, Status{int64(i), float64(seriesCycleTime), 5, sNewNewData[i+2].Timestamp})
								i = i + 2
							} else {
								//判斷這三組裡面有閒置
								for j := 0; j < 3; j++ {
									seriesCycleTime = sNewNewData[i+j+1].Timestamp - sNewNewData[i+j].Timestamp
									if seriesCycleTime > fiveMinute {
										statusInfo = append(statusInfo, Status{int64(i + j), 0, 3, sNewNewData[i+j].Timestamp})
										fmt.Println(i, j)
									}
								}
								i = i + 2
							}
						} else { //如果i是最後第4筆資料 則
							seriesCycleTime = sNewNewData[i+2].Timestamp - sNewNewData[i].Timestamp
							statusInfo = append(statusInfo, Status{int64(i), float64(seriesCycleTime), 5, sNewNewData[i+2].Timestamp}) //cycletime

							for j := 0; j < 2; j++ {
								seriesCycleTime = sNewNewData[i+j+1].Timestamp - sNewNewData[i+j].Timestamp
								if seriesCycleTime > fiveMinute {
									statusInfo = append(statusInfo, Status{int64(i + j), 0, 3, sNewNewData[i].Timestamp})
								}
							}
						}

					} else {
						//判斷第三組的狀態是2還是4
						if sNewNewData[i+2].Di0 == 0 {
							statusInfo = append(statusInfo, Status{int64(i + 2), 0, 2, sNewNewData[i+2].Timestamp})
						} else if sNewNewData[i+2].Di3 == 1 {
							statusInfo = append(statusInfo, Status{int64(i + 2), 0, 4, sNewNewData[i+2].Timestamp})
						} else {
							//***判斷是閒置(第一組跟第二組)
							for j := 0; j < 2; j++ {
								seriesCycleTime = sNewNewData[i+j+1].Timestamp - sNewNewData[i+j].Timestamp
								if seriesCycleTime > fiveMinute {
									statusInfo = append(statusInfo, Status{int64(i + j), 0, 3, sNewNewData[i].Timestamp})
								}
							}
						}
						i = i + 1
					}
				} else {
					//判斷第二組的狀態是4還是2
					if sNewNewData[i+1].Di0 == 0 {
						statusInfo = append(statusInfo, Status{int64(i + 1), 0, 2, sNewNewData[i+2].Timestamp})
					} else if sNewNewData[i+1].Di3 == 1 {
						statusInfo = append(statusInfo, Status{int64(i + 1), 0, 4, sNewNewData[i+2].Timestamp})
					} else {
						//***判斷是閒置(第一組)
						seriesCycleTime = sNewNewData[i+1].Timestamp - sNewNewData[i].Timestamp
						if seriesCycleTime > fiveMinute {
							statusInfo = append(statusInfo, Status{int64(i), 0, 3, sNewNewData[i].Timestamp})
						}

					}
					i = i + 1
				}
			}

		}

	}

	if sNewNewData[len(sNewNewData)-2].Di0 == 0 {
		statusInfo = append(statusInfo, Status{int64(len(sNewNewData) - 2), 0, 2, sNewNewData[len(sNewNewData)-2].Timestamp})
	} else if sNewNewData[len(sNewNewData)-2].Di3 == 1 {
		statusInfo = append(statusInfo, Status{int64(len(sNewNewData) - 2), 0, 4, sNewNewData[len(sNewNewData)-2].Timestamp})
	} else {
		seriesCycleTime = sNewNewData[len(sNewNewData)-1].Timestamp - sNewNewData[len(sNewNewData)-2].Timestamp
		if seriesCycleTime > fiveMinute {
			statusInfo = append(statusInfo, Status{int64(len(sNewNewData) - 2), 0, 3, sNewNewData[len(sNewNewData)-2].Timestamp})
		}
	}

	if sNewNewData[len(sNewNewData)-1].Di3 == 1 {
		statusInfo = append(statusInfo, Status{int64(len(sNewNewData) - 1), 0, 4, sNewNewData[len(sNewNewData)-1].Timestamp})
	}

	return statusInfo
}

/*
if sNewNewData[len(sNewNewData)-1].Di0 == 0 {
		statusInfo = append(statusInfo, Status{int64(len(sNewNewData) - 1), 0, 2, sNewNewData[len(sNewNewData)-1].Timestamp})

	} else */
