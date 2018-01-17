## TP12 - Méthodes

1. Compléter la définition du type `worker` qui est une struct avec plusieurs champs :
  - un champ anonyme de type `person`
  - `company` de type `company`
  - `role` de type `work`
  - `monthlyPay` de type `salary`
  - `certifications` de type slice de `certification`
2. Instancier deux `worker` (`worker1` et `worker2`)
3. Créer la méthode `certificationsNum()` qui a comme récepteur un `worker` et qui retourne le nombre de certifications du worker
4. Créer la méthode `isOfficeCompanyHeadquarters()` qui a toujours un `worker` comme récepteur et qui retourne `true` si le bureau du worker se trouve à la même adresse que le headquarter
5. Compléter le `main()` qui itère sur les éléments d'un tableau de `worker` et affiche le nombre de certifications de chaque worker et s'il travaille ou pas au headquarter.
