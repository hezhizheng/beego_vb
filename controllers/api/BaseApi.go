package api

// 成功状态码
const Success int = 0

type BaseStatus struct {
    Code int `json:"code"`
    Data interface{} `json:"data"`
    Error string `json:"error"`
}

func Response(data interface{}, code int , error string) BaseStatus {
    return BaseStatus{
        code,
        data,
        error,
    }
}


