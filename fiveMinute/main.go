package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//FiveMinute Output
type FiveMinute struct {
	StopTime     int64 `orm:"column(stop_time)"`
	IdleTime     int64 `orm:"column(idle_time)"`
	AbnormalTime int64 `orm:"column(abnormal_time)"`
	RunningTime  int64 `orm:"column(running_time)"`
	Status       int   `orm:"column(status)"`
	Timestamp    int64 `orm:"column(timestamp)"`
}

//Status Input
type Status struct {
	ID        int64   `orm:"column(id)" json:"id"`
	CycleTime float64 `orm:"column(cycle_time)" json:"cycleTime"`
	Status    int     `orm:"column(status)" json:"status"`
	Timestamp int64   `orm:"column(timestamp)" json:"timestamp"`
	GYR       string  `orm:"column(gyr)" json:"gyr"`
	SCT       float64 `orm:"column(sct)" json:"sct"`
}

//TimeSections times
type TimeSections struct {
	start int64
	end   int64
}

//StatusData including statusArry
type StatusData struct {
	statusData   []Status
	IdleTime     int64
	RunningTime  int64
	StopTime     int64
	AbnormalTime int64
}

//ErrorResponse include error string
type ErrorResponse struct {
	ErrString string
}

const stopStatus int = 2
const idleStatus int = 3
const abnormalStatus int = 4
const runningStatus int = 5

func init() {
	orm.RegisterDataBase("default", "mysql", "root:82589155@tcp(10.10.10.100:3306)/iom5?charset=utf8")
}

func main() {
	ginEngine().Run()
	fmt.Println(getSectionsNumber(1581058200000, 1581061500000))
	// fiveMinutesAnalyse("1581555000000", "1581558300000")
}

func ginEngine() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/fiveMinuteAnalyseHandler", fiveMinutesAnalyseHandler)
	return r
}

func fiveMinutesAnalyseHandler(c *gin.Context) {
	start := c.GetHeader("startTime")
	end := c.GetHeader("endTime")
	if result, err := fiveMinutesAnalyse(start, end); err != nil {
		var errResponse ErrorResponse
		errResponse.ErrString = err.Error()
		c.JSON(500, errResponse)
	} else {
		c.JSON(200, result)
	}
}

// func ireallywantu(lastStatus Status, data *StatusData, val int64) {
// 	switch lastStatus.Status {
// 	case stopStatus:
// 		data.StopTime += val
// 	case idleStatus:
// 		data.IdleTime += val
// 	case abnormalStatus:
// 		data.AbnormalTime += val
// 	case runningStatus:
// 		data.RunningTime += val
// 	}
// }

