package entities

const (
    OF = iota
    OR
)

type Vars struct {
    Vars        int     `json:"vars"`
}

type Vars2Template struct {
    Num         string `json:"num"`
    TemplateNum string `json:"templateNum"`
    Var1        string `json:"var1"`
    Var2        string `json:"var2"`
}

type Vars3Template struct {
    Num         string `json:"num"`
    TemplateNum string `json:"templateNum"`
    Var1        string `json:"var1"`
    Var2        string `json:"var2"`
    Var3        string `json:"var3"`
}

type Vars4Template struct {
    Num         string `json:"num"`
    TemplateNum string `json:"templateNum"`
    Var1        string `json:"var1"`
    Var2        string `json:"var2"`
    Var3        string `json:"var3"`
    Var4        string `json:"var4"`
}

type MessageTemplate struct {
    Id          string  `json:"_id"`
    Name        string  `json:"name"`
    Content     string  `json:"content"`
    Num         string  `json:"num"`
    Vars        int  `json:"vars"`
    Sign        string  `json:"sign"`
}

type MessageTemplateRes struct {
    Success     bool  `json:"success"`
    Data        []MessageTemplate  `json:"data"`
}

type ShortMessageSentRes struct {
    Success     bool    `json:"success"`
    Flag        string  `json:"flag"`
    Msgid       string  `json:"msgid"`
    Message     string  `json:"message"`
}

type WebCallRes struct {
    Command     string  `json:"Command"`
    Succeed     bool    `json:"Succeed"`
    ActionID    string  `json:"ActionID"`
    Response    string  `json:"Response"`
}

type WebCallInfo struct{
    Action      string `json:"Action"`
    ServiceNo   string `json:"ServiceNo"`
    Exten       string `json:"Exten"`
    WebCallType string `json:"WebCallType"`
    CallBackUrl string `json:"CallBackUrl"`
    Variable    string `json:"Variable"`
}



