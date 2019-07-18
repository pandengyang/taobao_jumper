package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const jumpTempStr = `<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
<title>雅赞小安家母婴小铺</title>
<style>
html,body {
    margin:0;
    height:100%;
}

#mask {
    width:100%;
    height:100%;
    background-color:#000;
    position:absolute;
    top:0;
    left:0;
    z-index:2;
    opacity:0.3;
    /*兼容IE8及以下版本浏览器*/
    filter: alpha(opacity=30);
}

#mask p {
    color: white;
    position: absolute;
    top:50%;
    left:50%;
    width:100%;
    transform:translate(-50%,-50%);
    text-align: center;
}

#mask p span {
    font-weight:700;
    color: #fff;
}

#arrow {
    position: absolute;
    right: 10px;
    top: 10px;
}
</style>
</head>

<body>
<div id="mask">
    <img id="arrow" src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iNjAiIGhlaWdodD0iOTAiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyI+CiA8IS0tIENyZWF0ZWQgd2l0aCBNZXRob2QgRHJhdyAtIGh0dHA6Ly9naXRodWIuY29tL2R1b3BpeGVsL01ldGhvZC1EcmF3LyAtLT4KIDxnPgogIDx0aXRsZT5iYWNrZ3JvdW5kPC90aXRsZT4KICA8cmVjdCBmaWxsPSIjbm9uZSIgaWQ9ImNhbnZhc19iYWNrZ3JvdW5kIiBoZWlnaHQ9IjkyIiB3aWR0aD0iNjIiIHk9Ii0xIiB4PSItMSIvPgogIDxnIGRpc3BsYXk9Im5vbmUiIG92ZXJmbG93PSJ2aXNpYmxlIiB5PSIwIiB4PSIwIiBoZWlnaHQ9IjEwMCUiIHdpZHRoPSIxMDAlIiBpZD0iY2FudmFzR3JpZCI+CiAgIDxyZWN0IGZpbGw9InVybCgjZ3JpZHBhdHRlcm4pIiBzdHJva2Utd2lkdGg9IjAiIHk9IjAiIHg9IjAiIGhlaWdodD0iMTAwJSIgd2lkdGg9IjEwMCUiLz4KICA8L2c+CiA8L2c+CiA8Zz4KICA8dGl0bGU+TGF5ZXIgMTwvdGl0bGU+CiAgPHBhdGggdHJhbnNmb3JtPSJyb3RhdGUoLTcwLjE2OTYwOTA2OTgyNDIyIDMwLjM0NTYwNzc1NzU2ODM3NCw0MS42MzIyODYwNzE3NzczNSkgIiBpZD0ic3ZnXzUiIGQ9Im0zNC41MzIzNjcsNjQuODk5MTYzYzIuMDgwNCwtMC43NjM1NzYgMy43ODY3OTUsLTEuNDc0NjMyIDAuNTU1MDk1LC0xLjQ5NTQ0MmMtMi44NTQwNSwwLjM4MzU4MyAtMTAuMjM3MjI0LC0yLjQzNTg4MyAtNC4zMTk5MTUsLTMuODIxNzE4YzMuNzAxOTQ3LC0xLjQ0NTQwNCAxMi4xMjY3NTgsMS44OTMxNTEgMTIuMTcxNjY2LC0zLjgxNzUwM2MtMC42Njk2ODUsLTMuNTc4ODMzIC0xMS41MzY5NzcsLTAuNzg0NTA3IC03Ljk2MTk5MSwtNS41ODQxOTRjMy45MzcxOSwtMS42MDk4ODggMTEuNzQ0MzE3LDIuMTgzMjE2IDEzLjA2NTA3NCwtMi43OTg0MTJjLTEuMjEwNzI1LC00LjgzMDI4IC04Ljc4MzQ3MywtMS43NDc2MTYgLTEyLjU4OTcyOCwtMy4wODI1NDFjLTMuNjg5ODY1LDAuNzYxMDk2IC00LjY2MTU5LC00LjA5MDk3OCAtMC41MjAzNjYsLTMuMzU4NjE2YzkuNTczMzcyLC0wLjUxNjQ4NSAxOS4yMTM1MTksMC4yNDE1MTcgMjguNzUzOTEsLTAuODEyNDk5YzEuNTQxMDk1LC0wLjg1MzYyMyAyLjY4MDI3MSwtNi4xMDQ2MzkgMC43MzU0NzQsLTYuNDU4MjUyYy0xMy44OTc0MDcsLTAuMjg3OTY1IC0yNy44MTkzNzMsLTAuMDIzMjg0IC00MS42OTQ3ODIsLTAuODgyMjMyYy00Ljk4NDk3NiwtMC4wNTE3NzMgLTEuMTc2MTgsLTUuMzc0NzUyIDIuNTA3NDM0LC00LjEyNDY2N2MzLjUwODg1NCwtMS4xNDkwMTcgNC44OTk0MjIsLTQuMTYxODI2IDcuNDc0Njg5LC03LjEzNjczOGMxLjMxNiwtNi41MDI1OSAtNS4wMDg4MzUsLTQuMDg5MTIgLTcuNTkxODU0LC0yLjE4NjMxMmMtMi44ODQzMTUsMS44OTA5MTggLTcuODQ1NDQ1LDUuODE1OTI5IC0xMC40NzcyNzksNy44NzY0MDhjLTIuMzA3MDY3LDEuNjY4MjI0IC03LjYyMjM3OSwzLjUxOTM4MiAtMTAuODU1NzM1LDUuNzU1NjEyYy0yLjc2Mzc1NywwLjk4MjE4MiAtNS43ODgxNDEsMC40OTMxOTggLTguNTAyNjU4LDAuNjY5NTY3YzAsOC44MTAyODcgMCwxNy42MjA0NDUgMCwyNi40MzA3MjhjNy42Nzk4MjIsMy4zNDQwMDUgMTUuNzYxMTI3LDYuMjgxMzgxIDI0LjI3NTI0LDUuOTQ1MzYyYzQuOTg3NTE0LC0wLjAyMDMwOSAxMC4xMTgyODUsMC4xNjYyMTkgMTQuOTc1NzIzLC0xLjExODU0OGwwLjAwMDAwMywtMC4wMDAwMDN6bS0zMS42NDEyMjIsMi44MjU3ODVjLTMuOTIwNDM3LC0xLjUwNjM0MyAtNy44MzM3NjMsLTMuMDMxODgxIC0xMS43MTg2NzksLTQuNjI3ODk4YzAuMTg1OTEzLC0xMS4wMDIxNzIgMC4zNzE4MjQsLTIyLjAwNDQ2OCAwLjU1NzcyNywtMzMuMDA2NjQxYzMuMTAxNzUsLTAuMTE3NzkyIDcuNzA0MjM1LDEuNDIwNzU1IDkuOTE1MDg5LC0wLjg0NDcwM2M1Ljk0Mzk5NywtMS45MDc3NjIgOS44MDMxNDksLTYuMTI3MzA3IDE1LjAxMzkwMiwtOS4yOTQzMTZjMy42NjY5NzIsLTIuODgxMzk3IDYuNTk1OTc3LC0zLjg2NDgxNyAxMC4zODM0NjUsLTYuMDUwMDE5YzIuODc2NzY1LC0xLjMzMDM0MyA0LjQzNDM5NSwtMC42NTA5OTEgNy42NzM4MDUsLTAuMDcwOTcxYzEuNDEyODUzLDIuMzEyMDI5IDMuMTY2ODAxLDMuNTA2OTk5IDIuMjU2NjYxLDcuOTExNDZjLTEuODk5MTM5LDAuOTU1MzA3IC00LjM2MzcxMyw3LjM3NjAyOCAtMC42MzE1OSw3LjQyNzA1NWMxMC4zOTAyNTksMC45OTU2ODEgMjAuOTQ3ODE3LC0wLjQwMzc3NCAzMS4yNTYzNDMsMS4zMzY0MTFjMi41NjY5MDYsMy40NzUxNjUgMy4wNTY0NDQsMTEuMjgyNDYgLTEuNjQxNjA5LDEzLjQ4NjI0MWMtNC40OTU5NDMsMS4yOTk4NjkgLTkuMjM2Mjc0LDAuNjU1MDc0IC0xMy44NTU0MzMsMC44MTEzODFjLTAuMDAwMjgxLDMuMjg3NjQ4IDAuNTQzMjY1LDcuNDgzMTYgLTMuMTA1ODI2LDkuMTM0NjYzYy0yLjY5NDQ2MywzLjUxNDMwNyAtMi44ODM5NDUsOC42NzcwMTcgLTcuMjI1NzEyLDExLjI0ODUyNmMtMy42NTE1NTcsMy43NDE3MDQgLTguODQ5NjQxLDQuODk3NTQxIC0xMy45MTk4Nyw0Ljc3NDE3OGMtOC4zMTM5MjksMC4xMTI3MDkgLTE3LjA2NTA5LDAuOTA5NzIgLTI0Ljk1ODI3MiwtMi4yMzUxMjJsMCwtMC4wMDAyNDlsLTAuMDAwMDAzLDAuMDAwMDA1eiIgZmlsbC1vcGFjaXR5PSJudWxsIiBzdHJva2Utb3BhY2l0eT0ibnVsbCIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZT0iIzAwMCIgZmlsbD0iI2ZmZiIvPgogPC9nPgo8L3N2Zz4="/>
    <p>1. 点击右上角的&nbsp;<img id="dots" src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMjIiIGhlaWdodD0iMjIiIHhtbG5zPSJodHRwOi8vd3d3LnczLm9yZy8yMDAwL3N2ZyIgc3R5bGU9InZlY3Rvci1lZmZlY3Q6IG5vbi1zY2FsaW5nLXN0cm9rZTsiIHN0cm9rZT0ibnVsbCI+CiA8IS0tIENyZWF0ZWQgd2l0aCBNZXRob2QgRHJhdyAtIGh0dHA6Ly9naXRodWIuY29tL2R1b3BpeGVsL01ldGhvZC1EcmF3LyAtLT4KIDxnIHN0cm9rZT0ibnVsbCI+CiAgPHRpdGxlIHN0cm9rZT0ibnVsbCI+YmFja2dyb3VuZDwvdGl0bGU+CiAgPHJlY3Qgc3Ryb2tlPSJudWxsIiBmaWxsPSJub25lIiBpZD0iY2FudmFzX2JhY2tncm91bmQiIGhlaWdodD0iMTYiIHdpZHRoPSIyNCIgeT0iLTEiIHg9Ii0xIi8+CiAgPGcgc3R5bGU9InZlY3Rvci1lZmZlY3Q6IG5vbi1zY2FsaW5nLXN0cm9rZTsiIHN0cm9rZT0ibnVsbCIgZGlzcGxheT0ibm9uZSIgb3ZlcmZsb3c9InZpc2libGUiIHk9IjAiIHg9IjAiIGhlaWdodD0iMTAwJSIgd2lkdGg9IjEwMCUiIGlkPSJjYW52YXNHcmlkIj4KICAgPHJlY3QgZmlsbD0idXJsKCNncmlkcGF0dGVybikiIHN0cm9rZS13aWR0aD0iMCIgeT0iMCIgeD0iMCIgaGVpZ2h0PSIxMDAlIiB3aWR0aD0iMTAwJSIvPgogIDwvZz4KIDwvZz4KIDxnIHN0cm9rZT0ibnVsbCI+CiAgPHRpdGxlIHN0cm9rZT0ibnVsbCI+TGF5ZXIgMTwvdGl0bGU+CiAgPGVsbGlwc2Ugc3Ryb2tlPSIjZGVkZWRlIiByeT0iMiIgcng9IjIiIGlkPSJzdmdfMSIgY3k9IjkuNzUiIGN4PSIxOS43NSIgc3Ryb2tlLXdpZHRoPSIxLjUiIGZpbGw9IiNkZWRlZGUiLz4KICA8ZWxsaXBzZSByeT0iMiIgcng9IjIiIGlkPSJzdmdfNCIgY3k9IjkuNzUiIGN4PSIxMC43NSIgc3Ryb2tlLW9wYWNpdHk9Im51bGwiIHN0cm9rZS13aWR0aD0iMS41IiBzdHJva2U9IiNkZWRlZGUiIGZpbGw9IiNkZWRlZGUiLz4KICA8ZWxsaXBzZSBzdHJva2U9IiNkZWRlZGUiIHJ5PSIyIiByeD0iMiIgaWQ9InN2Z182IiBjeT0iOS43NSIgY3g9IjIuNzUiIHN0cm9rZS1vcGFjaXR5PSJudWxsIiBzdHJva2Utd2lkdGg9IjEuNSIgZmlsbD0iI2RlZGVkZSIvPgogPC9nPgo8L3N2Zz4="/>
        <br />
        2. 选择&nbsp;<span>在浏览器打开</span>&nbsp;或&nbsp;<span>在Safari中打开</span>&nbsp;即可
    </p>
</div>
</body>
</html>
`

