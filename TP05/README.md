## TP05 - defer, panic et recover

Le code du tp05 compte de 0 à 10 via des appels récursif. Le but du TP est
de modifier le code en utilisant `defer`, `panic`et `recover`
pour avoir l'output suivant :

```
0
1
2
3
4
5
6
7
8
9
10
10
9
8
7
6
5
4
3
2
1
0
done!
```

1. Executer `tp05.go` avec la commande `go run tp05.go`
2. Modifier la fonction `count` en utilisant `defer` pour pouvoir compter à rebours.
3. Modifier la fonction `count` pour que lorsque `from` == `to`, on renvoie une panique avec la valeur "Done!"
4. Modifier la fonction `main` pour que lorsque la panique se produit, on affiche quand même le message de la panique
