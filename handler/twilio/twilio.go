package twilio

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	twilio "github.com/twilio/twilio-go"

	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"net/http"
)

const (
	accountSid = "SK2e136bbed5d27be5e6a2bc9b98113ffb"
	authToken  = "tnJZxVAOlZG3pFEmfHNYHXtq3X0ZCa1d"
	fromNumber = "+18006924636" // Twilio 虚拟号码
)

// accountSid := "SK2e136bbed5d27be5e6a2bc9b98113ffb"
// apiKey := "Taxi Virtual Number Proxy"
// apiSecret := "tnJZxVAOlZG3pFEmfHNYHXtq3X0ZCa1d"
//
//	client := twilio.NewRestClientWithParams(twilio.ClientParams{
//		Username:   apiKey,
//		Password:   apiSecret,
//		AccountSid: accountSid,
//	})
var client *twilio.RestClient

func TwilioSendMsg(from, to string) {
	//accountSid := "AC7f9abdea97d3280e8d3dae797d3ab16e"
	//authToken := "2f6a1dc05fc644ba203e82181c12d303"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetBody("Hello from Go!")

	resp, err := client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}
}

func TwilioCall(from, to string) {
	// Twilio 账户信息
	//accountSid := "AC7f9abdea97d3280e8d3dae797d3ab16e"
	//authToken := "2f6a1dc05fc644ba203e82181c12d303"
	accountSid := "SK2e136bbed5d27be5e6a2bc9b98113ffb"
	authToken := "tnJZxVAOlZG3pFEmfHNYHXtq3X0ZCa1d"

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	// 设置拨打电话的参数
	params := &twilioApi.CreateCallParams{}
	params.SetTo(to)
	params.SetFrom(from)
	params.SetTwiml("<Response><Say voice='alice'>你好，这是一个 Twilio 语音测试。</Say></Response>")

	// 拨打电话
	resp, err := client.Api.CreateCall(params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Call SID:", *resp.Sid)
}
func ResponseHandle(c *gin.Context) {
	speechResult := c.DefaultPostForm("SpeechResult", "") // 获取语音识别结果
	c.String(http.StatusOK, "<Response><Say>您说了："+speechResult+"。</Say></Response>")
}
func Call(c *gin.Context) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	var req VoiceCallRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少 to 参数"})
		return
	}
	//	twimlContent := `
	//<Response>
	//    <Say>您好，请问你的。</Say>
	//    <Gather input="speech" timeout="5" action="/process_speech" method="POST">
	//        <Say>请讲。</Say>
	//    </Gather>
	//    <Say>没有听到您的语音，正在结束通话。</Say>
	//</Response>`
	params := &twilioApi.CreateCallParams{}
	params.SetTo(req.To)
	params.SetFrom(fromNumber)
	params.SetUrl("https://handler.twilio.com/twiml/EH9ff52536f0d3842ab126a4073040440c")
	//params.SetTwiml(twimlContent)

	resp, err := client.Api.CreateCall(params)
	if err != nil {
		log.Println("拨号失败:", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sid":     *resp.Sid,
		"allresp": resp,
	})
}

// 请求结构体
type VoiceCallRequest struct {
	To string `json:"to" binding:"required"`
}
