package main

import (
	"bufio"
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"os"
	"runtime"
	"strings"
)

import (
	"github.com/anknown/ahocorasick"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("Error is ", err)
		os.Exit(-1)
	}
}

func ReadRunesByMysql() ([][]rune, error) {
	db, err := sql.Open("mysql", "root:xxx@tcp(127.0.0.1:3306)/xxx?charset=utf8")
	checkErr(err)

	//遍历敏感词表所有数据
	page := 1
	count := 100
	var dict [][]rune
	for{
		offset := (page-1)*count
		stmt, err := db.Prepare("select id, words from sensitive_words order by id limit ?, ?")
		checkErr(err)
		rows, err := stmt.Query(offset, count)
		checkErr(err)
		if !rows.Next() {
			break
		}else{
			var id int
			var words string
			err = rows.Scan(&id, &words)
			checkErr(err)
			fmt.Println(id)
			fmt.Println(words)
			dict = append(dict , []rune(words))
		}

		//查询数据
		for rows.Next(){
			var id int
			var words string
			err = rows.Scan(&id, &words)
			checkErr(err)
			fmt.Println(id)
			fmt.Println(words)
			dict = append(dict , []rune(words))
		}
		page++
	}
	return dict, nil
}

func main() {
	dict, err := ReadRunesByMysql()
	if err != nil {
		fmt.Println(err)
		return
	}

	m := new(goahocorasick.Machine)
	if err := m.Build(dict); err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()

	router.GET("/match", func(c *gin.Context) {
		words := c.Query("words")
		// 去除空格
		words = strings.Replace(words, " ", "", -1)
		// 去除换行符
		words = strings.Replace(words, "\n", "", -1)

		if words == "" {
			c.JSON(200, gin.H{
				"message": "没有参数",
			})
		}
		terms := m.MultiPatternSearch([]rune(words), false)
		var res []map[string]string
		for _, t := range terms {
			item := map[string]string{"pos": string(t.Pos), "word": string(t.Word)}
			res = append(res, item)
			//fmt.Printf("%d %s\n", t.Pos, string(t.Word))

		}
		c.JSON(200, gin.H{
			"message": res,
		})
	})

	router.Run(":8282")

}
