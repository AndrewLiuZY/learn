package main

import (
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		intarr := make([]float64, 0)
		for _, item := range strings.Split(request.FormValue("in"), " ") {
			num, _ := strconv.ParseFloat(item, 6)
			intarr = append(intarr, num)
		}
		io.WriteString(w, fmt.Sprintf("<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=\"1.0\"><title>Document</title></head><body><h1>Statistics</h1><p>Computes basic statistics for a given list of numbers</p><text>Numbers(comma or space-separared)</text><form action=\"/\" name=\"bar\" method=\"post\"><input type=\"text\" name=\"in\" /><br/><input type=\"submit\" value=\"calculate\"/></form><br><table border=\"1px\"><tbody><th colspan=\"3\">Result</th><tr><td> Numbers </td><td colspan=\"2\">%v</td></tr><tr><td> Count </td><td colspan=\"2\">%d</td></tr><tr><td> Mean </td><td colspan=\"2\">%f</td></tr><tr><td> Median </td><td colspan=\"2\">%f</td></tr></tbody></table></body></html>", request.FormValue("in"), len(intarr), Mean(intarr), Median(intarr)))
	}
}

func main() {
	http.HandleFunc("/", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func Mean(s []float64) float64 {
	sum := 0.0
	for _, item := range s {
		sum += item
	}
	return sum / float64(len(s))
}

func Median(s []float64) float64 {
	if len(s) <= 0 {
		return 0
	}
	sort.Float64s(s)
	for {
		if len(s) == 2 {
			return (s[0] + s[1]) / 2
		} else if len(s) == 1 {
			return s[0]
		}
		if len(s) <= 0 {
			return 0
		}
		s = s[1 : len(s)-1]
	}
}

const form = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>Statistics</h1>
    <p>Computes basic statistics for a given list of numbers</p>
    <text>Numbers(comma or space-separared)</text>
    <form action="/" name="bar" method="post">
		<input type="text" name="in" />
		<br/>
        <input type="submit" value="calculate"/>
    </form>
    <br>
    <table border="1px">
        <tbody>
            <th colspan="3">
                Result
            </th>
            <tr>
                <td> Numbers </td>
                <td colspan="2">1 2 3 4 5</td>
            </tr>
            <tr>
                <td> Count </td>
                <td colspan="2">5</td>
            </tr>
            <tr>
                <td> Mean </td>
                <td colspan="2">3</td>
            </tr>
            <tr>
                <td> Median </td>
                <td colspan="2">3</td>
            </tr>
        </tbody>
    </table>
</body>
</html>
`
