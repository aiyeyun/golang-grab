package ssc

import (
	"fmt"
	"xmn/core/model"
	"strings"
	"time"
	"log"
	"xmn/core/logger"
	"xmn/core/mail"
	"strconv"
)

//数据包
var contain_datapackage []*model.Packet

//重庆开奖数据
var contain_cq_data []*model.Cqssc

//重庆开奖数据
var contain_tj_data []*model.Tjssc

//重庆开奖数据
var contain_xj_data []*model.Xjssc

//重庆开奖数据
var contain_tw_data []*model.Twssc

//彩票类型
var contain_ssc_type map[int]string

//重庆时时彩 最新开奖号
var cq_ssc_new_code string

//天津时时彩 最新开奖号
var tj_ssc_new_code string

//新疆时时彩 最新开奖号
var xj_ssc_new_code string

//台湾五分彩 最新开奖号
var tw_ssc_new_code string


//时时彩
//包含数据包 算法
func Contain()  {
	fmt.Println("时时彩 - 包含数据包 算法")

	contain_ssc_type = make(map[int]string)
	contain_ssc_type[1] = "重庆时时彩"
	contain_ssc_type[2] = "天津时时彩"
	contain_ssc_type[3] = "新疆时时彩"
	contain_ssc_type[4] = "台湾五分彩"

	packet := new(model.Packet)
	contain_datapackage = packet.Query()

	cqssc := new(model.Cqssc)
	contain_cq_data = cqssc.Query()

	tjssc := new(model.Tjssc)
	contain_tj_data = tjssc.Query()

	xjssc := new(model.Xjssc)
	contain_xj_data = xjssc.Query()

	twssc := new(model.Twssc)
	contain_tw_data = twssc.Query()

	containAnalysis()
}

func containAnalysis()  {
	for i := range contain_datapackage {
		go containAnalysisCodes(contain_datapackage[i])
	}
}

