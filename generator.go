package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const dest, template, conf_d string = "demo_dir", "templates", "conf.d"

func openFile(srcName string) (fi *os.File, err error) {
	err = os.MkdirAll(path.Dir(srcName), 0777)
	if err != nil {
		return
	}

	fi, err = os.Create(srcName)
	return
}

func main() {
	//获取keys
	keys := [3]string{"/demo1/test", "/demo2/test", "/demo3/test"}

	//读取或创建template文件，依次写入template,读取或创建conf.d文件依次写入
	fi, err := os.Open("template.template")
	if err != nil {
		print(err)
	}
	defer fi.Close()
	fd_template, err := ioutil.ReadAll(fi)
	if err != nil {
		print(err)
	}

	fc, err := os.Open("confd.template")
	if err != nil {
		print(err)
	}
	defer fc.Close()
	fd_confd, err := ioutil.ReadAll(fc)
	if err != nil {
		print(err)
	}

	for _, key := range keys {
		key_template_name := template + key + ".tmpl"
		key_confd_name := conf_d + key + ".toml"
		fi_template, err := openFile(key_template_name)
		if err != nil {
			print(err)
		}

		fi_confd, err := openFile(key_confd_name)
		if err != nil {
			print(err)
		}

		wt := bufio.NewWriter(fi_template)
		fmt.Fprintf(wt, string(fd_template), key)
		wc := bufio.NewWriter(fi_confd)
		fmt.Fprintf(wc, string(fd_confd), key_template_name, dest+key, key)
		wt.Flush()
		wc.Flush()
		fi_template.Close()
		fi_confd.Close()
	}

}
