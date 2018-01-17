## TP24 - Vendoring et dep

Dans ce TP on développera une application qui utilise une librairie externe `github.com/briandowns/openweathermap` pour utiliser les API de `openweathermap` (https://home.openweathermap.org/). Pour la gestion de cette dépendance on utilisera la technique de vendoring.

1. Essayer d'exécuter le programme : `go run tp24.go`
2. Récupérer la dépendance : `go get github.com/briandowns/openweathermap`
    - Inspecter les répertoires src et pkg et vérifier que le package a bien été téléchargé
3. Re-essayer d'éxécuter le programme. Cette fois-ci vous devrez avoir un output de ce type :
```
Paris: ensoleillé, 4.1C
Cannes: ensoleillé, 3.3C
Rome: ensoleillé, 19.3C
```
4. Utiliser `dep` (https://github.com/golang/dep) pour créer le répertoire vendoring :
```
go get -u github.com/golang/dep/cmd/dep
dep init
```
    - Vérifier que vous avez bien un fichier `Gopkg.toml`
5. Effacer le répertoire `src/github.com/briandowns/openweathermap` et relancer le programme.
6. Pour faire un update de la librairie `openweathermap` il suffira de lancer :
```
dep ensure -update
```
