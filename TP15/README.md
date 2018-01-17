## TP15 - goroutines

Dans cette exercice on utilisera la méthode `curl()` qui récupère une page web et affiche le nom, les bytes et le temps écoulé :
```
zenika 37623 [0.36s]
```

`tp15.go` utilise `curl()` de façon séquentielle pour récupérer une liste d'urls. Le temps qu'il emploie pour les récupérer correspond donc à la somme de chaque `curl()` :
```
zenika 37623 [0.36s]
facebook 72097 [1.95s]
apple 29962 [2.72s]
microsoft 73704 [3.02s]
amazon 223652 [5.32s]
Time to fetch all urls: 6s
```

On veut améliorer `tp15.go` pour récupérer les url en parallèle.

1. Executer `tp15.go` : `go run tp15.go` et noter le temps que le programme emploie pour récupérer les 5 urls
2. Utiliser les goroutines pour que l'execution de chaque fonction `curl()` se fasse dans une goroutine distincte
3. Mettre un timeout dans `main()` pour attendre la fin de toutes les goroutines (on peut utiliser la fonction `time.Sleep()`) et exécuter