var (
	Info *log.Logger
)

func init() {
	file, err := os.OpenFile("log/taobao.maizimart.com.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file: ", err)
	}

	Info = log.New(file, "TAOBAO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	mux := httprouter.New()
	mux.GET("/shops/:id", shops)
	mux.GET("/items/:id", items)

	server := http.Server{
		Addr:    "127.0.0.1:10002",
		Handler: mux,
	}
	server.ListenAndServe()
}

func shops(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("shops error: %v", p)
			Info.Println(err)
		}

	}()

	userAgent := ""
	if r.Header != nil {
		if ua := r.Header["User-Agent"]; len(ua) > 0 {
			userAgent = ua[0]
		}
	}

	if strings.Contains(userAgent, "MicroMessenger") {
		Info.Println("MicroMessenger")
		io.WriteString(w, jumpTempStr)
	} else {
		// eg: https://shop362549991.taobao.com
		Info.Println("Other")
		w.Header().Set("Location", "https://shop"+p.ByName("id")+".taobao.com")
		w.WriteHeader(302)
	}
}

func items(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("items error: %v", p)
			Info.Println(err)
		}

	}()

	userAgent := ""
	if r.Header != nil {
		if ua := r.Header["User-Agent"]; len(ua) > 0 {
			userAgent = ua[0]
		}
	}

	if strings.Contains(userAgent, "MicroMessenger") {
		Info.Println("MicroMessenger")
		io.WriteString(w, jumpTempStr)
	} else {
		// eg: https://item.taobao.com/item.htm?id=591547056129
		Info.Println("Other")
		w.Header().Set("Location", "https://item.taobao.com/item.htm?id="+p.ByName("id"))
		w.WriteHeader(302)
	}
}
