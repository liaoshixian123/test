package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	// mysql driver

	"net/url"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type dbSetting struct {
	alias     string
	name      string
	user      string
	pwd       string
	host      string
	port      string
	charset   string
	timezone  string
	dataRange int
}

type iom5MoldOld struct {
	ID                 int64  `orm:"column(id)"`
	Seq                int64  `orm:"column(seq)"`
	ClientName         string `orm:"column(client_name)"`
	Name               string `orm:"column(name)"`
	PlasticType        string `orm:"column(plastic_type)"`
	ProductModel       string `orm:"column(product_model)"`
	Number             string `orm:"column(number)"`
	PlanCycleTime      string `orm:"column(plan_cycle_time)"`
	Created            int64  `orm:"column(created)"`
	GreenRange         string `orm:"column(green_range)"`
	YellowRange        string `orm:"column(yellow_range)"`
	MiddenumFirstRange string `orm:"column(middlenum_firstrange)"`
}

type iom5BasicMold struct {
	ID                     int64   `orm:"column(id)"`
	Name                   string  `orm:"column(name)"`
	ProductModel           string  `orm:"column(product_model)"`
	Number                 string  `orm:"column(number)"`
	TransferInDate         int64   `orm:"column(transfer_in_date)"`
	ManufactureDate        int64   `orm:"column(manufacture_date)"`
	Height                 float64 `orm:"column(height)"`
	Width                  float64 `orm:"column(width)"`
	MovableMoldThickness   float64 `orm:"column(movable_mold_thickness)"`
	FixedMoldThickness     float64 `orm:"column(fixed_mold_thickness)"`
	CavityNumber           int64   `orm:"column(cavity_number)"`
	ProductType            string  `orm:"column(product_type)"`
	ShrinkageRate          float64 `orm:"column(shrinkage_rate)"`
	ColdRunningWeight      float64 `orm:"column(cold_running_weight)"`
	HotRunningWeight       float64 `orm:"column(hot_running_weight)"`
	hoTrunningVolume       float64 `orm:"column(hot_running_volume)"`
	MoldStroke             float64 `orm:"column(mold_stroke)"`
	MinMoldStroke          float64 `orm:"column(min_mold_stroke)"`
	ColdRunningMaxDiameter float64 `orm:"column(cold_running_max_diameter)"`
	ColdRunningVolume      float64 `orm:"column(cold_running_volume)"`
	CompanyID              int64   `orm:"column(company_id)"`
	LastUpdateTime         int64   `orm:"column(last_update_time)"`
	GreenRange             float64 `orm:"column(green_range)"`
	YellowRange            float64 `orm:"column(yellow_range)"`
	CycleTime              float64 `orm:"column(cycle_time)"`
	UpTime                 int64   `orm:"column(up_time)"`
	DownTime               int64   `orm:"column(down_time)"`
	WorkShopID             int64   `orm:"column(work_shop_id)"`
	Status                 int64   `orm:"column(status)"`
}

type iom5ProductRecordOld struct {
	ID             int64  `orm:"column(id)"`
	SerialNumber   string `orm:"column(serial_number)"`
	StartTime      int64  `orm:"column(start_time)"`
	EndTime        int64  `orm:"column(end_time)"`
	NormalNumber   int64  `orm:"column(normal_number)"`
	AbnormalNumber int64  `orm:"column(abnormal_number)"`
	UpdateTime     int64  `orm:"column(update_time)"`
}

type iom5PostProductRecord struct {
	ID             int64  `orm:"column(id)"`
	SerialNumber   string `orm:"column(serial_number)"`
	ProductID      int64  `orm:"column(product_id)"` //N
	LastUpdateTime int64  `orm:"column(last_update_time)"`
	TaskID         int64  `orm:"column(task_id)"` //N
	StartTime      int64  `orm:"column(start_time)"`
	EndTime        int64  `orm:"column(end_time)"`
	NormalNumber   int64  `orm:"column(normal_number)"`
	AbnormalNumber int64  `orm:"column(abnormal_number)"`
}

