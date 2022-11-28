package email

import (
	"os"
	"strconv"
)

func EmailNotice(emailAddr, greet, res string) (err error) {
	message := Message{
		Title:   "疫情打卡",
		Email:   emailAddr,
		Content: template(greet, "打卡结果："+res),
	}
	sendCinfig := SendConfig{
		Username:  "COVID",
		EmailAddr: os.Getenv("COVID_EMAIL_ADDR"),
		Password:  os.Getenv("COVID_EMAIL_PASSWORD"),
		Host:      os.Getenv("COVID_EMAIL_HOST"),
		Port:      s2i(os.Getenv("COVID_EMAIL_PORT")),
	}
	return sendEmail(&message, &sendCinfig)
}

func template(greet string, content string) string {
	return `<head><base target="_blank" /><style type="text/css">::-webkit-scrollbar{ display: none; }</style><style id="cloudAttachStyle" type="text/css">#divNeteaseBigAttach, #divNeteaseBigAttach_bak{display:none;}</style><style id="blockquoteStyle" type="text/css">blockquote{display:none;}</style><style type="text/css">body{font-size:14px;font-family:arial,verdana,sans-serif;line-height:1.666;padding:0;margin:0;overflow:auto;white-space:normal;word-wrap:break-word;min-height:100px}td, input, button, select, body{font-family:Helvetica, 'Microsoft Yahei', verdana}pre {white-space:pre-wrap;white-space:-moz-pre-wrap;white-space:-pre-wrap;white-space:-o-pre-wrap;word-wrap:break-word;width:95%}th,td{font-family:arial,verdana,sans-serif;line-height:1.666}img{ border:0}header,footer,section,aside,article,nav,hgroup,figure,figcaption{display:block}blockquote{margin-right:0px}</style></head><body tabindex="0" role="listitem"><table width="700" border="0" align="center" cellspacing="0" style="width:700px;"><tbody><tr><td><div style="width:700px;margin:0 auto;border-bottom:1px solid #ccc;margin-bottom:30px;"><table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;"><tbody><tr><td width="210"></td></tr></tbody></table></div><div style="width:680px;padding:0 10px;margin:0 auto;"><div style="line-height:1.5;font-size:14px;margin-bottom:25px;color:#4d4d4d;"><strong style="display:block;margin-bottom:15px;">尊敬的用户：<span style="color:#AF2125;font-size: 16px;">` + greet + `</span>您好！</strong></div><div><strong style="display:block;margin-bottom:15px;">` + content + `</strong></div><div style="margin-bottom:30px;"><small style="display:block;margin-bottom:20px;font-size:12px;"><p style="color:#747474;">注意：该邮件为系统发送，请勿回复！</p><p style="color:#747474;"> © 2022 <a href="https://www.aecra.cn">aecra.cn</a> 版权所有</p></small></div></div><div style="width:700px;margin:0 auto;border-top:1px solid #ccc;"><table border="0" cellpadding="0" cellspacing="0" width="700" height="39" style="font:12px Tahoma, Arial, 宋体;"><tbody><tr><td width="210"></td></tr></tbody></table></div></td></tr></tbody></table></body>`
}

func s2i(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
