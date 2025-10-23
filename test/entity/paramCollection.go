package ystruct2

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
)

const (
	UserJSON     = `{"id": 1, "name": "John Doe", "email": "john.doe@example.com"}`
	ProductJSON  = `{"id": 101, "name": "Widget", "price": 19.99}`
	OrderJSON    = `{"id": 201, "userId": 1, "productIds": [101, 102], "total": 39.98}`
	TimeLineJSON = "{\n" +
		"    \"data\": {\n" +
		"        \"data\": \"{\\\"tenantId\\\":\\\"dcccvmvmrsptq3acr8hs\\\"," +
		"\\\"resourceId\\\":\\\"1879108022489600001\\\",\\\"bizCode\\\":\\\"COMMUNITY_MONITORING_PROTOCOL\\\"," +
		"\\\"bizData\\\":{\\\"requestId\\\":\\\"1430257704504393728176112699378622178\\\"," +
		"\\\"sessionId\\\":\\\"9876543210\\\",\\\"type\\\":\\\"time_line\\\"," +
		"\\\"timestamp\\\":1761128298741,\\\"from\\\":\\\"1879108022489600001\\\"," +
		"\\\"to\\\":\\\"6cc8de3d20fe8c391aj3wk\\\",\\\"contentType\\\":\\\"sd_playback\\\",\\\"content\\\":{}," +
		"\\\"realDev\\\":false,\\\"startTime\\\":1761116313000,\\\"endTime\\\":1761116329000,\\\"timezoneId\\\":\\\"Asia/Shanghai\\\"},\\\"deliveryType\\\":\\\"device\\\",\\\"deliveryTargetId\\\":\\\"1430257704504393728\\\"}\",\n" +
		"        \"protocol\": 700102,\n" +
		"        \"pv\": \"1.0\",\n" +
		"        \"sign\": \"397643435\",\n" +
		"        \"t\": 1761128298872,\n" +
		"        \"businessId\": \"dcccvmvmrsptq3acr8hs\",\n" +
		"        \"channel\": null,\n" +
		"        \"tenantInfo\": null\n" +
		"    },\n" +
		"    \"count\": 0,\n" +
		"    \"ct\": 0,\n" +
		"    \"rt\": 0,\n" +
		"    \"groupId\": null,\n" +
		"    \"offsets\": null\n" +
		"}"
)

func GetTimeLineStr() string {
	// timeLineJSONStr, err := json.Marshal(TimeLineJSON)
	// if err != nil {
	// 	log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	// }
	// return string(timeLineJSONStr)
	return strings.Replace(TimeLineJSON, "1430257704504393728176112699378622178", GetSnowFlakId(), -1)

	//return TimeLineJSON
}

