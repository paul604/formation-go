## TP02 - Packages

Le programme `tp02.go` affiche le nombre de jours qui se sont écoulés à partir du lancement de Go (10 novembre 2009). Le but de cet exercice est de déplacer une partie de ce programme dans un nouveau package `golaunch`.

1. Lancer `tp02.go` avec la commande `go run tp02.go` et vérifier que le message s'affiche bien.
2. Créer un dossier `golaunch` et :
     - créer un nouveau fichier `date.go` dans le dossier `golaunch/`
3. Déplacer la fonction `getDaysSinceGoLaunch()` dans le fichier `date.go` :
    - Le package de `date.go` doit être `golaunch`
    - Importer le package `golauch` dans le fichier `tp02.go` : `import "formation-go/TP02/golaunch"`
    - Le nom de la fonction `getDaysSinceGoLaunch()` doit commencer par une majuscule pour qu'elle soit *visible* de l'extérieur
4. Lancer `tp02.go` avec la commande `go run tp02.go` et vérifier que le message s'affiche bien.
