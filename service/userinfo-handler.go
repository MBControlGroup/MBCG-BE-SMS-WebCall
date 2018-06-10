package service

import (
	"io/ioutil"
	"strings"
	"encoding/base64"
	"encoding/json"
	//"fmt"
    "net/http"
    //"strconv"
    "fmt"
    "time"
    "crypto/md5"
    "github.com/MBControlGroup/MBCG-BE-SMS/entities"
    //"github.com/MBControlGroup/MBCG-BE-SMS/token"
    "github.com/unrolled/render"
    //"github.com/dgrijalva/jwt-go/request"
    //"github.com/dgrijalva/jwt-go"

)

const (
    smsAccount = "N00000015110"
    webcallAccount = "N00000016155"
    smsSecret = "637aac20-c5f6-11e7-91b8-7fa2dd1dcc08"
    webcallSecret = "929d29f0-ba21-11e7-a1ba-fb60feab0763"
    host = "http://apis.7moor.com"
)

func getDateTime() string {
    return time.Now().Format("20060102150405")
}

func mdfive(text string) string {
    data := []byte(text)
    has := md5.Sum(data)
    md5str := fmt.Sprintf("%x",has)
    return strings.ToUpper(md5str)
}

func baseSixFour(text string) string {
    b := []byte(text)
    return base64.URLEncoding.EncodeToString(b)
}

func webCallCallbackHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        fmt.Println(req.Form["actionid"][0])
        formatter.JSON(w, http.StatusOK, struct{ActionId string; Message string}{req.Form["actionid"][0], req.Form["Message"][0]})
    }
}

func webCallHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        var webCallInfo entities.WebCallInfo
        err := json.NewDecoder(req.Body).Decode(&webCallInfo)
        checkErr(err)

        tt := getDateTime()
        sig := mdfive(webcallAccount+webcallSecret+tt)
        fmt.Println(sig)
        //sendInterfaceTemplateSms getSmsTemplate
        interfacePath := "/v20160818/webCall/webCall/"
        url := host+interfacePath+webcallAccount+"?sig="+sig;
        auth := baseSixFour(webcallAccount+":"+tt)
        
        client := &http.Client{}
        
        /*myjson := struct{
            Action string `json:"Action"`
            ServiceNo string `json:"ServiceNo"`
            Exten string `json:"Exten"`
            WebCallType string `json:"WebCallType"`
            CallBackUrl string `json:"CallBackUrl"`
            Variable string `json:"Variable"`
        }{
            "Webcall",
            "02033275113",
            "13719327791",
            "asynchronous",
            "http://172.17.0.1:8080/webCall/callback",
            "role:1",
        }*/

        ttjson, err := json.Marshal(webCallInfo)

        request, err := http.NewRequest("POST",url, strings.NewReader(string(ttjson)))
        checkErr(err)

        request.Header.Add("Accept", "application/json")
        request.Header.Add("Content-Type","application/json;charset=utf-8")
        request.Header.Add("Authorization",auth)
        
        response, err := client.Do(request)
        checkErr(err)

        defer response.Body.Close()

        body, err := ioutil.ReadAll(response.Body)

        fmt.Println(string(body))

        var webCallRes entities.WebCallRes
        err = json.Unmarshal(body, &webCallRes)
        checkErr(err)
        formatter.JSON(w, http.StatusOK, webCallRes)
    }
} 

func getSmsTempHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        tt := getDateTime()
        sig := mdfive(smsAccount+smsSecret+tt)
        fmt.Println(sig)
        //sendInterfaceTemplateSms getSmsTemplate
        interfacePath := "/v20160818/sms/getSmsTemplate/"
        url := host+interfacePath+smsAccount+"?sig="+sig;
        auth := baseSixFour(smsAccount+":"+tt)
        
        client := &http.Client{}

        request, err := http.NewRequest("POST",url, nil/*strings.NewReader(string(ttjson))*/)
        checkErr(err)

        request.Header.Add("Accept", "application/json")
        request.Header.Add("Content-Type","application/json;charset=utf-8")
        request.Header.Add("Authorization",auth)
        
        response, err := client.Do(request)
        checkErr(err)

        defer response.Body.Close()

        body, err := ioutil.ReadAll(response.Body)

        fmt.Println(string(body))

        var messageTemplateRes entities.MessageTemplateRes
        err = json.Unmarshal(body, &messageTemplateRes)
        checkErr(err)
        formatter.JSON(w, http.StatusOK, messageTemplateRes)
    }
} 

func sendTempSmsHandler(formatter *render.Render) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()

        if len(req.Form["vars"]) == 0 {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}; Message string}{400, "fail", "失败", nil, "without vars!"})
            return
        }

        var reqBody interface{}

        if req.Form["vars"][0] == "2" {
            var var2 entities.Vars2Template
            err := json.NewDecoder(req.Body).Decode(&var2)
            checkErr(err)
            reqBody = var2
        } else if req.Form["vars"][0] == "3" {
            var var3 entities.Vars3Template
            err := json.NewDecoder(req.Body).Decode(&var3)
            checkErr(err)
            reqBody = var3
        } else if req.Form["vars"][0] == "4" {
            var var4 entities.Vars4Template
            err := json.NewDecoder(req.Body).Decode(&var4)
            checkErr(err)
            reqBody = var4
        } else {
            formatter.JSON(w, http.StatusBadRequest, struct{ Code int;Enmsg string;Cnmsg string; Data interface{}; Message string}{400, "fail", "失败", nil, "Bad vars!"})
            return
        }

        tt := getDateTime()
        sig := mdfive(smsAccount+smsSecret+tt)
        fmt.Println(sig)
        //sendInterfaceTemplateSms getSmsTemplate
        interfacePath := "/v20160818/sms/sendInterfaceTemplateSms/"
        url := host+interfacePath+smsAccount+"?sig="+sig;
        auth := baseSixFour(smsAccount+":"+tt)
        
        client := &http.Client{}
        
        /*myjson := struct{
            Num string `json:"num"`
            TemplateNum string `json:"templateNum"`
            Var1 string `json:"var1"`
            Var2 string `json:"var2"`,
        }{
            "13711112396",
            "1",
            "1234",
            "5678",
        }*/

        ttjson, err := json.Marshal(reqBody)

        request, err := http.NewRequest("POST",url, strings.NewReader(string(ttjson)))
        checkErr(err)

        request.Header.Add("Accept", "application/json")
        request.Header.Add("Content-Type","application/json;charset=utf-8")
        request.Header.Add("Authorization",auth)
        
        response, err := client.Do(request)
        checkErr(err)

        defer response.Body.Close()

        body, err := ioutil.ReadAll(response.Body)

        fmt.Println(string(body))

        var shortMessageSentRes entities.ShortMessageSentRes
        err = json.Unmarshal(body, &shortMessageSentRes)
        checkErr(err)
        formatter.JSON(w, http.StatusOK, shortMessageSentRes)
    }
} 