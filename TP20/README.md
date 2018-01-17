## TP20 - Benchmark

`tp20.go` contient l'implementation de trois algorithmes de tri : `bubblesort()`, `selectionsort()` et `quicksort()`. L'objectif de ce TP est de comparer les performances de ces trois algorithmes en utilisant l'outil de benchmark de `go`.

1. Créer un fichier de tests `tp20_test.go` avec trois fonctions de test `BenchmarkBubbleSort()`, `BenchmarkSelectionsort()` et `BenchmarkQuickSort()`.
2. Lancer l'outil de benchmark : `go test -bench=. -benchmem`
