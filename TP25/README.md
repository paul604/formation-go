## Web service et JSON

Dans ce TP nous allons développer une API REST qui renvoie la météo d'une ville. Les données météo sont récupérées du service Yahoo Weather. Le TP se déroule en 3 phases :

1. Compléter le client Yahoo pour décoder les réponses json :
    - Lancer les test de l'API yahoo et vérifier qu'ils ne passent pas : `go test formation-go/TP25/yahoo`
    - Compléter les deux fonctions `unmarshalWeatherResult` et `unmarshalWoeidResult` du fichier `yahoo_api.go` qui font le mapping entre le format json (`data`) et les struct de type `weatherQueryResult` et `woeidQueryResult` (`result`)
    - Une fois les fonctions complétées vérifier que les tests passent
2. Ajouter la route pour servir l'API REST /cities en utilisant le package `http` de la standared library`
    - Lancer le serveur `go run tp25.go`. Vérifier que le serveur répond bien à l'adresse http://localhost:8000/
    - Dans le `main` de `tp25.go` ajouter un deuxième appel à `http.HandleFunc()`. Le path à servir sera `/cities` et `citiesHandler()` sera la fonction qui servira ce deuxième path.
    - Relancer le serveur et vérifier que l'adresse http://localhost:8000/cities répond une liste d'objets de type `Weather` en format json.
3. Utiliser un package de routage plus avancé (github.com/gorilla/mux) pour compléter l'API.
    - Dans le `main` de `tp25.go` commenter les premières lignes, celles ou le routage est fait avec le package `http` de la standard library. Et décommenter les lignes ou le routage est faite avec le package `mux` qui offre plus de options par rapport au premier.
    - Relancer le serveur et vérifier qu'il répond correctement aux URL http://localhost:8000/ et http://localhost:8000/cities
    - Ajouter une option de routage pour servir les path du type `/cities/{name}` avec la fonction `cityHandler`
    - Relancer le serveur et vérifier que l'application répond bien aux adresses :
        - http://localhost:8000/
        - http://localhost:8000/cities
        - http://localhost:8000/cities/Paris
        - http://localhost:8000/cities/Rome
        - http://localhost:8000/cities/Cannes
