package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

// Yield information
type Yield struct {
	MachineName      string  `orm:"column(machine_name)" json:"machineName"`
	MachineNumber    string  `orm:"column(machine_number)" json:"machineNumber"`
	GreenProduction  float64 `orm:"column(green_production)" json:"greenProduction"`
	YellowProduction float64 `orm:"column(yellow_production)" json:"yellowProduction"`
	RedProduction    float64 `orm:"column(red_production)" json:"redProduction"`
}

// YieldTwo information
type YieldTwo struct {
	GreenProduction  float64 `orm:"column(green_production)" json:"greenProduction"`
	YellowProduction float64 `orm:"column(yellow_production)" json:"yellowProduction"`
	RedProduction    float64 `orm:"column(red_production)" json:"redProduction"`
	Timestamp        int64   `orm:"column(timestamp)" json:"timestamp"`
}

// ErrorResponse error response
type ErrorResponse struct {
	ErrorRes string
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:82589155@tcp(10.10.10.100:3306)/iom5?charset=utf8") //註冊資料庫
}

func main() {
	ginEngine().Run()
}

func ginEngine() *gin.Engine {
	r := gin.Default()
	r.GET("/findYieldByMachineID", findYieldByMachineIDHandler)
	r.GET("/findYieldPerHourByWorkShopID", findYieldPerHourByWorkShopIDHendler)
	return r
}

func findYieldByMachineIDHandler(c *gin.Context) {
	workShopID := c.GetHeader("wsID")
	startTime := c.GetHeader("startTime")
	endTime := c.GetHeader("endTime")

	if results, err := findYieldByMachineID(workShopID, startTime, endTime); err != nil {
		var errResponse ErrorResponse
		errResponse.ErrorRes = err.Error()
		c.JSON(http.StatusInternalServerError, errResponse)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

func findYieldPerHourByWorkShopIDHendler(c *gin.Context) {
	workShopID := c.GetHeader("wsID")
	startTime := c.GetHeader("startTime")
	endTime := c.GetHeader("endTime")
	if result, err := findYieldPerHourByWorkShopID(workShopID, startTime, endTime); err != nil {
		var errResponse ErrorResponse
		errResponse.ErrorRes = err.Error()
		c.JSON(http.StatusInternalServerError, errResponse)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

func findYieldByMachineID(workShopID, startTime, endTime string) (results []Yield, err error) {
	o := orm.NewOrm()
	IDint64, err := strconv.ParseInt(workShopID, 10, 64)
	if err != nil {
		return nil, err
	}

	start, err := strconv.ParseInt(startTime, 10, 64)
	if err != nil {
		return nil, err
	}

	end, err := strconv.ParseInt(endTime, 10, 64)
	if err != nil {
		return nil, err
	}

	results, err = getYieldByMachineID(start, end, IDint64, o)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func findYieldPerHourByWorkShopID(wsID, start, end string) (results []YieldTwo, err error) {
	o := orm.NewOrm()
	const anHour = int64(time.Hour / time.Millisecond)

	wsIDInt64, err := strconv.ParseInt(wsID, 10, 64)
	if err != nil {
		fmt.Println("1")
		return nil, err
	}
	startTime, err := strconv.ParseInt(start, 10, 64)
	if err != nil {
		fmt.Println("2")
		return nil, err
	}
	endTime, err := strconv.ParseInt(end, 10, 64)
	if err != nil {
		fmt.Println("3")
		return nil, err
	}
	dataArr, err := getRangeDataByMachineIDAndStartAndEnd(startTime, endTime, wsIDInt64, o)
	if err != nil {
		fmt.Println("4")
		return nil, err
	}
	tmpMap := make(map[int64]YieldTwo)
	for _, v := range dataArr {
		var tmpTime int64
		if v.Timestamp%anHour == 0 {
			tmpTime = v.Timestamp
		} else {
			tmpTime = startTime + (anHour - v.Timestamp%anHour)
		}
		tmpMapp := tmpMap[tmpTime]
		tmpMapp.GreenProduction += v.GreenProduction
		tmpMapp.YellowProduction += v.YellowProduction
		tmpMapp.RedProduction += v.RedProduction
		tmpMap[tmpTime] = tmpMapp
	}
	for k, v := range tmpMap {
		v.Timestamp = k
		results = append(results, v)
	}
	return
}

func getYieldByMachineID(start, end, workShopID int64, o orm.Ormer) (results []Yield, err error) {
	sql := "SELECT bs.machine_number, bs.machine_name, an.green_production, an.yellow_production , an.red_production from iom5.iom5_basic_machine as bs Left JOIN (select machine_number,sum(green_production) as 'green_production', sum(yellow_production) as 'yellow_production', sum(red_production) as 'red_production' from iom5.iom5_analyze where timestamp between ? and ? group by machine_number) as an on bs.machine_number = an.machine_number where bs.work_shop_id = ? group by bs.machine_number;"
	_, err = o.Raw(sql, start, end, workShopID).QueryRows(&results)
	return
}

func getRangeDataByMachineIDAndStartAndEnd(start, end, wsID int64, o orm.Ormer) (YieldArr []YieldTwo, err error) {
	sql := "select bs.work_shop_id,an.machine_number, an.green_production, an.yellow_production, an.red_production, an.timestamp from iom5_analyze as an left join iom5.iom5_basic_machine as bs on an.machine_number = bs.machine_number where bs.work_shop_id =? and timestamp between ? and ?"
	_, err = o.Raw(sql, wsID, start, end).QueryRows(&YieldArr)
	return
}
