package login

import (
	"net/http"
	"time"

	"github.com/imroc/req/v3"
)

type LoginResult struct {
	Client    *req.Client
	Status    int
	Location  string
	SetCookie []*http.Cookie
	Body      string
}

func LoginMelonWithReq(email, pwd string, proxyURL string) (*LoginResult, error) {
	// 1) 构建 req 客户端：开启 Cookie、HTTP/2、浏览器指纹、代理、不跟随重定向
	c := req.C().
		// 自动管理 Cookie（服务端返回 Set-Cookie 会自动存进 Jar）
		// 文档：Cookie 自动管理。:contentReference[oaicite:0]{index=0}
		// 默认就启用 Jar，这里显式留着注释
		// SetCookieJar(http.DefaultClient.Jar).
		// 强制 HTTP/2（如后端不支持会报错；你也可以去掉这行走 ALPN 自动协商）

		// 使用浏览器 HTTP 指纹（包含 TLS 指纹 + HTTP/2 伪首部/优先级/头顺序等）
		// 文档：HTTP Fingerprint（ImpersonateChrome/Firefox/Edge）。:contentReference[oaicite:1]{index=1}
		ImpersonateChrome().
		// 不自动跟随跳转，保留 302 以读取 Set-Cookie/Location
		// 文档：Redirect Policy（NoRedirectPolicy）。:contentReference[oaicite:2]{index=2}
		SetRedirectPolicy(req.NoRedirectPolicy()).
		// 超时
		SetTimeout(15 * time.Second)

	// 2) 代理（支持 http/socks5，也支持函数）
	// 文档：Proxy。:contentReference[oaicite:3]{index=3}
	if proxyURL != "" {
		c.SetProxyURL(proxyURL)
	}

	// 3) 拼表单（Content-Type 会自动设为 application/x-www-form-urlencoded）
	// 文档：Set Form Data。:contentReference[oaicite:4]{index=4}

	// 4) 构建请求（把你 curl 里的头都塞进来；UA 建议与 Impersonate 的浏览器一致）
	// 注意：req 会处理 HTTP/2 伪首部和头顺序，不必强行控制大小写
	r := c.ImpersonateChrome().R().
		// 这些是请求级 Cookie（与 curl -b 等价）；如果你已有这三枚就加上，
		// 没有也可以不加，通常登录页面会下发。
		SetCookies(
			&http.Cookie{Name: "PCID", Value: "17320157945852523050993", Domain: ".melon.com"},
			&http.Cookie{Name: "TKT_POC_ID", Value: "WP19", Domain: ".melon.com"},
			&http.Cookie{Name: "JSESSIONID", Value: "E874130B959DC968E6C7EFCF2425A99F", Domain: "gmember.melon.com"},
		).
		SetHeaders(map[string]string{
			"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
			"accept-language":           "zh-CN,zh;q=0.9",
			"cache-control":             "max-age=0",
			"content-type":              "application/x-www-form-urlencoded",
			"origin":                    "https://gmember.melon.com",
			"priority":                  "u=0, i",
			"referer":                   "https://gmember.melon.com/login/login_form.htm?langCd=EN&redirectUrl=https://tkglobal.melon.com/performance/index.htm?langCd=EN&prodId=211818",
			"sec-ch-ua-mobile":          "?0",
			"sec-ch-ua-platform":        `"Windows"`,
			"sec-fetch-dest":            "document",
			"sec-fetch-mode":            "navigate",
			"sec-fetch-site":            "same-origin",
			"sec-fetch-user":            "?1",
			"upgrade-insecure-requests": "1",
		}).
		SetFormData(map[string]string{
			"rtnUrl": "https://tkglobal.melon.com/main/index.htm",
			"langCd": "EN",
			"email":  email,
			"pwd":    pwd,
		})

	// 5) 发送登录
	resp, err := r.Post("https://gmember.melon.com/login/login_proc.htm")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	res := &LoginResult{
		Client:    c,
		Status:    resp.GetStatusCode(),
		Location:  resp.GetHeader("Location"),
		SetCookie: resp.Cookies(), // 解析所有 Set-Cookie
	}
	// 读一点点 body（通常 302 没啥 body）
	b, _ := resp.ToString()
	res.Body = b

	// 可选：把新下发的 cookie 同步到“后续所有请求”的公共 cookie 中
	if len(res.SetCookie) > 0 {
		c.SetCommonCookies(res.SetCookie...)
		c.SetCommonCookies(&http.Cookie{Name: "PCID", Value: "17320157945852523050993", Domain: ".melon.com"},
			&http.Cookie{Name: "TKT_POC_ID", Value: "WP19", Domain: ".melon.com"},
			&http.Cookie{Name: "JSESSIONID", Value: "E874130B959DC968E6C7EFCF2425A99F", Domain: "gmember.melon.com"})
	}

	return res, nil
}
