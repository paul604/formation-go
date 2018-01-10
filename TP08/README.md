## TP08 - Composition de struct

A partir du code du TP07 changer le type `book` : le champ `author` est maintenant de type `person` :

```go
    type book struct {
        title  string
        author person
        year   int
    }
```

1. Compléter la définition du type `person` avec les champs `firstName`, `lastName`

2. Adapter le reste de tp08.go afin qu'il compile
