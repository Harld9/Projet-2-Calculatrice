# Projet 2 : Calculatrice simple

Ce projet consiste en la création d'une calculatrice capable d'effectuer les opérations arithmétiques fondamentales, développée en Go. Il se divise en deux étapes : une application en ligne de commande (CLI) et une interface web interactive.

---

## Phase 1 : Version Terminal (CLI)

La première phase se concentre sur la logique de calcul et la robustesse de l'interaction utilisateur.

### Objectif
Mettre en œuvre les concepts de modularité et de gestion d'erreurs : utilisation de fonctions dédiées pour chaque opération, sécurisation des saisies numériques et gestion des cas d'exclusion mathématique (division par zéro).

### Fonctionnalités
* **Opérations disponibles** :
    * Addition (+)
    * Soustraction (-)
    * Multiplication (x)
    * Division (/)
* **Gestion des erreurs** : Le programme empêche les crashs en cas de saisie de texte et bloque les divisions par zéro avec un message d'alerte.
* **Navigation fluide** : Utilisation de labels et de boucles pour naviguer entre les opérations et le menu principal.
* **Interface claire** : Nettoyage automatique du terminal pour une meilleure lisibilité.

### Installation et Lancement
1. Accéder au répertoire :

    cd "Phase 1 - Calculatrice CLI"

2. Exécuter le programme :

    go run main.go

### Utilisation
1. Sélectionner l'opération souhaitée dans le menu principal (1 à 4).
2. Saisir le premier nombre, puis le deuxième.
3. Lire le résultat affiché.
4. Choisir de continuer ou de revenir au menu (0).
5. Utiliser l'option 5 pour quitter.

---

## Phase 2 : Extension Web (HTML + Go)

Cette phase transpose la logique de calcul sur le web en proposant une interface utilisateur moderne et ergonomique.

### Objectif
Gérer des formulaires multiples sur une même interface et assurer la transmission des données via des requêtes POST, tout en maintenant une structure de code organisée et évolutive.

### Fonctionnalités
* **Interface Multi-blocs** : Des sections distinctes et stylisées pour chaque type de calcul.
* **Validation côté Serveur** : Vérification en temps réel de la validité des nombres saisis.
* **Retour d'état** : Affichage différencié des résultats (succès en vert) et des erreurs de saisie (alertes en rouge).
* **Architecture modulaire** : Séparation stricte entre les routes, les contrôleurs et la logique de calcul.

### Installation et Lancement
1. Accéder au répertoire :
```bash
    cd "Phase 2 - Calculatrice Web"
```
2. Démarrer le serveur :
```bash
    go run main.go
```
3. Accéder à l'application :
Ouvrir un navigateur à l'adresse : http://localhost:8080

### Structure du Projet Web
* **main.go** : Point d'entrée et démarrage du serveur HTTP.
* **controller/** : Gestion des données de formulaire et rendu des templates.
* **logic/** : Noyau de calcul (fonctions arithmétiques).
* **router/** : Configuration des points d'accès (endpoints) de l'application.
* **template/** : Fichiers HTML avec injection de données dynamiques.
* **static/** : Design de l'interface via CSS.

---

## Annexes : Logique de calcul

Le projet implémente les opérations de base selon les principes suivants :
* **Addition** : a + b
* **Soustraction** : a - b
* **Multiplication** : a * b
* **Division** : a / b (avec vérification préalable b différent de 0)