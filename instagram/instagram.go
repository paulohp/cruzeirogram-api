package instagram

import "net/http"
import "log"
import "encoding/json"
import "fmt"
import "io/ioutil"
import "os"
import "github.com/gorilla/mux"

type Pic struct {
	Name    string `bson:"name"`
	Picture string `bson:"picture"`
}
