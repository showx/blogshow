
/**
 * 匹配imei号码包
 * Author:show
 * 需要去除重复的设备码，以及不正确的设备码（比如字母开头或者字符不在14-15位之间的）
 */
package main

import "fmt"
import "os"
import (
	"bufio"
	"io"
	"strings"
//	"unicode/utf8"
	"io/ioutil"
	"unicode"
)
func Ioutil(name string) {
		if contents,err := ioutil.ReadFile(name);err == nil {
				//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
				result := strings.Replace(string(contents),"\n","",1)
				fmt.Println(len(result))
				fmt.Println("Use ioutil.ReadFile to read a file:",result)
		}
}
func main() {
	
	fi, err := os.Open("1.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()
	
	
	outputFile, outputError := os.OpenFile("ok.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	
	outputFile2, outputError2 := os.OpenFile("error.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError2 != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return  
	}
	defer outputFile2.Close()

	outputWriter2 := bufio.NewWriter(outputFile2)
	
//	var m1 map[string]string
	m1 := make(map[string]string)
	
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		cc := strings.Replace(string(a),"\n","",-1)
		cc = strings.Replace(string(cc)," ","",-1)
		strlen := len(cc)
		szstr := string(a);
		
		
		
		if strlen == 0 {
			fmt.Println("you didn't type add!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			continue
		}
		isInt := unicode.IsDigit([]rune(szstr)[0])
		
		
		nofound := true
		if _, ok := m1[szstr]; ok {
			nofound = false
		} else {
			nofound = true
		}
		
		m1[szstr] = "1"
		
		if( (strlen>=14 && strlen<16)  && isInt && nofound ){
			fmt.Println("ok:"+szstr);
			outputWriter.WriteString(szstr+"\n")
		}else{
			fmt.Println("err"+szstr)
			outputWriter2.WriteString(szstr+"\n")
		}
	}
	outputWriter.Flush()
	outputWriter2.Flush()
	return
		
}
