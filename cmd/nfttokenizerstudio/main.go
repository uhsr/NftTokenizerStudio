// cmd/nfttokenizerstudio/main.go
package main

import (
"flag"
"log"
"os"

"nfttokenizerstudio/internal/nfttokenizerstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nfttokenizerstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
