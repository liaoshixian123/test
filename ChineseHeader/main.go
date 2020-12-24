package main

import (
	"fmt"
	"net/url"
)

// 通過http需要傳中文的header時， 後端這裡無法直接接到中文的header因為編碼的問題
// 為了解決這個問題 前端需要先進行轉碼 後端接取之後則在解碼則可正常使用

// 以下为vue中对中文进行编码解码的方式：

//  编码
// encodeURIComponent(str)
//  解码
// decodeURIComponent(str)

// 后台进行接收：
// 编码
// java.net.URLEncoder.encode(token,"UTF-8")
// 解码
// java.net.URLDecoder.decode(token,"UTF-8")

// 此部分為Golang解碼的method
func main() {

	encodedHeader := "%E4%B8%AD%E6%96%87" // 這裡是前端將轉完碼的參數透過header的形式來傳遞

	if tmpStr, err := url.QueryUnescape(encodedHeader); err != nil {
		fmt.Println("Err", err)
	} else {
		fmt.Println(tmpStr)
	}

}
