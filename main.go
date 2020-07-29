package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
var fileName string

func init()  {
	flag.StringVar(&fileName,"path","default","nil")
}

func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".json")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}
func Format(fileName string)  {
	saveFileName :=fileName[:len(fileName)-5]+"_FORMATED.json"

	filePtr, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Open file failed [Err:%s]", err.Error())
		return
	}
	defer filePtr.Close()

	byteValue,_ :=ioutil.ReadAll(filePtr)
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue),&result)
	//fmt.Println(result)

	// 创建文件
	fileSavePtr, err := os.Create(saveFileName)
	if err != nil {
		fmt.Println("Create file failed", err.Error())
		return
	}
	defer fileSavePtr.Close()

	// 带JSON缩进格式写文件　
	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println("Format failed", err.Error())
	} else {
		fmt.Println("Format success")
	}
	fileSavePtr.Write(data)
}
func main() {
	//暂停获取参数filename
	flag.Parse()

	files,_,err :=GetFilesAndDirs(fileName)
	if err!=nil{
		fmt.Println("export dir fail!")
		return
	}

	if(fileName=="default"){
		fmt.Println("Useage: -path FILEPATH")
		return
	}

	for i,file := range files{
		fmt.Println("Format No.",i+1," FilePath:"+file)
		Format(file)
	}

}
