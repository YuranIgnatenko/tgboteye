package bot

import (
	"errors"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func LineGet(line string) string {
	re := exec.Command("bash", "-c", line)
	res, _ := re.Output()
	return string(res)
}

func ConvertStatus(s string) (string, bool, error) {
	status_user := strings.Split(s, ":")[0]
	var b bool = false
	var err error = nil
	switch status_user {
	case StatusUserNotFound:
		b = false
		err = errors.New(StatusUserNotFound)
	case StatusUserOnline:
		b = true
	case StatusUserOffline:
		b = false
	case StatusUserRecently:
		b = false
		err = errors.New(StatusUserRecently)
	}
	return status_user, b, err
}

func LoadTasks() ([]string, error) {
	content, err := ioutil.ReadFile("tasks.save")
	if err != nil {
		log.Println(err)
		return []string{}, err
	}
	data:=string(content) 
	if strings.TrimSpace(data) == "" {
		return []string{}, nil
	} else{
		list:= strings.Split(data,"\n")
		return list, nil
	} 
}

func SaveTasks(ts []*Task) error {
	line := "" 
	for _, t := range ts {
		line += t.TargetUser + "\n" 
	}
	line = line[:len(line)-1]
	err := ioutil.WriteFile("tasks.save", []byte(line), 0755)
	return err
}