type iom5ProductNgType struct {
	ID     int64  `orm:"column(id)"`
	Name   string `orm:"column(name)"`
	Delete bool   `orm:"column(delete)"`
}

type iom5BasicDefect struct {
	ID     int64  `orm:"column(id)"`
	Name   string `orm:"column(name)"`
	Remake string `orm:"column(remake)"`
	Delete bool   `orm:"column(delete)"`
	Type   int64  `orm:"column(type)"`
}

type iom5ProductRecordNgInfo struct {
	ID        int64 `orm:"column(id)"`
	ProductID int64 `orm:"column(product_id)"` //123
	NGTypeID  int64 `orm:"column(ng_type_id)"` //121
	Number    int64 `orm:"column(number)"`     //122
}

type iom5BasicDefectRel struct {
	ID       int64 `orm:"column(id)"`
	DefectID int64 `orm:"column(defect_id)"`
	Number   int64 `orm:"column(number)"`
	RecordID int64 `orm:"column(record_id)"`
}

type iom5Schedule struct {
	ID               int64   `orm:"column(id)"`
	ScheduleSerial   string  `orm:"column(schedule_serial)"`
	Seq              int64   `orm:"column(seq)"` //--
	Qty              string  `orm:"column(qty)"`
	MoldNumber       string  `orm:"column(mold_number)"`
	MachineName      string  `orm:"column(machine_name)"` //--
	TimeStart        string  `orm:"column(time_start)"`   //**
	TimeEnd          string  `orm:"column(time_end)"`     //**
	MachineNumber    string  `orm:"column(machine_number)"`
	MoldID           int64   `orm:"column(mold_id)"`
	MachineID        int64   `orm:"column(machine_id)"`
	Update           int64   `orm:"column(update)"`          //--
	TimeStartinput   string  `orm:"column(time_startinput)"` //--
	TimeEndinput     string  `orm:"column(time_endinput)"`   //--
	TimeStamp        int64   `orm:"column(time_stamp)"`      //--
	MacAddress       string  `orm:"column(mac_address)"`     //--
	Tr               float64 `orm:"column(tr)"`              //--
	GreenRange       string  `orm:"column(green_range)"`
	YellowRange      string  `orm:"column(yellow_range)"`
	MiddleFirstrange string  `orm:"column(middle_firstrange)"` //--
	ProductSerial    string  `orm:"column(product_serial)"`    //--
	//create time
	//status
	//delay time
	//material id
}

type iom5InfoSchedule struct {
	ID              int64   `orm:"column(id)"`
	ScheduleSerial  string  `orm:"column(schedule_serial)"`
	MoldID          int64   `orm:"column(mold_id)"`
	MoldGreenRange  float64 `orm:"column(mold_green_range)"`
	MoldYellowRange float64 `orm:"column(mold_yellow_range)"`
	MoldCycleTime   float64 `orm:"column(mold_cycle_time)"`
	MachineID       int64   `orm:"column(machine_id)"`
	MaterialID      int64   `orm:"column(material_id)"`
	StartTime       int64   `orm:"column(start_time)"`
	EndTime         int64   `orm:"column(end_time)"`
	DelayTime       int64   `orm:"column(delay_time)"`
	Update          bool    `orm:"column(update)"`
	Qty             int64   `orm:"column(qty)"`
	Status          int64   `orm:"column(status)"`
	CreatedTime     int64   `orm:"column(created_time)"`
}

