package parsing

import (
	"bufio"
	"context"
	"csv-storage/db"
	"csv-storage/models"
	"csv-storage/utils"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func ParseAndStore(filePath string) {
	earseStore()
	c := make(chan string, 200)
	for i := 1; i <= 500; i++ {
		go splitLineThenStore(c)
	}
	readFile(filePath, c)
}
func readFile(filePath string, c chan string) {
	csvFile, err := os.Open(filePath)
	utils.HandleError(err)
	defer csvFile.Close()

	s := bufio.NewScanner(csvFile)
	for s.Scan() {
		c <- s.Text()
		fmt.Println(s.Text())
	}
	fmt.Println("Finished reading the file")
	close(c)
}

func splitLineThenStore(c chan string) {
	for s := range c {

		var promotion models.Promotion

		columns := strings.Split(s, ",")
		promotion.Id = columns[0]

		var err error
		promotion.Price, err = strconv.ParseFloat(columns[1], 64)
		utils.HandleError(err)

		promotion.ExpirationDate = columns[2]
		ctx := context.Background()
		db.CreatePromotion(ctx, promotion)
	}
}

func earseStore() {
	cmd := exec.Command("/bin/sh", "earse-redis.sh")
	_, err := cmd.Output()
	utils.HandleError(err)
}
