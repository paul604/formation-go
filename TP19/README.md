## TP19 - Couverture des tests

1. Analyser la couverture des tests avec `go test -cover formation-go/TP19/anagram`
2. Générer un rapport de profiling avec les deux commandes :
```bash
go test -coverprofile=c.out formation-go/TP19/anagram
go tool cover -html=c.out -o coverage.html
```
3. Ajouter un test pour faire passer la couverture de code à 100%