type iom5BasicMachine struct {
	ID                         int64   `orm:"column(id)"`
	Brand                      string  `orm:"column(brand)"`
	MachineName                string  `orm:"column(machine_name)"`
	MachineNumber              string  `orm:"column(machine_number)"`
	Model                      string  `orm:"column(model)"`
	PurchasDate                int64   `orm:"column(purchas_date)"`
	ManufactureDate            int64   `orm:"column(manufacture_date)"`
	ScrewDiameter              float64 `orm:"column(screw_diameter)"`
	ScrewRatio                 float64 `orm:"column(screw_ratio)"`
	TheoreticalInjectionVolume float64 `orm:"column(theoretical_injection_volume)"`
	InjectionVolume            float64 `orm:"column(injection_volume)"`
	MaxInjectionPressure       float64 `orm:"column(max_injection_pressure)"`
	InjectionRate              float64 `orm:"column(injection_rate)"`
	InjectionSpeed             float64 `orm:"column(injection_speed)"`
	ShotStroke                 float64 `orm:"column(shot_stroke)"`
	ShootingStroke             float64 `orm:"column(shooting_stroke)"`
	NozzleClosedPower          float64 `orm:"column(nozzle_closed_power)"`
	HeatingSegmentNumber       float64 `orm:"column(heating_segment_number)"`
	TubeTotalHeat              float64 `orm:"column(tube_total_heat)"`
	ClampingPower              float64 `orm:"column(clamping_power)"`
	MaxOpenStroke              float64 `orm:"column(max_open_stroke)"`
	MinMoldThickness           float64 `orm:"column(min_mold_thickness)"`
	MaxMoldThickness           float64 `orm:"column(max_mold_thickness)"`
	MoldWidth                  float64 `orm:"column(mold_width)"`
	MoldHeight                 float64 `orm:"column(mold_height)"`
	MaxPlasticAmount           float64 `orm:"column(max_plastic_amount)"`
	LastUpdateTime             int64   `orm:"column(last_update_time)"`
	DcID                       int64   `orm:"column(dc_id)"`
	WorkShopID                 int64   `orm:"column(work_shop_id)"`
	Analyze                    bool    `orm:"column(analyze)"`
	Status                     int64   `orm:"column(status)"`
	MaterialType               int64   `orm:"column(material_type)"`
}

type machineOld struct {
	ID            int64  `orm:"column(id)"`
	MacAddress    string `orm:"column(mac_address)"`
	MachineName   string `orm:"column(machine_name)"`
	MachineNumber string `orm:"column(machine_number)"`
	Brand         string `orm:"column(brand)"`
	Seq           int64  `orm:"column(seq)"`
	Modified      string `orm:"column(modified)"`
	CollectionID  int64  `orm:"column(collection_id)"`
	Shot          int64  `orm:"column(shot)"`
	MaxPressure   int64  `orm:"column(max_pressure)"`
	LockPower     int64  `orm:"column(lock_power)"`
	MachineModel  string `orm:"column(machine_model)"`
}

//WorkShop work shop
type WorkShop struct {
	ID     int64  `orm:"column(id)"`
	Number string `orm:"column(number)"`
	Name   string `orm:"column(name)"`
	Type   int64  `orm:"column(type)"`
	Remark string `orm:"column(remark)"`
}

//Iom5AuthMailAccount email's infomation
type Iom5AuthMailAccount struct {
	ID              int64  `orm:"column(id)"`
	Email           string `orm:"column(email)"`
	Recipient       string `orm:"column(recipient)"`
	WorkShopSetting string `orm:"column(work_shop_setting)"`
}

//MailAccountOld old mail struct
type MailAccountOld struct {
	Parameter string `orm:"column(parameter)"`
	Value     string `orm:"column(value)"`
}

//SystemSetting system setting
type SystemSetting struct {
	Parameter string `orm:"column(parameter);pk"`
	Value     string `orm:"column(value)"`
}

// machine (OK)          	edge40 setting -> iom5 mail_account
// mold(OK)            work_shop - 預設一個 BDI
// product_ng_type(OK)
// product_record(OK)
// product_record_ng_info(OK)
// schedule(OK)
// system_setting

// const startIndex = 10
const tabName = "setting"

var fromSetting dbSetting
var toSetting dbSetting
var oldNewMap map[string]string

