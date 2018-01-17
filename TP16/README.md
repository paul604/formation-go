## TP16 - channels

En repartant du TP précédent, nous allons utiliser les *channels* pour profiter au maximum de la concurrence des `goroutines`. `curl()` pourra communiquer avec `main()` pour lui dire qu'il a terminé son execution à l'aide d'un channel.

1. Instancier un `chan` de type `string` au début du `main()`
2. Passer cette variable `chan` comme paramètre à la fonction `curl()`. Celle-ci devra être modifiée afin qu'elle écrive les messages sur le channel et non plus sur `stdout`.
3. À la place de l'instruction `time.Sleep()` dans `main()`, il faudra écrire une boucle `for` qui, pour chaque `goroutine` affichera le messages reçus via le channel (`fmt.Print(<-c)`).
4. (Facultatif) Remplacer l'utilisation du `chan` par un `sync.WaitGroup`.
