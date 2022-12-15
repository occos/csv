package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// WriteCsvFile 往csv文件中写入数据
func WriteCsvFile(fileName string, row []string) {
	//fileName := fmt.Sprintf("./data-%s.csv", time.Now().Format("20060102_150405"))
	//fileName := fmt.Sprintf("./data-%s.csv", time.Now().Format("20060102_150405"))
	//fileName := fmt.Sprintf("./data-%s.csv", time.Now().Format("20060102_150405"))

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
	}
	defer file.Close()
	file.Seek(0, io.SeekEnd)

	w := csv.NewWriter(file)
	w.Comma = ','
	w.UseCRLF = true

	//row := []string{userInfo.Result.Userid, userInfo.Result.Name, userInfo.Result.JobNumber, userInfo.Result.Mobile, userInfo.Result.Email}
	//row := []string{userInfo.Result.Userid, userInfo.Result.Name, userInfo.Result.JobNumber, userInfo.Result.Mobile, userInfo.Result.Email}
	//row := []string{userInfo.Result.Userid, userInfo.Result.Name, userInfo.Result.JobNumber, userInfo.Result.Mobile, userInfo.Result.Email}

	if err1 := w.Write(row); err1 != nil {
		log.Println("写入文件失败...")
	}
	w.Flush()

	log.Println("数据写入成功...")
	log.Printf("%s \n\n", fileName)
}