func init() {

	// oldNewMap = make(map[string]string)
	file, err := excelize.OpenFile("./db_setting.xlsx")
	if err != nil {
		fmt.Println("read setting file failed", err)
		return
	}

	fromSetting.alias, _ = file.GetCellValue(tabName, "B1")
	fromSetting.name, _ = file.GetCellValue(tabName, "B2")
	fromSetting.user, _ = file.GetCellValue(tabName, "B3")
	fromSetting.pwd, _ = file.GetCellValue(tabName, "B4")
	fromSetting.host, _ = file.GetCellValue(tabName, "B5")
	fromSetting.port, _ = file.GetCellValue(tabName, "B6")
	fromSetting.charset, _ = file.GetCellValue(tabName, "B7")
	fromSetting.timezone, _ = file.GetCellValue(tabName, "B8")

	toSetting.alias, _ = file.GetCellValue(tabName, "D1")
	toSetting.name, _ = file.GetCellValue(tabName, "D2")
	toSetting.user, _ = file.GetCellValue(tabName, "D3")
	toSetting.pwd, _ = file.GetCellValue(tabName, "D4")
	toSetting.host, _ = file.GetCellValue(tabName, "D5")
	toSetting.port, _ = file.GetCellValue(tabName, "D6")
	toSetting.charset, _ = file.GetCellValue(tabName, "D7")
	toSetting.timezone, _ = file.GetCellValue(tabName, "D8")

	dataRange, _ := file.GetCellValue(tabName, "D9")
	if i, err := strconv.Atoi(dataRange); err != nil {
		panic("datarange 轉換失敗:" + err.Error())
	} else {
		toSetting.dataRange = i
	}

	fmt.Println(fromSetting)
	fmt.Println(toSetting)
}

func main() {

	fromURL := fromSetting.user + ":" + fromSetting.pwd + "@tcp(" + fromSetting.host + ":" + fromSetting.port + ")/" + fromSetting.name + "?charset=" + fromSetting.charset + "&loc=" + url.QueryEscape(fromSetting.timezone)
	toURL := toSetting.user + ":" + toSetting.pwd + "@tcp(" + toSetting.host + ":" + toSetting.port + ")/" + toSetting.name + "?charset=" + toSetting.charset + "&loc=" + url.QueryEscape(toSetting.timezone)
	orm.RegisterDataBase(fromSetting.alias, "mysql", fromURL, 50)
	orm.RegisterDataBase(toSetting.alias, "mysql", toURL, 50)

	orm.RegisterModel(new(iom5BasicMold), new(iom5PostProductRecord), new(iom5BasicDefect), new(iom5BasicDefectRel), new(iom5InfoSchedule), new(iom5BasicMachine), new(WorkShop), new(Iom5AuthMailAccount), new(SystemSetting))
	orm.RunSyncdb(toSetting.alias, false, false)

	var err error

	err = allActions()
	if err != nil {
		panic(err)
	}
}

func allActions() (err error) {

	err = actionsForProductRecord()
	if err != nil {
		fmt.Println("ProductRecord: ", err)
	} else {
		fmt.Println("Product Record Copy success")
	}

	err = actionsForMold()
	if err != nil {
		fmt.Println("Mold:", err)
	} else {
		fmt.Println("Mold Copy success")
	}

	err = actionsForProductNGType()
	if err != nil {
		fmt.Println("ProductNGType:", err)
	} else {
		fmt.Println("ProductNGType Copy success")
	}

	err = actionsForProductRecordNGInfo()
	if err != nil {
		fmt.Println("ProductNGInfo:", err)
	} else {
		fmt.Println("ProductNGInfo Copy success")
	}

	err = actionsForSchedule()
	if err != nil {
		fmt.Println("Schedule:", err)
	} else {
		fmt.Println("Schedule Copy success")
	}

	err = actionsForMachine()
	if err != nil {
		fmt.Println("Machine:", err)
	} else {
		fmt.Println("Machine Copy success")
	}

	err = insertFirstWorkShopData()
	if err != nil {
		fmt.Println("WorkShop:", err)
	} else {
		fmt.Println("WorkShop create success")
	}

	err = insertAndUpdateEmailData()
	if err != nil {
		fmt.Println("EmailAccount:", err)
	} else {
		fmt.Println("EmailAccount create success")
	}

	err = actionsForSystemSetting()
	if err != nil {
		fmt.Println("SystemSetting:", err)
	} else {
		fmt.Println("SystemSetting create success")
	}

	return nil
}

