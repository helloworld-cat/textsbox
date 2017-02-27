## About
Simple and easy lib. to manage translations and texts, with cache, in Go project.

## Installation

```bash
go get -u github.com/pagedegeek/textsbox
```

## Usage

```go
// main.go
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
# config/locales/en.yml
en:
  messages:
    welcome: Hello World !
```

```yaml
# config/locales/fr.yml
fr:
  messages:
    welcome: Bonjour tout le monde !
```

## Other features

### Load YAML content

```go
tb.LoadFile("...") // load from file
```

```go
tb.Load(reader) // load from io.Reader
```

### Reset cache
```go
tb.ResetCache()
```

