package main

import (
	"crypto/sha1"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/satori/go.uuid"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	//process form submission
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf") // 获取上传的文件名
		if err != nil {
			fmt.Println(err)
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		}
		defer mf.Close() // 关闭 文件句柄

		// create sha for file name   随机生成文件名
		ext := strings.Split(fh.Filename, ".")[1] // 获取扩展名
		h := sha1.New()
		// copy uploaded file to new sha1 filename 上传文件到新的SHA1文件副本
		io.Copy(h, mf)
		// add file extension to this filename, note that filename is hexadecimal
		// 将文件扩展名添加到这个文件名，注意文件名是十六进制
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname) // 文件名绝对路径
		nf, err := os.Create(path)                         // 创建文件
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()
		// copy, note need to reset read/write head because of io.Copy above
		// 复制，注释需要重写读写头，因为IO。
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs[1:])
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
			Path:  "/",
		}
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	// append
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}

	c.Value = s
	http.SetCookie(w, c)
	return c
}