// ------------------------------------------------ Actions

func actionsForMold() (err error) {
	var arr []iom5MoldOld
	var newArr []iom5BasicMold

	if arr, err = getMoldData(); err != nil {
		return err
	}

	for _, v := range arr {
		greenRange, err := strconv.ParseFloat(v.GreenRange, 10)
		if err != nil {
			panic(err)
		}

		yellowRange, err := strconv.ParseFloat(v.YellowRange, 10)
		if err != nil {
			panic(err)
		}

		planCycleTime, err := strconv.ParseFloat(v.PlanCycleTime, 10)
		if err != nil {
			panic(err)
		}

		newArr = append(newArr, iom5BasicMold{v.ID, v.Name, v.ProductModel, v.Number, 0, 0, 0, 0, 0, 0, 0, v.PlasticType, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, greenRange, yellowRange, planCycleTime, 0, 0, 0, 0})
	}

	err = insertMoldDataToDB(newArr)
	if err != nil {
		return err
	}
	return
}

func actionsForProductRecord() (err error) {
	var arr []iom5ProductRecordOld
	var newArr []iom5PostProductRecord

	if arr, err = getProductRecordData(); err != nil {
		return err
	}

	for _, v := range arr {
		newArr = append(newArr, iom5PostProductRecord{v.ID, v.SerialNumber, 0, v.UpdateTime, 0, v.StartTime, v.EndTime, v.NormalNumber, v.AbnormalNumber})
	}

	err = insertProductRecordDataToDB(newArr)
	if err != nil {
		return err
	}
	return nil
}

func actionsForProductNGType() (err error) {
	var arr []iom5ProductNgType
	var newArr []iom5BasicDefect

	if arr, err = getProductNGTypeData(); err != nil {
		return err
	}

	for _, v := range arr {
		newArr = append(newArr, iom5BasicDefect{v.ID, v.Name, "", v.Delete, 0})
	}

	err = insertProductNGTypeToDB(newArr) //塞資料
	if err != nil {
		return err
	}
	return nil
}

func actionsForProductRecordNGInfo() (err error) {
	var arr []iom5ProductRecordNgInfo
	var newArr []iom5BasicDefectRel

	if arr, err = getProductRecordNGInfoData(); err != nil {
		return err
	}

	for _, v := range arr {
		newArr = append(newArr, iom5BasicDefectRel{v.ID, v.NGTypeID, v.Number, v.ProductID})
	}

	err = insertProductNGTypeInfoToDB(newArr)
	if err != nil {
		return err
	}
	return nil
}

func actionsForSchedule() (err error) {
	var arr []iom5Schedule
	var newArr []iom5InfoSchedule
	timeForm := "2006-01-02 15:04"

	if arr, err = getScheduleData(); err != nil {
		return err
	}

	for _, v := range arr {
		var upDateBool bool

		number, err := strconv.ParseInt(v.Qty, 10, 64)
		if err != nil {
			panic(err)
		}

		greenRange, err := strconv.ParseFloat(v.GreenRange, 10)
		if err != nil {
			panic(err)
		}

		yellowRange, err := strconv.ParseFloat(v.YellowRange, 10)
		if err != nil {
			panic(err)
		}

		loc, _ := time.LoadLocation("Local")
		theStartTime, _ := time.ParseInLocation(timeForm, v.TimeStartinput, loc)
		startTime := theStartTime.UnixNano() / 1000000
		theEndTime, _ := time.ParseInLocation(timeForm, v.TimeEndinput, loc)
		endTime := theEndTime.UnixNano() / 1000000
		now := time.Now()

		if v.Update == 1 {
			upDateBool = true
		} else {
			upDateBool = false
		}
		newArr = append(newArr, iom5InfoSchedule{v.ID, v.ScheduleSerial, v.MoldID, greenRange, yellowRange, 0, v.MachineID, 0, startTime, endTime, 0, upDateBool, number, 0, now.UnixNano() / 1000000})
	}

	err = insertScheduleDataToDB(newArr)
	if err != nil {
		return err
	}
	return nil
}

