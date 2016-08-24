package main

import (
	"bufio"
	//"gitlab.renrenche.com/web/artemis/service/common"
	"os"
	//"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {

		text, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		out := MyProcess(text)
		os.Stdout.WriteString(out + "\n")

	}
}

//just say hello.
func MyProcess(in []byte) string {
	return "hello pipe " + string(in)
}

//jiemi
//carID, _ := strconv.Atoi(text[0:len(text)-1])
//os.Stdout.WriteString(common.Encrypt(carID))
//os.Stdout.WriteString("\n")
//jiami
//temp := common.Decrypt(string(text))
//os.Stdout.WriteString(strconv.Itoa(temp))
