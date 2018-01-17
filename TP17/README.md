## TP17 - select et timeout

`tp17.go` visite récursivement le contenu du dossier passé en paramètre (la fonction `filepath.Walk()` de la librairie standard est utilisée).

```bash
    $ go run tp17.go .
    filepath.Walk() visited 5 elements
```

La visite des répertoires est faite dans une `goroutine` : `go doWalk(root, c)`
Une autre `goroutine` est créée pour le spinner : `go spinner(100 * time.Millisecond)`

Le but de cet exercice est de mettre en place un timeout qui termine le programme si la fonction `doWalk()` ne rend pas la main au bout de 5 secondes. Pour mettre en place le timeout, il faudra utiliser l'instruction `select`.