func fiveMinutesAnalyse(start, end string) (fiveMinutesResult []FiveMinute, err error) {
	o := orm.NewOrm()
	var startTime, endTime int64

	startTime, err = strconv.ParseInt(start, 10, 64)
	if err != nil {
		return nil, err
	}
	endTime, err = strconv.ParseInt(end, 10, 64)
	if err != nil {
		return nil, err
	}

	var fiveMin int64
	fiveMin = int64(5 * time.Minute / time.Millisecond)
	var statusArr []StatusData

	dataArr, err := getAlldataFromStatusTableInPeriodM01(startTime, endTime, o)
	if err != nil {
		return
	}

	lastStatus, err := getLastStatusDataBeforeATime(startTime, o)
	if err != nil && err != orm.ErrNoRows {
		return
	}
	sections := int(getSectionsNumber(startTime, endTime))

	for i := 0; i < sections; i++ {
		statusArr = append(statusArr, StatusData{})
		fiveMinutesResult = append(fiveMinutesResult, FiveMinute{})
	}

	for _, value := range dataArr {
		location := (value.Timestamp - startTime) / fiveMin
		statusArr[location].statusData = append(statusArr[location].statusData, value)
	}

	if len(dataArr) > 0 {
		judgeTime := endTime - dataArr[len(dataArr)-1].Timestamp
		if judgeTime > int64(time.Hour/time.Millisecond) {
			err = errors.New("not enough data")
			return nil, err
		}
	} else {
		err = errors.New("no data")
		return nil, err
	}

	var flag int
	var previousStatusPoint Status
	previousStatusPoint = lastStatus
	for i, value := range statusArr {
		var time int64
		if len(value.statusData) == 0 {
			if i == 0 && flag == 0 {
				flag = 1
			} else if flag == 0 {
				lastStatus = previousStatusPoint
			}
			time = fiveMin
		} else {
			firstTimeStamp := value.statusData[0].Timestamp
			time = (firstTimeStamp - startTime) % fiveMin
			if flag == 0 {
				lastStatus = previousStatusPoint
			}
			previousStatusPoint = statusArr[i].statusData[len(statusArr[i].statusData)-1]
			flag = 0
		}

		// ireallywantu(lastStatus, &value, time)

		switch lastStatus.Status { //處理第一筆時間點之前的資料
		case stopStatus:
			statusArr[i].StopTime += time
		case idleStatus:
			statusArr[i].IdleTime += time
		case abnormalStatus:
			statusArr[i].AbnormalTime += time
		case runningStatus:
			statusArr[i].RunningTime += time
		}

		for number := range statusArr[i].statusData {
			var time int64
			if number == len(statusArr[i].statusData)-1 {
				tempEnd := startTime + fiveMin + fiveMin*int64(i)
				time = tempEnd - statusArr[i].statusData[number].Timestamp
			} else {
				time = statusArr[i].statusData[number+1].Timestamp - statusArr[i].statusData[number].Timestamp
			}

			switch statusArr[i].statusData[number].Status {
			case stopStatus:
				statusArr[i].StopTime += time
			case idleStatus:
				statusArr[i].IdleTime += time
			case abnormalStatus:
				statusArr[i].AbnormalTime += time
			case runningStatus:
				statusArr[i].RunningTime += time
			}
		}
	}

	for i := 0; i < len(statusArr); i++ {
		tempMap := make(map[int]int64)
		fiveMinutesResult[i].AbnormalTime = statusArr[i].AbnormalTime
		fiveMinutesResult[i].IdleTime = statusArr[i].IdleTime
		fiveMinutesResult[i].RunningTime = statusArr[i].RunningTime
		fiveMinutesResult[i].StopTime = statusArr[i].StopTime
		fiveMinutesResult[i].Timestamp = startTime + int64(i)*fiveMin + fiveMin
		tempMap[abnormalStatus] = fiveMinutesResult[i].AbnormalTime
		tempMap[idleStatus] = fiveMinutesResult[i].IdleTime
		tempMap[stopStatus] = fiveMinutesResult[i].StopTime
		tempMap[runningStatus] = fiveMinutesResult[i].RunningTime
		var tempNumber int64
		for number, value := range tempMap {
			if value > tempNumber {
				tempNumber = value
				fiveMinutesResult[i].Status = number
			}
		}

	}

	fmt.Println(statusArr[5].statusData)

	return
}

func getSectionsNumber(start, end int64) (section int64) {
	var fiveMin int64
	fiveMin = int64(5 * time.Minute / time.Millisecond)
	section = int64((end - start) / fiveMin)

	return
}

func getAlldataFromStatusTableInPeriodM01(start, end int64, o orm.Ormer) (dataArr []Status, err error) {
	sql := "select * from iom5_data_status_m02 where timestamp>=? and timestamp <? order by timestamp asc"
	_, err = o.Raw(sql, start, end).QueryRows(&dataArr)
	if err != nil {
		return
	}
	return
}

func getLastStatusDataBeforeATime(start int64, o orm.Ormer) (lastStatus Status, err error) {
	sql := "SELECT * FROM iom5_data_status_m02 where timestamp < ? order by timestamp desc limit 1;"
	err = o.Raw(sql, start).QueryRow(&lastStatus)
	if err != nil {
		return
	}
	return
}

// func getFirstStatusDataAfterATime(start int64, o orm.Ormer) (firstStatus Status, err error) {
// 	sql := "SELECT * FROM iom5_data_status_m01 where timestamp > ? order by timestamp asc limit 1;"
// 	err = o.Raw(sql, start).QueryRow(&firstStatus)
// 	if err != nil {
// 		fmt.Println("wrong", err)

// 	}
// 	return
// }
