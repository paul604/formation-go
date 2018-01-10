## TP01 - Premiers pas

1. Installation des Go Tools :
  - Télécharger les Go Tools https://golang.org/dl/. Sur linux, préférer la version `tar.gz` plutôt que le package manager de votre distribution. Cela permet de récupérer la dernière version de Go et tous les outils du SDK.
  - Utiliser le `GOPATH` par défaut (`$HOME/go`) ou configurer la variable d'environnement `GOPATH` avec le dossier de votre choix. Par exemple sous linux ou Mac OS : `export GOPATH=$HOME/go-workspace`
  - Ajouter le répertoire `$GOPATH/bin/` dans le `PATH`
      Par exemple, sour linux ou Mac OS : `export PATH=$PATH:$GOPATH/bin/`
  - Vérifier l'installation :
      `go version`
  - Ajouter les tps dans le répertoire `$GOPATH/src/` de manière à obtenir l'arborescence suivante :
      ```sh
      $GOPATH
      └── src
          └── formation-go
              ├── TP01
              ├── TP02
              ├── TP03
              ...
      ```
2. Dans le répertoire `TP01/`, créer le ficher `HelloWorld.go` et l'exécuter avec la commande `go run` :
    ```go
    package main

    import (
        "fmt"
    )

    func main() {
        fmt.Println("Hello, world!")
    }
    ```
3. Télécharger l'outil `golint` à l'aide de la commande `go get github.com/golang/lint/golint`
4. Utiliser `golint` pour analyser `HelloWorld.go`. Si `golint` ne rencontre pas d'anomalies, il n'affiche rien. Pour générer un warning on peut, par exemple, déclarer une variable avec un nom qui contient un `_` (*underscore*).