func actionsForMachine() (err error) {
	var arr []machineOld
	var newArr []iom5BasicMachine
	timeForm := "2006-01-02 15:04:05"

	if arr, err = getMachineData(); err != nil {
		return err
	}

	for _, v := range arr {
		loc, _ := time.LoadLocation("Local")
		theModifiedTime, _ := time.ParseInLocation(timeForm, v.Modified, loc)
		modifiedTime := theModifiedTime.UnixNano() / 1000000

		newArr = append(newArr, iom5BasicMachine{v.ID, v.Brand, v.MachineName, v.MachineNumber, v.MachineModel, 0, 0, 0, 0, float64(v.Shot), 0, float64(v.MaxPressure), 0, 0, 0, 0, 0, 0, 0, float64(v.LockPower), 0, 0, 0, 0, 0, 0, modifiedTime, 0, 0, false, 0, 0})
	}

	err = insertMachineDataToDB(newArr)
	if err != nil {
		return err
	}
	return nil
}

func actionsForSystemSetting() (err error) {
	var arr []SystemSetting

	if arr, err = getSystemSettingData(); err != nil {
		return err
	}

	err = insertSystemSettingData(arr)
	if err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------ ProductRecordNGTypeInfo COPY
func insertProductNGTypeInfoToDB(arr []iom5BasicDefectRel) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func getProductRecordNGInfoData() (arr []iom5ProductRecordNgInfo, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllProductRecordNGInfoFromDB(o)
	return
}

//	------------------------------------------------WorkShop insert
func insertFirstWorkShopData() (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	var ws WorkShop
	ws.Name = "BDI"
	ws.Number = "BDI"
	ws.Remark = "BDI"
	ws.Type = 0

	_, err = o.Insert(&ws)
	return
}

// ------------------------------------------------ ProductNGType Copy

func insertProductNGTypeToDB(arr []iom5BasicDefect) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getProductNGTypeData() (arr []iom5ProductNgType, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllProductNGTypeDataFromDB(o)
	return
}

// ------------------------------------------------ ProductRecord COPY

func insertProductRecordDataToDB(arr []iom5PostProductRecord) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
			fmt.Println("haha")
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getProductRecordData() (arr []iom5ProductRecordOld, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllProductRecordDataFromDB(o)
	return
}

// ------------------------------------------------MOLD COPY

func insertMoldDataToDB(arr []iom5BasicMold) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getMoldData() (arr []iom5MoldOld, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllMoldDataFromDB(o)
	return
}

// ------------------------------------------------Schedule COPY

func insertScheduleDataToDB(arr []iom5InfoSchedule) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getScheduleData() (arr []iom5Schedule, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllScheduleFromDB(o)
	return
}

// ------------------------------------------------Schedule COPY

func insertMachineDataToDB(arr []iom5BasicMachine) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getMachineData() (arr []machineOld, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)
	arr, err = getAllMachineFromDB(o)
	return
}

//	----------------------------------------------- Email copy

func insertAndUpdateEmailData() (err error) {
	var arr []Iom5AuthMailAccount
	o := orm.NewOrm()
	arr, err = getEmailDataFromOld()
	if err != nil {
		panic(err)
	}

	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	oldArr, err := getMailAccountFromNewDB(o)
	if err != nil {
		panic(err)
	}
	if !isMailAccountEmpty() {
		arr[0].ID = oldArr[0].ID
		arr[1].ID = oldArr[1].ID
		arr[2].ID = oldArr[2].ID
		for i := 0; i < len(arr); i++ {
			updateMailAccount(arr[i])
		}
	} else {
		o.InsertMulti(len(arr), arr)
	}
	return
}

func updateMailAccount(mail Iom5AuthMailAccount) (err error) {
	o := orm.NewOrm()
	if _, err := updateNewMailAccountTableFromDB(mail, o); err != nil {
		panic(err)
	}
	return nil
}

func isMailAccountEmpty() bool {
	o := orm.NewOrm()

	count, err := countMailAccountFromNewDB(o)
	if err != nil {
		panic(err)
	}
	return count == 0
}