func containAnalysisCodes(packet *model.Packet)  {
	log.Println("时时彩－包含数据包 正在分析　数据包别名:", packet.Alias, "彩种:", contain_ssc_type[packet.Type])
	slice_dataTxt := strings.Split(packet.DataTxt, "\r\n")
	//slice data txt to slice data txt map
	dataTxtMap := make(map[string]string)
	for i := range slice_dataTxt {
		dataTxtMap[slice_dataTxt[i]] = slice_dataTxt[i]
	}

	//fmt.Println(dataTxtMap)

	//检查是否在报警时间段以内
	if (packet.Start >0 && packet.End >0) && (time.Now().Hour() < packet.Start || time.Now().Hour() > packet.End)  {
		log.Println("彩票类型:", contain_ssc_type[packet.Type], "数据包别名:", packet.Alias, "报警通知非接受时间段内")
		logger.Log("彩票类型: " +  contain_ssc_type[packet.Type] + " 数据包别名: " + packet.Alias + " 报警通知非接受时间段内 ")
		return
	}

	//开奖数据
	codes := make([]string, 0)
	//重庆时时彩
	if packet.Type == 1 {
		//检查 该彩种到最新的一起 是否重复分析
		if len(contain_cq_data) > 0 {
			new_code := contain_cq_data[len(contain_cq_data) - 1].One + contain_cq_data[len(contain_cq_data) - 1].Two + contain_cq_data[len(contain_cq_data) - 1].Three + contain_cq_data[len(contain_cq_data) - 1].Four + contain_cq_data[len(contain_cq_data) - 1].Five
			if new_code == cq_ssc_new_code {
				log.Println(contain_ssc_type[packet.Type], "最新的一期 已经分析过了... 等待出现新的开奖号")
				return
			}
		}

		for i := range contain_cq_data {
			code := contain_cq_data[i].One + contain_cq_data[i].Two + contain_cq_data[i].Three + contain_cq_data[i].Four +contain_cq_data[i].Five
			codes = append(codes, code)
			//刷新该彩种的最新开奖号
			cq_ssc_new_code = code
		}
	}
	//天津时时彩
	if packet.Type == 2 {
		//检查 该彩种到最新的一起 是否重复分析
		if len(contain_tj_data) > 0 {
			new_code := contain_tj_data[len(contain_tj_data) - 1].One + contain_tj_data[len(contain_tj_data) - 1].Two + contain_tj_data[len(contain_tj_data) - 1].Three + contain_tj_data[len(contain_tj_data) - 1].Four + contain_tj_data[len(contain_tj_data) - 1].Five
			if new_code == tj_ssc_new_code {
				log.Println(contain_ssc_type[packet.Type], "最新的一期 已经分析过了... 等待出现新的开奖号")
				return
			}
		}

		for i := range contain_tj_data {
			code := contain_tj_data[i].One + contain_tj_data[i].Two + contain_tj_data[i].Three + contain_tj_data[i].Four +contain_tj_data[i].Five
			codes = append(codes, code)
			//刷新该彩种的最新开奖号
			tj_ssc_new_code = code
		}
	}
	//新疆时时彩
	if packet.Type == 3 {
		//检查 该彩种到最新的一起 是否重复分析
		if len(contain_xj_data) > 0 {
			new_code := contain_xj_data[len(contain_xj_data) - 1].One + contain_xj_data[len(contain_xj_data) - 1].Two + contain_xj_data[len(contain_xj_data) - 1].Three + contain_xj_data[len(contain_xj_data) - 1].Four + contain_xj_data[len(contain_xj_data) - 1].Five
			if new_code == xj_ssc_new_code {
				log.Println(contain_ssc_type[packet.Type], "最新的一期 已经分析过了... 等待出现新的开奖号")
				return
			}
		}

		for i := range contain_xj_data {
			code := contain_xj_data[i].One + contain_xj_data[i].Two + contain_xj_data[i].Three + contain_xj_data[i].Four +contain_xj_data[i].Five
			codes = append(codes, code)
			//刷新该彩种的最新开奖号
			xj_ssc_new_code = code
		}
	}
	//台湾时时彩
	if packet.Type == 4 {
		//检查 该彩种到最新的一起 是否重复分析
		if len(contain_tw_data) > 0 {
			new_code := contain_tw_data[len(contain_tw_data) - 1].One + contain_tw_data[len(contain_tw_data) - 1].Two + contain_tw_data[len(contain_tw_data) - 1].Three + contain_tw_data[len(contain_tw_data) - 1].Four + contain_tw_data[len(contain_tw_data) - 1].Five
			if new_code == tw_ssc_new_code {
				log.Println(contain_ssc_type[packet.Type], "最新的一期 已经分析过了... 等待出现新的开奖号")
				return
			}
		}

		for i := range contain_tw_data {
			code := contain_tw_data[i].One + contain_tw_data[i].Two + contain_tw_data[i].Three + contain_tw_data[i].Four +contain_tw_data[i].Five
			codes = append(codes, code)
			//刷新该彩种的最新开奖号
			tw_ssc_new_code = code
		}
	}

	//fmt.Println(contain_ssc_type[packet.Type])
	//fmt.Println(codes)

	//各单位报警期数 初始化
	var q3_number int = 0
	var z3_number int = 0
	var h3_number int = 0

	//各单位报警 是否有上期参考对象 初始化
	var q3_reference bool = false
	var z3_reference bool = false
	var h3_reference bool = false
	for i := range codes{
		code_byte := []byte(codes[i])
		//前三号码
		q3 := string(code_byte[0]) + string(code_byte[1]) + string(code_byte[2])
		//中三号码
		z3 := string(code_byte[1]) + string(code_byte[2]) + string(code_byte[3])
		//后三号码
		h3 := string(code_byte[2]) + string(code_byte[3]) + string(code_byte[4])

		//各单位是否在 数据包内 初始化
		var q3_in bool = false
		var z3_in bool = false
		var h3_in bool = false

		//前三号码 是否在数据包内
		if _, ok := dataTxtMap[q3]; ok {
			q3_in = true
		}
		//中三号码 是否在数据包内
		if _, ok := dataTxtMap[z3]; ok {
			z3_in = true
		}
		//后三号码 是否在数据包内
		if _, ok := dataTxtMap[h3]; ok {
			h3_in = true
		}

		//前三没有上一期 开奖数据 参考对象 and 前三出现在数据包里
		if !q3_reference && q3_in {
			q3_number = q3_number + 1
			//fmt.Println(contain_ssc_type[packet.Type], "q3", q3, "+1=", q3_number)
		} else if q3_reference && q3_in  {
			//前三有上一期 开奖数据 参考对象 and 前三出现在数据包里
			q3_number = 0
			q3_number = q3_number + 1
			//fmt.Println(contain_ssc_type[packet.Type], "q3", q3, "清0 +1=", q3_number)
		}

		//中三没有上一期 开奖数据 参考对象 and 中三出现在数据包里
		if !z3_reference && z3_in {
			z3_number = z3_number + 1
		} else if z3_reference && z3_in  {
			//中三有上一期 开奖数据 参考对象 and 中三出现在数据包里
			z3_number = 0
			z3_number = z3_number + 1
		}

		//后三没有上一期 开奖数据 参考对象 and 后三出现在数据包里
		if !h3_reference && h3_in {
			h3_number = h3_number + 1
		} else if h3_reference && h3_in  {
			//后三有上一期 开奖数据 参考对象 and 后三出现在数据包里
			h3_number = 0
			h3_number = h3_number + 1
		}

		//前三参考对象
		if q3_in {
			q3_reference = true
		} else {
			q3_reference = false
		}

		//中三参考对象
		if z3_in {
			z3_reference = true
		} else {
			z3_reference = false
		}

		//后三参考对象
		if h3_in {
			h3_reference = true
		} else {
			h3_reference = false
		}
	}

	//最新的一期有数据包里的数据 才报警
	if !q3_reference {
		q3_number = 0
	}
	if !z3_reference {
		z3_number = 0
	}
	if !h3_reference {
		h3_number = 0
	}

	//fmt.Println(contain_ssc_type[packet.Type], "q3 期数", q3_number)
	//fmt.Println(contain_ssc_type[packet.Type], "z3 期数", z3_number)
	//fmt.Println(contain_ssc_type[packet.Type], "h3 期数", h3_number)

	var body string = ""

	//前三报警
	if q3_number >= packet.RegretNumber {
		body += "<div>数据包别名: " + packet.Alias + " 位置 前三 " + strconv.Itoa(q3_number) + " 期 报警！</div><br/>"
	}

	//中三报警
	if z3_number >= packet.RegretNumber {
		body += "<div>数据包别名: " + packet.Alias + " 位置 中三 " + strconv.Itoa(q3_number) + " 期 报警！</div><br/>"
	}

	//后三报警
	if h3_number >= packet.RegretNumber {
		body += "<div>数据包别名: " + packet.Alias + " 位置 后三 " + strconv.Itoa(q3_number) + " 期 报警！</div><br/>"
	}

	//发送邮件
	if body != "" {
		go mail.SendMail(contain_ssc_type[packet.Type] + " 包含数据包", body)
	}
}