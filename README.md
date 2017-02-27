## Installation
```bash
go get -u github.com/pagedegeek/textsbox
```

## Usage

```go
package main

import (
        "github.com/pagedegeek/textsbox"
)

func main() {
        tb := textsbox.New()

        tb.LoadFile("./config/locales/en.yml")
        tb.LoadFile("./config/locales/fr.yml")

        tb.AddKeyAlias("en", "en-En")
        tb.AddKeyAlias("fr", "fr-Fr", "fr-FR")

        // ...
        // locale := session.Get("locale") // can be: en, en-En, fr, fr-Fr, fr-FR

        txt, _ := tb.Find(locale, "messages.welcome")
        welcomeMsg := txt.(string)
        view.render(welcomeMsg)
}
```

```yaml
en:
  messages:
    welcome: Hello World !
```