func getEmailDataFromOld() (mailAcount []Iom5AuthMailAccount, err error) {
	o := orm.NewOrm()
	o.Using(fromSetting.alias)

	var mailAccount1 = new(Iom5AuthMailAccount)
	var mailAccount2 = new(Iom5AuthMailAccount)
	var mailAccount3 = new(Iom5AuthMailAccount)

	stringMap, err := getSystemSettingMapFromNewDB(o)
	if err != nil {
		panic(err)
	}

	var settingOne []string
	var settingOneInt []int64
	var s1 string
	_ = json.Unmarshal([]byte(stringMap["daily_group_setting_one"].(string)), &settingOne)
	if len(settingOne) != 0 {
		for _, v := range settingOne {
			j, err := strconv.Atoi(v)
			if err == nil {
				settingOneInt = append(settingOneInt, int64(j)) //將string字串改成int arr
			}
		}
		byteArrayOne, _ := json.Marshal(settingOneInt) //將int arr 轉成byteArr
		s1 = string([]byte(byteArrayOne))
	} else {
		s1 = ""
	}

	var settingTwo []string
	var settingTwoInt []int64
	var s2 string
	_ = json.Unmarshal([]byte(stringMap["daily_group_setting_two"].(string)), &settingTwo)
	if len(settingTwo) != 0 {
		for _, v := range settingTwo {
			j, err := strconv.Atoi(v)
			if err == nil {
				settingTwoInt = append(settingTwoInt, int64(j)) //將string字串改成int arr
			}
		}
		byteArrayTwo, _ := json.Marshal(settingTwoInt) //將int arr 轉成byteArr
		s2 = string([]byte(byteArrayTwo))

	} else {
		s2 = ""
	}

	var settingThree []string
	var settingThreeInt []int64
	var s3 string
	_ = json.Unmarshal([]byte(stringMap["daily_group_setting_three"].(string)), &settingThree)
	if len(settingThree) != 0 {
		for _, v := range settingThree {
			j, err := strconv.Atoi(v)
			if err == nil {
				settingThreeInt = append(settingThreeInt, int64(j)) //將string字串改成int arr
			}
		}
		byteArrayThree, _ := json.Marshal(settingThreeInt) //將int arr 轉成byteArr
		s3 = string([]byte(byteArrayThree))
	} else {
		s3 = ""
	}

	mailAccount1.Email = stringMap["daily_recipient_mail_one"].(string)
	mailAccount1.Recipient = stringMap["daily_recipient_name_one"].(string)
	mailAccount1.WorkShopSetting = s1
	mailAccount2.Email = stringMap["daily_recipient_mail_two"].(string)
	mailAccount2.Recipient = stringMap["daily_recipient_name_two"].(string)
	mailAccount2.WorkShopSetting = s2
	mailAccount3.Email = stringMap["daily_recipient_mail_three"].(string)
	mailAccount3.Recipient = stringMap["daily_recipient_name_three"].(string)
	mailAccount3.WorkShopSetting = s3

	mailAcount = append(mailAcount, *mailAccount1, *mailAccount2, *mailAccount3)

	return
}

//	----------------------------------------------- system setting copy

func insertSystemSettingData(arr []SystemSetting) (err error) {
	o := orm.NewOrm()
	o.Using(toSetting.alias)
	o.Begin()
	defer func() {
		if r := recover(); r != nil {
			o.Rollback()
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("Unknown panic")
			}
		} else {
			o.Commit()
		}
	}()

	if len(arr) != 0 {
		_, err = o.InsertMulti(toSetting.dataRange, arr)
		if err != nil {
			panic(err)
		}
		return
	}

	return nil
}

