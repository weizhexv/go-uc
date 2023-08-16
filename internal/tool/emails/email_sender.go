package emails

import (
	"bytes"
	dginfra "dghire.com/libs/go-infra-sdk"
	"errors"
	"github.com/google/uuid"
	"go-uc/internal/tlog"
	"go-uc/internal/tool/httpclient"
	"go-uc/internal/tool/nonces"
	"go-uc/pkg/var/domains"
	"go-uc/pkg/vo"
	"go-uc/vconfig"
	"html/template"
	"os"
	"strings"
)

type HrefPlaceHolder struct {
	DomainName string
	Email      string
	Nonce      string
}

type HtmlPlaceHolder struct {
	Href       string
	DomainName string
}

type EmailParams struct {
	Subject    string
	ToEmail    string
	DomainName string

	HrefTemplate        string
	HtmlContentTemplate *template.Template
}

var emailClient = initEmailClient()

func initEmailClient() *dginfra.InfraSendEmailClient {
	return dginfra.NewSendEmailClient(httpclient.Ins(), vconfig.HostInfra())
}

type Locale string

const (
	SubjectZH = "欢迎来到DG Hire"
	SubjectEN = "Welcome to DG Hire"
)

var (
	localeMap = map[domains.Domain]string{
		domains.Business: "zh",
		domains.Platform: "zh",
		domains.Supplier: "en",
		domains.Employee: "en",
	}
	subjectMap = map[domains.Domain]string{
		domains.Business: SubjectZH,
		domains.Platform: SubjectZH,
		domains.Supplier: SubjectEN,
		domains.Employee: SubjectEN,
	}
)

func SendRegisterEmail(ctx *tlog.Ctx, detail *vo.UserDetail) error {
	ctx.Infof("send register email %v", detail)
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(rootPath + "/tmpl/" + getLocale(detail.Domain) + "/register.html")
	if err != nil {
		return err
	}
	p := EmailParams{
		ToEmail:             detail.Email,
		Subject:             getSubject(detail.Domain),
		DomainName:          detail.DomainName,
		HrefTemplate:        getHref(detail.Domain),
		HtmlContentTemplate: tmpl,
	}
	ctx.Infof("register email params %v", p)

	if err = doSend(ctx, p); err != nil {
		return err
	}
	return nil
}

func SendResetPasswordEmail(ctx *tlog.Ctx, detail *vo.UserDetail) error {
	ctx.Infof("send reset password %v", detail)
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(rootPath + "/tmpl/" + getLocale(detail.Domain) + "/reset_password.html")
	if err != nil {
		return err
	}
	p := EmailParams{
		ToEmail:             detail.Email,
		Subject:             getSubject(detail.Domain),
		DomainName:          detail.DomainName,
		HrefTemplate:        getHref(detail.Domain),
		HtmlContentTemplate: tmpl,
	}
	ctx.Infof("reset password email params %v", p)
	if err = doSend(ctx, p); err != nil {
		return err
	}
	return nil
}

func SendInviteEmployeeEmail(ctx *tlog.Ctx, domainName string, email string) error {
	ctx.Infof("send invite employee email, domainName %s, email %s", domainName, email)
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	tmpl, err := template.ParseFiles(rootPath + "/tmpl/" + getLocale(domains.Employee) + "/invite_employee.html")
	if err != nil {
		return err
	}
	p := EmailParams{
		Subject:             SubjectEN,
		ToEmail:             email,
		DomainName:          domainName,
		HrefTemplate:        vconfig.EmailHrefInvite(),
		HtmlContentTemplate: tmpl,
	}
	ctx.Infof("invite employee email params %v", p)
	if err = doSend(ctx, p); err != nil {
		return err
	}
	return nil
}

func doSend(ctx *tlog.Ctx, params EmailParams) error {
	//解析邮件模版
	htmlContent, err := parseHtmlContent(ctx, params)
	if err != nil {
		return err
	}

	//发送邮件
	ret := emailClient.SendEmail(ctx.DgContext, &dginfra.SendEmailRequest{
		To:      []string{params.ToEmail},
		Subject: params.Subject,
		Content: htmlContent,
	})

	ctx.Infof("send email ret %v", ret)
	if !ret.Success {
		return errors.New(ret.Message)
	} else {
		return nil
	}
}

func parseHtmlContent(ctx *tlog.Ctx, params EmailParams) (string, error) {
	//生成nonce
	nonce, err := nonces.New(ctx, params.ToEmail)
	if err != nil {
		return "", err
	}

	//生成href
	hrefTmpl, err := template.New("href_" + uuid.NewString()).Parse(params.HrefTemplate)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	err = hrefTmpl.Execute(buf, HrefPlaceHolder{
		DomainName: encodeUrl(params.DomainName),
		Email:      encodeUrl(params.ToEmail),
		Nonce:      encodeUrl(nonce),
	})

	if err != nil {
		return "", err
	}
	href := string(buf.Bytes())
	ctx.Infof("email url link: %v", href)

	//生成邮件内容
	buf = new(bytes.Buffer)
	err = params.HtmlContentTemplate.Execute(buf, HtmlPlaceHolder{
		Href:       href,
		DomainName: params.DomainName,
	})
	if err != nil {
		return "", err
	}
	return string(buf.Bytes()), nil
}

func encodeUrl(val string) string {
	val = strings.ReplaceAll(val, " ", "%20")
	val = strings.ReplaceAll(val, "&", "%26")
	return val
}

func getLocale(domain domains.Domain) string {
	return localeMap[domain]
}
func getSubject(domain domains.Domain) string {
	return subjectMap[domain]
}

func getHref(domain domains.Domain) string {
	if domain.Is(domains.Supplier) {
		return vconfig.EmailHrefS()
	} else if domain.Is(domains.Platform) {
		return vconfig.EmailHrefA()
	} else if domain.Is(domains.Business) {
		return vconfig.EmailHrefB()
	} else {
		return vconfig.EmailHrefC()
	}
}
