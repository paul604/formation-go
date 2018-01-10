## TP03 - Variadic functions et closures

Le but de cet exercice est de modifier le code de `tp03.go` pour transformer `isPI()` en une fonction variadique et utiliser des closures pour afficher les messages :

1. Vérifier que le programme fonctionne bien : `go run tp03.go`
2. Dé-commenter le code commenté dans `tp03.go` et vérifier qu'il ne compile plus : `go run tp03.go`
3. Modifier la déclaration de la fonction `isPI()` pour supprimer les erreurs de compilation (N.B. un paramètre variadique est de type `slice`. C'est également le type de la variable `digits` au début de `isPI()`. Même si on n'a pas encore vu le type `slice`, vous pourrez itérer sur le paramètre variadique de la même façon qu'on itère sur la variable `digits`)
4. Assigner les deux variables `successMessage` et `failureMessage` en utilisant deux fonctions