func getSystemSettingData() (arr []SystemSetting, err error) {
	// var arr []SystemSetting
	o := orm.NewOrm()
	o.Using(fromSetting.alias)

	stringMap, err := getSystemSettingMapFromNewDB(o)
	if err != nil {
		panic(err)
	}

	stringMap["mail_host"] = stringMap["host"]
	stringMap["mail_password"] = stringMap["password"]
	stringMap["mail_port"] = stringMap["port"]
	stringMap["mail_user"] = stringMap["user_name"]

	var slashArr []int

	s := stringMap["web_daily_report_url"].(string)
	if strings.HasPrefix(s, "http") || strings.HasPrefix(s, "https") {
		for i := 0; i < len(s)-1; i++ {
			if string(s[i]) == "/" {
				slashArr = append(slashArr, i)
			}
		}
		s = s[slashArr[1]+1 : slashArr[2]]
	} else {
		s = ""
	}

	arr = append(arr, SystemSetting{"carousel_frequency", stringMap["carousel_frequency"].(string)})
	arr = append(arr, SystemSetting{"mail_host", stringMap["mail_host"].(string)})
	arr = append(arr, SystemSetting{"mail_password", stringMap["mail_password"].(string)})
	arr = append(arr, SystemSetting{"mail_port", stringMap["mail_port"].(string)})
	arr = append(arr, SystemSetting{"mail_user", stringMap["mail_user"].(string)})
	arr = append(arr, SystemSetting{"report_red_boundary", stringMap["report_red_boundary"].(string)})
	arr = append(arr, SystemSetting{"report_yellow_boundary", stringMap["report_yellow_boundary"].(string)})
	arr = append(arr, SystemSetting{"web_daily_report_url", stringMap["web_daily_report_url"].(string)})
	arr = append(arr, SystemSetting{"web_url", s})
	return
}

// -----------------------------------------------DAO

func getAllProductRecordDataFromDB(o orm.Ormer) (arr []iom5ProductRecordOld, err error) { //ProductRecordData
	sql := "SELECT id, serial_number, start_time,end_time,normal_number,abnormal_number,update_time FROM edge.edge_40_product_record"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getAllMoldDataFromDB(o orm.Ormer) (arr []iom5MoldOld, err error) { //Mold
	sql := "select * from edge.edge_40_mold"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getAllProductNGTypeDataFromDB(o orm.Ormer) (arr []iom5ProductNgType, err error) { //ProductNGType
	sql := "SELECT * FROM edge.edge_40_product_ng_type;"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getAllProductRecordNGInfoFromDB(o orm.Ormer) (arr []iom5ProductRecordNgInfo, err error) { //ProductRecordNGInfo
	sql := "SELECT * FROM edge.edge_40_product_record_ng_info;"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getAllScheduleFromDB(o orm.Ormer) (arr []iom5Schedule, err error) { //Schedule
	sql := "SELECT * FROM edge.edge_40_schedule;"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getAllMachineFromDB(o orm.Ormer) (arr []machineOld, err error) { //Schedule
	sql := "SELECT * FROM edge.edge_40_machine;"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func getSystemSettingFromOldDB(o orm.Ormer) (arr []MailAccountOld, err error) {

	sql := "select * from edge.edge_40_system_setting"
	_, err = o.Raw(sql).QueryRows(&arr)
	return
}

func countMailAccountFromNewDB(o orm.Ormer) (int, error) {
	var count int
	o.Using(toSetting.alias)
	sql := "SELECT count(*) FROM iom5_auth_mail_account;"
	err := o.Raw(sql).QueryRow(&count)
	return count, err
}

func getMailAccountFromNewDB(o orm.Ormer) (arr []Iom5AuthMailAccount, err error) {

	sql := "SELECT * FROM iom5_auth_mail_account;"
	_, err = o.Raw(sql).QueryRows(&arr)
	return arr, err
}

func getSystemSettingMapFromNewDB(o orm.Ormer) (ormMap orm.Params, err error) {
	o.Using(toSetting.alias)
	ormMap = make(orm.Params)
	sql := "select parameter, value from edge.edge_40_system_setting"
	_, err = o.Raw(sql).RowsToMap(&ormMap, "parameter", "value")

	return
}

func updateNewMailAccountTableFromDB(mail Iom5AuthMailAccount, o orm.Ormer) (int64, error) {
	o.Using(toSetting.alias)
	affectedNum, err := o.Update(&mail)
	return affectedNum, err

}
