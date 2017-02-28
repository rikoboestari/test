package connection

import (
	"database/sql"
	_ "github.com/lib/pq"

	"github.com/rikoboestari/test/src/aservice"
	"fmt"
	"log"
)

func QueryTalk(howmany int) []aservice.Talk {
	db, err := sql.Open("postgres", "user=techacademy password=123qwe!@#QWE dbname=tokopedia-talk host=192.168.100.126 port=5432 sslmode=disable")
	defer db.Close()

	if err!=nil {
		fmt.Println("db con err")
		return nil
	}
	fmt.Println("db con success")
	rows, err := db.Query("select * from ws_talk order by talk_id limit 10")

	if err!=nil {
		fmt.Println("query err", err)
		return nil
	}
	for rows.Next() {
		talkId, shopId, productId := -1, -1, -1


		err := rows.Scan(&talkId, &shopId, &productId)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(talkId, shopId, productId)

	}
	return nil
}