func GetSnowFlakId() string {
	// 初始化雪花节点
	node, err := snowflake.NewNode(1) // 1 是机器 ID
	if err != nil {
		fmt.Println("Error creating snowflake node:", err)
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return node.Generate().String()
}

func GetInitiativeMessageMsg() string {
	msg := "{\"data\":{\"data\":\"{\\\"partitionKey\\\":\\\"bf282f5350df04a24cnnhv\\\",\\\"requestId\\\":\\\"d07957b9-36a6-4e29-a751-310c571783cd\\\",\\\"bizCode\\\":\\\"dpReport\\\",\\\"deviceId\\\":\\\"bf282f5350df04a24cnnhv\\\",\\\"productId\\\":\\\"ax83yfv609mkiabw\\\",\\\"category\\\":\\\"cdsxj\\\",\\\"categoryCode\\\":\\\"DAYJ6P2A35YWR725WH\\\",\\\"time\\\":1744266099418,\\\"uuid\\\":\\\"uuid1471926be280991b\\\",\\\"spaceId\\\":\\\"216340247\\\",\\\"statusList\\\":[{\\\"id\\\":212,\\\"code\\\":\\\"initiative_message\\\",\\\"value\\\":\\\"eyJ2IjoiNS4wIiwiY2FwdHVyZUlkIjogIjE3NDQyNjU4NzIxNjMwMCIsImNvc3RfdGltZSI6IjEwMjkiLCJjbWQiOiJpcGNfY3VzdG9tXzEiLCJ0eXBlIjoiaW1hZ2UiLCJ3aXRoIjoicmVzb3VyY2UiLCJhbGFybSI6dHJ1ZSwidGltZSI6MTc0NDI2NjEwMCwiZmlsZXMiOltbInR5LWV1LXN0b3JhZ2UzMCIsIi83MjQxNTUtMjE2MzQwMjQ3LXV1aWQxNDcxOTI2YmUyODA5OTFiL3VuaWZ5LzE3NDQyNjYxMDAuanBlZyIsImEwYjJiYzBjNzRjYjRmZDYiLCIxNzQ2ODU4MTAxIl1dfQ==\\\",\\\"reportTime\\\":1744266099050}]}\",\"protocol\":50402,\"pv\":\"1.0\",\"sign\":\"-117422073\",\"t\":1744266099418,\"businessId\":\"8qga8euft3rdgtakyfg3\",\"channel\":null,\"tenantInfo\":\"{\\\"mgmtTenantId\\\":\\\"222786871153922874\\\",\\\"region\\\":\\\"eu\\\",\\\"tenantId\\\":\\\"8qga8euft3rdgtakyfg3\\\"}\"},\"count\":0,\"ct\":0,\"rt\":0,\"groupId\":null,\"offsets\":null}"
	return msg
}

func GetImgRetryMsg() string {
	msg := "{\"dpRecordMsgBody\":{\"bizCode\":\"dpReport\",\"category\":\"cdsxj\",\"categoryCode\":\"DAYJ6P2A35YWR725WH\",\"deviceId\":\"6cfc1657e31bc5caddqaig\",\"productId\":\"ax83yfv609mkiabw\",\"spaceId\":\"154107613\",\"statusList\":[{\"code\":\"initiative_message\",\"id\":212,\"reportTime\":1740556970620,\"value\":\"eyJ2IjoiNS4wIiwic25hcF9pZCI6ICJTTkFQMDAwMTIwMTUwMjEwMDAyMyIsImNvc3RfdGltZSI6IjQ5NCIsImNtZCI6ImlwY19jdXN0b21fNiIsInR5cGUiOiJpbWFnZSIsIndpdGgiOiJyZXNvdXJjZSIsImFsYXJtIjp0cnVlLCJ0aW1lIjoxNzQwNTU2OTcxLCJmaWxlcyI6W1sidHktY24tc3RvcmFnZTMwLTEyNTQxNTM5MDEiLCIvNGNkMjVkLTE1NDEwNzYxMy11dWlkZDUwNWI3NDU4NzQ3MzAwZi91bmlmeS8xNzQwNTU2OTcxLmpwZWciLCIyYmQ3NGQzYWMwNTA0NmIwIiwiMTc0MzE0ODk3MSJdXX0=\"}],\"tenantId\":\"dcccvmvmrsptq3acr8hs\",\"time\":\"1740556970869\",\"uuid\":\"uuidd505b7458747300f\"},\"dpStatus\":{\"$ref\":\"$.dpRecordMsgBody.statusList[0]\"},\"originalMsg\":true,\"retryCount\":2,\"uid\":\"bay1724834401259Cea3\"}"
	return msg
}

func GetMonitorOfferMsg() string {
	msg := "{\"businessId\":\"dcccvmvmrsptq3acr8hs\"," +
		"\"data\":\"{\\\"content\\\":{\\\"sdp\\\":\\\"v=0\\\\no=- 4611733160928125423 2 IN IP4\\\"},\\\"contentType\\\":\\\"live\\\",\\\"from\\\":\\\"uid-1234565\\\",\\\"needReminder\\\":false,\\\"requestId\\\":\\\"1234567890\\\",\\\"sessionId\\\":\\\"1338454802485481472173918160648379883\\\",\\\"timestamp\\\":1739159828000,\\\"to\\\":\\\"6ca338ae37e4bddc44xm1i\\\",\\\"type\\\":\\\"offer\\\"}\",\"protocol\":1,\"pv\":\"pv\",\"t\":1739522749049}"
	return msg
}

func GetMonitorOfferMsg2() string {
	msg := "{\n" +
		"    \"count\": 0,\n" +
		"    \"ct\": 0,\n" +
		"    \"data\": {\n" +
		"        \"businessId\": \"tenantId-test-123\",\n" +
		"        \"data\": \"{\\\"bizCode\\\":\\\"web_rtc\\\",\\\"data\\\":{\\\"from\\\":\\\"devId123456\\\",\\\"needReminder\\\":false,\\\"requestId\\\":\\\"reqId123456\\\",\\\"timestamp\\\":0,\\\"to\\\":\\\"uid123456\\\"},\\\"needSlice\\\":false,\\\"requestId\\\":\\\"reqId123456\\\",\\\"t\\\":1743477908424}\",\n" +
		"        \"protocol\": 600102,\n" +
		"        \"pv\": \"1.0\",\n" +
		"        \"t\": 1743477908703\n" +
		"    },\n" +
		"    \"rt\": 0\n" +
		"}"
	return msg
}

func GetAiNews() string {
	msg := "{\n" +
		"    \"audio_url\": \"\",\n" +
		"    \"author\": \"Bangkok Post Public Company Limited\",\n" +
		"    \"category\": \"ความคิดเห็น\",\n" +
		"    \"channel\": \"tuya\",\n" +
		"    \"content\": \"ภาพถ่ายเมื่อวันที่ 9 ตุลาคม 2554 แสดงให้เห็นสถานการณ์น้ำท่วมรุนแรงในจังหวัดอยุธยาในปี 2554 แหล่งที่มา: Bangkok Post/ประกิจ จันทวงษ์\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\t\\\\n\\\\n\\\\nประเทศไทยเผชิญกับปัญหาน้ำท่วมมานานกว่าศตวรรษ บางปีสถานการณ์เลวร้ายกว่าปีอื่น ๆ แต่รูปแบบยังคงเหมือนเดิม เหตุการณ์น้ำท่วมครั้งใหญ่ในปี 2554 ยังคงเป็นสิ่งที่เจ็บปวดที่สุด: ธนาคารโลกระบุว่าสร้างความเสียหายทางเศรษฐกิจ 46.5 พันล้านดอลลาร์สหรัฐฯ (1.5 ล้านล้านบาท) ทำให้ผู้คน 13 ล้านคนต้องพลัดถิ่น และคร่าชีวิตผู้คนไปประมาณ 800 ราย ศูนย์กลางอุตสาหกรรมของประเทศส่วนใหญ่จมอยู่ใต้น้ำเป็นเวลาหลายเดือน ส่งผลกระทบอย่างรุนแรงต่อห่วงโซ่อุปทานทั่วโลก ข้ามมาถึงปี 2568 สถานการณ์ก็ยังคงไม่เปลี่ยนแปลงมากนัก น้ำท่วมครั้งล่าสุดในน่าน (กรกฎาคม) เชียงใหม่ (สิงหาคม) และเพชรบูรณ์ (กันยายน) เน้นย้ำว่าน้ำท่วมที่สร้างความเสียหายยังคงเกิดขึ้นอย่างต่อเนื่อง.\\\\nน้ำท่วมไม่ได้จำกัดอยู่แค่จังหวัดในชนบท กรุงเทพฯ เองก็เคยประสบกับน้ำท่วมรุนแรงในปี 2481, 2538 และ 2554.\\\\nน้ำท่วมครั้งใหญ่ยังเกิดขึ้นในหลายจังหวัดในปี 2526, 2543, 2553, 2556, 2557, 2559, 2562, 2567 และ 2568 ความถี่ของเหตุการณ์เหล่านี้ตอกย้ำว่าปัญหานี้ฝังรากลึกในชีวิตของคนไทยเพียงใด.\\\\nทำไมน้ำท่วมจึงเกิดขึ้นบ่อยครั้ง\\\\nน้ำท่วมในประเทศไทยเป็นผลมาจากทั้งสภาพธรรมชาติและความเปราะบางที่มนุษย์สร้างขึ้น:\\\\nภูมิประเทศ: ประเทศไทยตั้งอยู่บนที่ราบลุ่มปากแม่น้ำที่มีแม่น้ำสายหลักไหลจากเหนือลงใต้ ทำให้มีแนวโน้มที่จะเกิดน้ำท่วมตามธรรมชาติ\\\\nสภาพอากาศ: ฤดูมรสุมและพายุโซนร้อนนำมาซึ่งปริมาณน้ำฝนที่ตกหนัก บางครั้งรุนแรงขึ้นจากพายุไต้ฝุ่นจากมหาสมุทรแปซิฟิก.\\\\nการทรุดตัวของแผ่นดิน: การทรุดตัวที่ไม่สามารถย้อนกลับได้ของที่ราบลุ่มปากแม่น้ำทำให้ปัญหารุนแรงขึ้น โดยเฉพาะในกรุงเทพฯ และจังหวัดใกล้เคียง.\\\\nการขยายตัวของเมืองและการใช้ประโยชน์ที่ดิน: การขยายตัวของเมืองอย่างรวดเร็วและขาดการวางแผนที่ดีได้ปิดกั้นเส้นทางระบายน้ำตามธรรมชาติ ถมพื้นที่ชุ่มน้ำ และเพิ่มปริมาณน้ำไหลบ่า.\\\\nการตัดไม้ทำลายป่า: การลดลงของป่าไม้ลดความสามารถในการดูดซับน้ำตามธรรมชาติและเร่งการกัดเซาะของดิน.\\\\nการจัดการทางน้ำที่ไม่ดี: คลองและแม่น้ำมักจะอุดตันด้วยขยะหรือบำรุงรักษาไม่เพียงพอ.\\\\nโครงสร้างพื้นฐานที่ไม่เพียงพอ: คันกั้นน้ำและเขื่อนที่มีอยู่มีความเสี่ยงที่จะพังทลาย ในขณะที่ระบบป้องกันน้ำท่วมขนาดใหญ่ยังไม่ตอบสนองต่อความต้องการ.\\\\nความท้าทายด้านธรรมาภิบาล: การทับซ้อนของความรับผิดชอบระหว่างหน่วยงาน การประสานงานที่ไม่ดี และการตัดสินใจแบบรวมศูนย์ขัดขวางการจัดการที่มีประสิทธิภาพ.\\\\nการเปลี่ยนแปลงสภาพภูมิอากาศ: ระดับน้ำทะเลที่สูงขึ้น อุณหภูมิที่ร้อนขึ้น และพายุที่รุนแรงขึ้นทำให้ความเปราะบางที่มีอยู่แย่ลง.\\\\n2554 'เป็นสิ่งเตือนใจ'\\\\nน่าสนใจว่าหลายเดือนก่อนน้ำท่วมครั้งใหญ่ในปี 2554 ทีมผู้เชี่ยวชาญชาวดัตช์ได้เตือนว่าประเทศไทยมีความเสี่ยงสูงที่จะเกิดน้ำท่วมร้ายแรงภายในไม่กี่ปีข้างหน้า การคาดการณ์ของพวกเขาพิสูจน์ให้เห็นว่าแม่นยำอย่างน่าตกใจเมื่อประเทศถูกน้ำท่วมในอีกสี่เดือนต่อมา ชาวดัตช์ซึ่งมาจากประเทศที่ประมาณ 26% ของพื้นที่อยู่ต่ำกว่าระดับน้ำทะเล และประมาณ 59% ของประเทศมีความเสี่ยงต่อน้ำท่วม เป็นผู้นำระดับโลกในการจัดการน้ำท่วมมาอย่างยาวนาน คำเตือนของพวกเขาเน้นย้ำถึงความจำเป็นที่ประเทศไทยจะต้องนำกลยุทธ์ที่ครอบคลุมและระยะยาวมาใช้ แทนที่จะเป็นการตอบสนองแบบเร่งด่วน.\\\\nแม้จะมีคำมั่นสัญญาในการปฏิรูปและแผนป้องกันน้ำท่วมที่ทะเยอทะยานหลังปี 2554 แต่ความคืบหน้าก็เป็นไปอย่างช้าๆ โครงการขนาดใหญ่มักจะหยุดชะงักในขั้นวางแผน โดยมีอุปสรรคจากค่าใช้จ่าย การเมือง หรือการขาดฉันทามติ ข้อเสนอโครงการขนาดใหญ่ปี 2554 ของอดีตรัฐมนตรีว่าการกระทรวงวิทยาศาสตร์และเทคโนโลยี ปลอดประสพ สุรัสวดี ก็เป็นตัวอย่างหนึ่ง ซึ่งยิ่งใหญ่ในวิสัยทัศน์แต่สุดท้ายก็ไม่สำเร็จ.\\\\nทัศนคติต่อปัญหาน้ำท่วม\\\\nแม้ว่าน้ำท่วมจะถูกมองว่าเป็นภัยพิบัติ แต่บางชุมชนก็ยอมรับว่าเป็นส่วนหนึ่งของชีวิต ดังที่ชาวบ้านคนหนึ่งกล่าวว่า: \\\\\\\"เมื่อน้ำท่วมมา เราก็ยกเฟอร์นิเจอร์ขึ้นชั้นบน รอให้น้ำลด แล้วก็ทำความสะอาดและใช้ชีวิตตามปกติ.\\\\\\\"\\\\nบางคนยังเห็นประโยชน์ เกษตรกรและชาวประมงในจังหวัดริมแม่น้ำโขงและสาขาต่างๆ สังเกตว่าน้ำท่วมตามฤดูกาลช่วยเติมเต็มแหล่งประมงและเสริมสร้างความอุดมสมบูรณ์ของที่ดินเพื่อการเกษตร การยอมรับน้ำท่วมในวัฒนธรรมนี้ทำให้รัฐบาลสร้างฉันทามติเกี่ยวกับมาตรการป้องกันที่มีค่าใช้จ่ายสูงได้ยากขึ้น.\\\\nอย่างไรก็ตาม ความจริงทางเศรษฐกิจนั้นยากที่จะมองข้าม นอกเหนือจากความเสียหายโดยตรงแล้ว น้ำท่วมที่เกิดขึ้นซ้ำๆ ยังบั่นทอนความเชื่อมั่นของนักลงทุน สร้างความเสียหายต่อโครงสร้างพื้นฐาน และคุกคามแหล่งมรดกทางวัฒนธรรม ยูเนสโกได้แสดงความกังวลซ้ำแล้วซ้ำเล่าเกี่ยวกับความเสี่ยงจากน้ำท่วมต่อวัด วัง และโบราณสถานของไทย รวมถึงในอยุธยา.\\\\nข้อบกพร่องทางนโยบาย\\\\nหนึ่งในอุปสรรคที่ใหญ่ที่สุดในการแก้ไขปัญหาน้ำท่วมของประเทศไทยคือธรรมาภิบาล แม้จะมีหลายหน่วยงานที่รับผิดชอบการจัดการน้ำและภัยพิบัติ เช่น กรมชลประทาน (RID), กรมทรัพยากรน้ำ (DWR) และกรมป้องกันและบรรเทาสาธารณภัย (DDPM) แต่การประสานงานกลับไม่ดี ความรับผิดชอบทับซ้อนกัน และความเป็นผู้นำขาดการรวมศูนย์.\\\\nแม้ว่าจะมีการจัดตั้งหน่วยงานใหม่ เช่น สำนักงานทรัพยากรน้ำแห่งชาติ (OWRM) และศูนย์จัดการภัยพิบัติแห่งชาติ (NDMC) เพื่อเสริมสร้างการประสานงาน แต่การแยกส่วนยังคงเป็นจุดอ่อนที่สำคัญ รัฐบาลมักจะเลือกใช้นโยบายแบบบนลงล่าง ซึ่งมักจะละเลยความรู้ท้องถิ่น นักการเมืองมักจะให้ความสำคัญกับนโยบายระยะสั้นที่ให้ผลลัพธ์ภายในช่วงเวลาการเลือกตั้งของตน มากกว่าที่จะมุ่งมั่นในกลยุทธ์หลายทศวรรษที่จำเป็นสำหรับการป้องกันน้ำท่วม.\\\\nบทเรียนจากต่างประเทศ\\\\nประเทศไทยไม่ได้เผชิญกับความท้าทายจากน้ำท่วมเพียงลำพัง เนเธอร์แลนด์ เยอรมนี และประเทศอื่นๆ ได้พัฒนากลยุทธ์การจัดการน้ำแบบบูรณาการระยะยาวที่ผสมผสานวิศวกรรมเข้ากับการมีส่วนร่วมของชุมชน.\\\\nตัวอย่างที่น่าสนใจคือสหภาพเมืองแม่น้ำไรน์ ซึ่งรัฐบาลท้องถิ่นริมแม่น้ำร่วมมือกันทุกปีเพื่อแบ่งปันประสบการณ์ ประสานงานโครงการป้องกันน้ำท่วม และรวบรวมข้อมูลจากสาธารณะ แนวทางนี้เน้นย้ำถึงความร่วมมือ การสร้างขีดความสามารถ และวิสัยทัศน์ระยะยาว ซึ่งเป็นหลักการที่ประเทศไทยสามารถนำมาปรับใช้ได้ ประสบการณ์จากนานาชาติแสดงให้เห็นว่าการจัดการน้ำท่วมไม่สามารถพึ่งพาวิศวกรรมเพียงอย่างเดียวได้ แนวทางแก้ไขที่มีประสิทธิภาพจะต้องรวมถึงการปฏิรูปสถาบัน การมีส่วนร่วมของผู้มีส่วนได้ส่วนเสีย การให้ความรู้แก่สาธารณะ และการวางแผนการปรับตัวต่อสภาพภูมิอากาศ.\\\\ nแนวทางแก้ไขที่เป็นไปได้\\\\nการทำลายวงจรของน้ำท่วมที่เกิดขึ้นซ้ำๆ จะต้องอาศัยการแทรกแซงที่กล้าหาญและระยะยาว ลำดับความสำคัญหลัก ได้แก่:\\\\nการจัดการน้ำแบบบูรณาการ: การประสานงานที่ดีขึ้นระหว่างหน่วยงานและระดับการปกครอง โดยมีผู้รับผิดชอบที่ชัดเจน.\\\\nการลงทุนในโครงสร้างพื้นฐาน: การสร้างและบำรุงรักษาคันกั้นน้ำ อ่างเก็บน้ำ ระบบระบายน้ำ และทางระบายน้ำ โดยได้รับการสนับสนุนจากเทคโนโลยีที่ทันสมัย.\\\\nการปฏิรูปการใช้ที่ดิน: การบังคับใช้กฎหมายผังเมือง การปกป้องพื้นที่ชุ่มน้ำ และการป้องกันการก่อสร้างที่มีความเสี่ยงตามแนวทางน้ำ.\\\\nการมีส่วนร่วมของชุมชน: การดึงชุมชนท้องถิ่นเข้ามามีส่วนร่วมในการวางแผน การเตรียมพร้อม และการบำรุงรักษา.\\\\nการปรับตัวต่อสภาพภูมิอากาศ: การออกแบบนโยบายและโครงสร้างพื้นฐานโดยคำนึงถึงระดับน้ำทะเลที่สูงขึ้นและพายุที่รุนแรงขึ้น.\\\\nกลไกทางการเงิน: การแนะนำเครื่องมือต่างๆ เช่น ภาษีน้ำหรือโครงการประกันภัย เพื่อให้แน่ใจว่ามีเงินทุนที่ยั่งยืนสำหรับโครงการระยะยาว.\\\\nการเปลี่ยนแปลงเหล่านี้ไม่ใช่เรื่องง่าย ต้องใช้ทรัพยากรทางการเงิน เจตจำนงทางการเมือง และการเปลี่ยนแปลงทัศนคติของประชาชน แต่หากขาดการสนับสนุนนโยบาย การดำเนินการ และความมุ่งมั่น ประเทศไทยก็เสี่ยงที่จะกลับไปเผชิญกับฝันร้ายในปี 2554 ซ้ำอีก หรือแย่กว่านั้นมาก.\\\\nGeorge G van der Meulen, Ph.D, เป็นอาจารย์ประจำ Asian Institute of Technology (AIT) และกรรมการผู้จัดการของ Compuplan Knowledge Institute of applied geo-Spatial Informatics (CKI) Chamniern Paul Vorratnchaiphan, Ph.D, เป็นประธานของ Institute for Research and Social Action Foundation.\",\n" +
		"    \"description\": \"ประเทศไทยเผชิญกับปัญหาน้ำท่วมมานานกว่าศตวรรษ บางปีสถานการณ์เลวร้ายกว่าปีอื่น ๆ แต่รูปแบบยังคงเหมือนเดิม เหตุการณ์น้ำท่วมครั้งใหญ่ในปี 2554 ยังคงเป็นสิ่งที่เจ็บปวดที่สุด: ธนาคารโลกระบุว่าสร้างความเสียหายทางเศรษฐกิจ 46.5 พันล้านดอลลาร์สหรัฐฯ (1.5 ล้านล้านบาท) ทำให้ผู้คน 13 ล้านคนต้องพลัดถิ่น และคร่าชีวิตผู้คนไปประมาณ 800 ราย ศูนย์กลางอุตสาหกรรมของประเทศส่วนใหญ่จมอยู่ใต้น้ำเป็นเวลาหลายเดือน ส่งผลกระทบอย่างรุนแรงต่อห่วงโซ่อุปทานทั่วโลก\",\n" +
		"    \"keywords\": [\n" +
		"        \"น้ำท่วมประเทศไทย\",\n" +
		"        \"การจัดการน้ำท่วม\",\n" +
		"        \"ผลกระทบทางเศรษฐกิจ\",\n" +
		"        \"การปรับตัวต่อการเปลี่ยนแปลงสภาพภูมิอากาศ\",\n" +
		"        \"ทรัพยากรน้ำ\",\n" +
		"        \"การพัฒนาโครงสร้างพื้นฐาน\",\n" +
		"        \"การขยายตัวของเมือง\",\n" +
		"        \"การเตรียมพร้อมรับมือภัยพิบัติ\",\n" +
		"        \"ธรรมาภิบาล\",\n" +
		"        \"ธนาคารโลก\"\n" +
		"    ],\n" +
		"    \"language\": \"th\",\n" +
		"    \"location\": [\n" +
		"        \"ประเทศไทย\",\n" +
		"        \"อยุธยา\",\n" +
		"        \"ประเทศไทย, กรุงเทพฯ\",\n" +
		"        \"น่าน\",\n" +
		"        \"เชียงใหม่\",\n" +
		"        \"เพชรบูรณ์\",\n" +
		"        \"กรุงเทพฯ\",\n" +
		"        \"เนเธอร์แลนด์\",\n" +
		"        \"เยอรมนี\",\n" +
		"        \"กรุงเทพฯ\",\n" +
		"        \"แม่น้ำไรน์\"\n" +
		"    ],\n" +
		"    \"news_id\": \"7232ea4520265f6bef7e20edfb09e7c1c83586e8b17457b093e2072700e6ec66\",\n" +
		"    \"poster_url\": \"https://static.bangkokpost.com/media/content/dcx/2025/09/16/5780842_700.jpg\",\n" +
		"    \"published_at\": 1757959260000,\n" +
		"    \"recorded_at\": 1757980865351,\n" +
		"    \"source\": \"Bangkok Post\",\n" +
		"    \"source_url\": \"https://www.bangkokpost.com/opinion/opinion/3105186/floods-a-recurring-nightmare\",\n" +
		"    \"speech\": \"ข่าวนี้กล่าวถึงปัญหาอุทกภัยที่ประเทศไทยต้องเผชิญมานานกว่าศตวรรษ โดยเรียกมันว่าเป็น 'ฝันร้ายที่กลับมาหลอกหลอนซ้ำแล้วซ้ำเล่า' น้ำท่วมครั้งใหญ่ในปี 2554 สร้างความเสียหายทางเศรษฐกิจถึง 46.5 พันล้านดอลลาร์สหรัฐฯ ทำให้ผู้คน 13 ล้านคนต้องพลัดถิ่น และส่งผลกระทบอย่างรุนแรงต่อห่วงโซ่อุปทานทั่วโลก แม้จะมีคำเตือนจากผู้เชี่ยวชาญชาวดัตช์และคำมั่นสัญญาในการปฏิรูป แต่ความคืบหน้าในการจัดการน้ำท่วมยังคงเป็นไปอย่างช้าๆ เนื่องจากปัจจัยซับซ้อนหลายประการ เช่น สภาพภูมิประเทศ ฤดูมรสุมที่รุนแรง การทรุดตัวของแผ่นดิน การขยายตัวของเมืองที่ขาดการวางแผน การทำลายป่า และโครงสร้างพื้นฐานและการกำกับดูแลที่ไม่เพียงพอ บทความนี้เน้นย้ำว่าแม้บางชุมชนจะยอมรับน้ำท่วมตามฤดูกาล แต่ผลกระทบทางเศรษฐกิจอย่างมหาศาลต่อความเชื่อมั่นของนักลงทุนและมรดกทางวัฒนธรรมนั้นปฏิเสธไม่ได้ มีการเสนอแนวทางแก้ไขที่ครอบคลุม เช่น การจัดการน้ำแบบบูรณาการ การลงทุนในโครงสร้างพื้นฐาน การปฏิรูปการใช้ที่ดิน การมีส่วนร่วมของชุมชน และการปรับตัวต่อสภาพภูมิอากาศ โดยเรียนรู้จากประเทศอย่างเนเธอร์แลนด์และเยอรมนี ซึ่งเน้นย้ำถึงความจำเป็นเร่งด่วนสำหรับกลยุทธ์ระยะยาวและเจตจำนงทางการเมืองเพื่อหลีกเลี่ยงหายนะในอนาคต\",\n" +
		"    \"subtitle\": \"\",\n" +
		"    \"tags\": [\n" +
		"        \"น้ำท่วม\",\n" +
		"        \"การเปลี่ยนแปลงสภาพภูมิอากาศ\",\n" +
		"        \"โครงสร้างพื้นฐาน\",\n" +
		"        \"การลงทุน\",\n" +
		"        \"ประเด็นร้อนวันนี้\",\n" +
		"        \"การจัดการขยะ\"\n" +
		"    ],\n" +
		"    \"tenantId\": \"ay_h98yfqy7xmgs3gtcwn3j\",\n" +
		"    \"title\": \"น้ำท่วม 'ฝันร้ายที่กลับมาหลอกหลอนซ้ำแล้วซ้ำเล่า'\"\n" +
		"}"
	return msg
}
func GetAiNews2() string {
	msg := "{\n" +
		"    \"title\": \"MEA Cracks Down on Illegal Cable Cutting\",\n" +
		"    \"subtitle\": \"The Metropolitan Electricity Authority warns offenders that cable theft, electricity tampering, and meter fraud are life-threatening crimes that will lead to arrest and prosecution.\",\n" +
		"    \"published_at\": 1758266700000,\n" +
		"    \"recorded_at\": 1758268075982,\n" +
		"    \"author\": \"Bangkok Post Public Company Limited\",\n" +
		"    \"source\": \"Bangkok Post\",\n" +
		"    \"source_url\": \"https://www.bangkokpost.com/business/general/3107341/mea-cracks-down-on-illegal-cable-cutting\",\n" +
		"    \"keywords\": [\n" +
		"      \"MEA\",\n" +
		"      \"electricity theft\",\n" +
		"      \"cable cutting\",\n" +
		"      \"meter tampering\",\n" +
		"      \"Bitcoin mining\",\n" +
		"      \"public safety\",\n" +
		"      \"power distribution\"\n" +
		"    ],\n" +
		"    \"category\": \"BUSINESS\",\n" +
		"    \"description\": \"The Metropolitan Electricity Authority (MEA) has issued a warning against illegal activities such as cutting power cables, using stolen electricity for Bitcoin mining, and tampering with electricity meters, all of which are considered MEA property.\",\n" +
		"    \"content\": \"The Metropolitan Electricity Authority (MEA) has issued a warning against illegal activities such as cutting power cables, using stolen electricity for Bitcoin mining, and tampering with electricity meters, all of which are considered MEA property.Recently, MEA’s Yannawa District Office, in collaboration with local police, arrested a suspect attempting to cut underground power cables near Narathiwas Soi 11. MEA emphasises that it takes these matters seriously, deploying teams to regularly inspect power lines, working closely with local police stations, and pursuing legal action against offenders.\\nSuch actions violate both civil and criminal laws and can cause severe damage — not only endangering the perpetrator’s life but also putting the public at risk. These illegal acts may lead to fatal accidents such as electrocution or fires, disrupt the power distribution system, cause widespread blackouts, and impact thousands of electricity users, in addition to damaging MEA-owned public infrastructure.\\nMEA urges the public to remain vigilant and report any suspicious activities, such as illegal cable cutting, electricity theft, or unsafe electrical equipment. Reports can be made via the MEA Smart Life Application, available free on iOS and Android, or through MEA’s official social media channels, including Line: MEA Connect (@MEAthailand) (with the green shield badge), by selecting “Contact MEA Call Centre Online.” Reports may also be submitted through the MEA Call Centre 1130, available 24 hours a day.\",\n" +
		"    \"poster_url\": \"https://static.bangkokpost.com/media/content/20250919/c1_3107341_700.png\",\n" +
		"    \"location\": [\n" +
		"      \"Bangkok, Thailand\"\n" +
		"    ],\n" +
		"    \"language\": \"en\",\n" +
		"    \"tags\": [\n" +
		"      \"energy\",\n" +
		"      \"infrastructure\",\n" +
		"      \"today's hot topics\"\n" +
		"    ],\n" +
		"    \"speech\": \"The Metropolitan Electricity Authority (MEA) is cracking down on illegal activities like cable cutting, electricity theft for Bitcoin mining, and meter tampering. These actions, which are crimes against MEA property, can lead to severe consequences including electrocution, fires, widespread blackouts, and damage to public infrastructure, endangering both perpetrators and the public. MEA, in collaboration with local police, regularly inspects power lines and pursues legal action against offenders, as demonstrated by a recent arrest in Yannawa. The authority urges the public to report any suspicious activities through the MEA Smart Life Application, social media, or their 24-hour call center at 1130 to ensure public safety and maintain power distribution integrity.\",\n" +
		"    \"audio_url\": \"\"\n" +
		"  }"
	return msg
}
