package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = "4000"
	}
	http.HandleFunc("/", hello)
	fmt.Println("listening at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	i, err := strconv.Atoi(os.Getenv("CF_INSTANCE_INDEX"))
	if err != nil {
		i = -1
	}
	fmt.Fprintln(res, hackertacle(toEmoji(i)))
}

func toEmoji(i int) string {
	switch i {
	case 0:
		return "0️⃣ "
	case 1:
		return "1️⃣ "
	case 2:
		return "2️⃣ "
	case 3:
		return "3️⃣ "
	case 4:
		return "4️⃣ "
	case 5:
		return "5️⃣ "
	case 6:
		return "6️⃣ "
	case 7:
		return "7️⃣ "
	case 8:
		return "8️⃣ "
	case 9:
		return "9️⃣ "
	case 10:
		return "🔟 "
	case 11:
		return "🍜 "
	default:
		return "🎾 "
	}
}

func hakata(s string) string {
	return strings.Replace("*  *  **   *  *   **   ****  ** \n"+
		"*  * *  *  * *   *  *     *    *  *\n"+
		"*** ***  **    ***    *    ***\n"+
		"*  * *   * *  *  *   *    *    *  *\n"+
		"*  * *   * *   * *   *    *    *  *\n", "*", s, -1)
}

func hackertacle(s string) string {
	return strings.Replace("*  *  **    *** *   * ***  ***  \n"+
		"*  * *  *  *      *  *  *      *   * \n"+
		"*** ***  *      **    **    *** \n"+
		"*  * *   * *      *  *  *      *   * \n"+
		"*  * *   *  *** *   * ***  *   * \n"+
		"                                                 \n"+
		"****  **    *** *   * *      *** \n"+
		"   *    *  *  *      *  *  *      *\n"+
		"   *    ***  *      **    *      **   \n"+
		"   *    *   * *      *  *  *      * \n"+
		"   *    *   *  *** *   * ***  *** \n", "*", s, -1)
}
