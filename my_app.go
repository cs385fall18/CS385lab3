package main

import (
	//"fmt"
	"log"
	"net/http"
	"os"
	//"io"
)

func main() {
	// Connect to the database
	//dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	//models.InitDB(dbConn)
	//fmt.Println("Hello World")
	http.HandleFunc("/", RegisterHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//http.ListenAndServe(port(), nil)
	//fs := http.FileServer(http.Dir("static/"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//http.ListenAndServe(":3001", nil)
	//fmt.Println("got here")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
/*
func (r Registration) ToJSON() string {
        jsonbytes, err := json.Marshal(r)
        if err != nil {
                log.Panic(err)
        }
        return string(jsonbytes)
}
*/
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("doesn't print")
	//fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
	name, err := os.Hostname()
	//resp, err := http.PostForm("http://example.com/form",
	//url.Values{"key": {"Value"}, "id": {"123"}})

	//fmt.Println("", name)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
	if err != nil {
        panic(err)
    }
	/*
        if r.Method == http.MethodPost {
                body, err := ioutil.ReadAll(r.Body)
                if err != nil {
                        w.WriteHeader(http.StatusBadRequest)
                        w.Write([]byte("Unable to read request body."))
                }
                registration := Registration{}
                err = json.Unmarshal(body, &registration)
                if err != nil {
                        w.WriteHeader(http.StatusBadRequest)
                        w.Write([]byte("Unable to parse registration data."))
                } else {
                        strlength := len(registration.Password)
                        if strlength >= 10 {
                                hashedpw, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
                                if err != nil {
                                        w.WriteHeader(http.StatusInternalServerError)
                                        w.Write([]byte("Unable to process request."))
                                }

				if err != nil {
                                        w.WriteHeader(http.StatusInternalServerError)
                                        w.Write([]byte("Unable to register new account"))
                                } else {
                                        last_id, _ := res.LastInsertId()
                                        w.Write([]byte(fmt.Sprintf("Successfully registered account %s", string(last_id))))
                                }
                        } else {
                                w.WriteHeader(http.StatusBadRequest)
                                w.Write([]byte("Passwords must be at least 10 characters long"))
                        }
                }
        } else {
                w.WriteHeader(http.StatusBadRequest)
                w.Write([]byte("Unsupported Method."))
        }
	*/
}

