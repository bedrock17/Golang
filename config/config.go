package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Config 구조체에 서버 실행에 관련된 정보를 저장
type Config struct {
	Port int
}

//지정 해준 경로에 파일을 파싱해 설정값을 로딩
func (conf *Config) Load(path string) error {

	file, ferr := os.Open("serverconfig.ini")

	if ferr != nil {
		return ferr
	}

	defer file.Close()

	fileInfo, fierr := file.Stat()

	if fierr != nil {
		return fierr
	}

	data := make([]byte, fileInfo.Size())

	file.Read(data)

	inistring := string(data[:])

	// fmt.Println(inistring)
	conf.serverInfoParse(inistring)

	return nil

}

//server 정보를 읽어오는 부분
func (conf *Config) serverInfoParse(ini string) error {

	idx := strings.Index(ini, "server")

	if idx == -1 { //server string not found
		return fmt.Errorf("invalid ini file [server] not found")
	}

	portString := "port="
	idx = strings.Index(ini, portString)
	if idx == -1 {
		return fmt.Errorf("port info not found")
	}
	endidx := strings.Index(ini[idx:], "\n")

	if endidx == -1 {
		return fmt.Errorf("port info not found (end char)")
	}

	endidx += idx
	idx += len(portString)

	var err error
	conf.Port, err = strconv.Atoi(ini[idx : endidx-1])

	// fmt.Println(idx, endidx, ini[idx:])
	// fmt.Printf("server port [%s]\n", ini[idx:endidx-1])
	// fmt.Println(conf.Port)

	if err != nil {
		return err
	}

	return nil
}
