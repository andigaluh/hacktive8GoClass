package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Elemen struct {
	Water, Wind int
}

func index(w http.ResponseWriter, r *http.Request) {
	var watervalue, windvalue int
	watervalue = rand.Intn(100)
	windvalue = rand.Intn(100)
	data := Elemen{
		Water: watervalue,
		Wind:  windvalue,
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	jsonFile, err := os.Open("test.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("Successfully Opened sess8.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Elemen

	json.Unmarshal([]byte(byteValue), &users)

	const tpl = `
        <!DOCTYPE html>
        <html lang="en">
          <head>
            <!-- Required meta tags -->
            <meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<meta http-equiv="refresh" content="5; URL=http://localhost:9090/index">

            <!-- Bootstrap CSS -->
            <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" integrity="sha384-ggOyR0iXCbMQv3Xipma34MD+dH/1fQ784/j6cY/iJTQUOhcWr7x9JvoRxT2MZw1T" crossorigin="anonymous">

            <title>Session 8</title>
          </head>
          <body>
          <div class="row">
            <div class="col-md-6">
                <table class="table">
                    <thead>
                        <tr>
						  <th scope="col">Elemen</th>
						  <th scope="col">Value</th>
						  <th scope="col">Status</th>
                        </tr>
                    </thead>
                    <tbody>
							<tr>
							  <td>Water</td>
							  <td>{{ .Water }} meter</td>
							  {{$wtr:=.Water}}
							  <td>
								{{if lt $wtr 30}}
								<span style="color:green">Aman</span>
								{{else if and (ge $wtr 30 ) (le $wtr 60)}}
								<span style="color:orange">Siaga</span>
								{{else if gt $wtr 60}}
								<span style="color:red">Bahaya</span>
								{{end}}
							  </td>
							</tr>
							<tr>
							  <td>Wind</td>
							  <td>{{ .Wind }} meter / detik</td>
							  {{$wind:=.Wind}}
							  <td>
								{{if lt $wind 30}}
								<span style="color:green">Aman</span>
								{{else if and (ge $wind 30 ) (le $wind 60)}}
								<span style="color:orange">Siaga</span>
								{{else if gt $wind 60}}
								<span style="color:red">Bahaya</span>
								{{end}}
							  </td>
                            </tr>
                    </tbody>
                </table>
            </div>
          </div>
            <!-- Optional JavaScript -->
            <!-- jQuery first, then Popper.js, then Bootstrap JS -->
            <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
            <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.7/umd/popper.min.js" integrity="sha384-UO2eT0CpHqdSJQ6hJty5KVphtPhzWj9WO1clHTMGa3JDZwrnQq4sF86dIHNDz0W1" crossorigin="anonymous"></script>
            <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/js/bootstrap.min.js" integrity="sha384-JjSmVgyd0p3pXB1rRibZUAYoIIy6OrQ6VrjIEaFf/nJGzIxFDsf4x0xIM+B07jRM" crossorigin="anonymous"></script>
          </body>
        </html>`

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := template.New("index").Parse(tpl)
	check(err)

	err = t.Execute(w, users)
	check(err)
}

func student(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It's Works!!")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/student", student) // set router
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Error running service: ", err)
	}
}
