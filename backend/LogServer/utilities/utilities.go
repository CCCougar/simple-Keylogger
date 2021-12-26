package utilities

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

const FILETIMEOUT = 5 * time.Minute

// 将标准logger输出重定向到一个文件
func InitLog(file string) {
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
}

// 判断所给路径文件/文件夹是否存在
func Filexists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 通过data-form的形式以POST包向服务器发送数据
func Upload(url string, values map[string]io.Reader) (err error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			log.Println("%s\n", fw)
			return err
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// fmt.Printf("%s\n", b)

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)
	// var k bytes.Buffer
	// req, err := http.NewRequest("POST", url, &k)
	if err != nil {
		return err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())
	// fmt.Println(w.FormDataContentType())

	// fmt.Printf("%s", req)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   FILETIMEOUT, // 文件（如：截屏，文件监控，录音）等，超时时间为
	}

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Check the response
	if res.StatusCode != http.StatusOK {
		log.Println("bad status: %s", res.Status)
		err = fmt.Errorf("bad status: %s", res.Status)
		return err
	}
	defer res.Body.Close()

	respBody, _ := ioutil.ReadAll(res.Body)

	// Debug
	log.Printf("[[[resp.Body %s: %s]]]", values["device_id"], values["data"], string(respBody))
	return nil
}

func CheckProcess(pid int) bool {
	_, err := os.FindProcess(pid)
	if err != nil {
		return false
	} else {
		return true
	}
}
