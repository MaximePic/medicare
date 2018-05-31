# Blockchain - Médicare

	Authors : Maxime Pomier, Maxime Pic, Léa Deschamps, Camille Orgambide
	Version : 1.0

Ce document spécifie le projet "blockchain médicament", appelé "Médicare". 

## Problème


> 800 000 personnes décédées dus à la consommation de médicaments falsifiés dans le monde chaque année

> 15% de médicaments falsifiés en circulation dans le monde

> 300% l’augmentation du trafic de médicaments falsifiés entre 2007 et 2010

> Le marché de la contrefaçon de médicaments est estimé à 200 milliards de dollars par an


## Pourquoi une blockchain ?
> Vérification de l'authenticité du médicament 

> Traçabilité entre différents acteurs

> Procédure rapide

> Registre immuable et durable dans le temps

> Suivi de la non-contrefaçon

## User stories 

### Pour le client 

> Permet de vérifier que la dernière position du produit est bien le lieu où il l’a acheté

### Pour l'état 

> Permet de tracer toutes les contrefaçons existantes dans son pays (au niveau de l’import / export)

## Technologies utilisées 

Architecture basée sur le projet : [https://github.com/hyperledger/education/tree/master/LFS171x/fabric-material/tuna-app](https://github.com/hyperledger/education/tree/master/LFS171x/fabric-material/tuna-app)

## Modélisation d'un médicament 

| Nom du champ | Description  |  Format  |
|--|--|--|
| refProduit | Référence unique à chaque produit (présent dans un datamatrix) | Chaîne de caractères | 
| nomProduit | Nom du produit | Chaîne de caractères |
| nomFabricant | Nom du fabriquant du produit | Chaîne de caractères |
| numeroLot | Numéro de lot du produit | Chaîne de caractères |
| dateExp | Date d'expiration | Date |
| localisationVille | Ville actuelle du produit | Chaîne de caractères |
| locationEtablissement | Etablissement de réception du produit | Chaîne de caractères |
| dateCreation | Date de création du médicament | Date |
| dateReception | Date de réception associée au lieu | Date |
| venteClient | Booléen permettant de savoir si le produit a été vendu | Booléen |


## CRUD 

### CREATE 

La partie CREATE du client web permet l’ajout d’un médicament dans la blockchain. Les différents champs sur la page CREATE sont tous les champs présentés ci-dessus (Partie Modélisation d’un médicament).

### READ 

La partie READ du client web permet de visualiser les différents médicaments présents dans la blockchain. Tous les champs de la blockchain sont visibles dans la partie READ.

### UPDATE 

La partie UPDATE permet de mettre à jour la position actuelle du médicament, ainsi que sa date de vente à partir de la référence du produit.

| Nom du champ | Description  |  Format  |
|--|--|--|
| refProduit | Référence unique à chaque produit (présent dans un datamatrix) | Chaîne de caractères | 
| localisationVille | Ville actuelle du produit | Chaîne de caractères |
| locationEtablissement | Etablissement de réception du produit | Chaîne de caractères |
| dateReception | Date de réception associée au lieu | Date |
| venteClient | Booléen permettant de savoir si le produit a été vendu | Booléen |

### DELETE 

Pas de possibilité de supprimer de médicaments. 


