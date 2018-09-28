package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Connect to the database
	//dbConn := fmt.Sprintf("minibank:minibank@tcp(mysql)/minibank")
	//models.InitDB(dbConn)
	fmt.Println("Hello World")
	http.HandleFunc("/api/account/register", RegisterHandler)

	http.ListenAndServe(port(), nil)
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
	fmt.Println("Hello World")
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
