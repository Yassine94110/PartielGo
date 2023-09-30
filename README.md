
# Introduction

Ce programme est écrit en Go et vise à tester la connectivité à une série de ports sur un serveur spécifique. Pour chaque port ouvert, il effectue une série de requêtes HTTP pour interagir avec des endpoints spécifiques.

# Fonctionnalités

Test de connectivité: Le programme tente de se connecter à chaque port dans la plage spécifiée pour déterminer s'il est ouvert.
Requêtes HTTP: Pour chaque port ouvert, le programme effectue une série de requêtes HTTP, notamment des requêtes GET pour /ping et des requêtes POST pour /signup, /check, /getUserLevel, /getUserPoints, /getUserSecret, /enterChallenge, et /submitSolution.
Extraction du niveau: Le programme extrait le niveau d'un utilisateur à partir de la réponse de l'endpoint /getUserLevel.
Comment exécuter
Pour exécuter le programme, utilisez la commande suivante:

bash
Copy code
go run .

# Configuration
Adresse IP du serveur: L'adresse IP du serveur est actuellement définie sur 10.49.122.144. Si nécessaire, vous pouvez la modifier dans la fonction main.
Plage de ports: Le programme teste tous les ports de 1000 à 65535. Vous pouvez ajuster cette plage en modifiant les variables minPort et maxPort dans la fonction main.
Remarques
Assurez-vous d'avoir Go installé sur votre machine pour exécuter ce programme.
Le programme utilise des goroutines pour tester simultanément plusieurs ports, ce qui le rend plus rapide mais peut également consommer plus de ressources système.
Certaines parties du code, comme la requête pour /iNeedAHint, sont actuellement commentées. Vous pouvez les décommenter si nécessaire.
