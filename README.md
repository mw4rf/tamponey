# Tamponey

Ajoute des tampons de façon séquentielle sur la première page de tous les fichiers PDF présents dans le dossier donné en argument ou, à défaut, dans le dossier courant.

La séquence est définie par le nom de chaque fichier, qui doit commencer par un nombre suivi d'un espace.

Les fichiers originaux sont préservés et des copies tamponnées sont créées.

Le texte du tampon et sa configuration (position, dimensions, couleurs, etc.) sont définis dans un fichier YAML placé dans le dossier de configuration par défaut:

- Linux: `~/.config/tamponey/config.yaml`
- MacOS: `~/Library/Application Support/tamponey/config.yaml`
- Windows: `%APPDATA%\tamponey\config.yaml`

Ce fichier de configuration définit `stamp.template` et `stamp.configuration`.

`stamp.template` est un modèle de texte avec un emplacement pour le numéro à ajouter identifié par `%s`. Des sauts de ligne peuvent être ajoutés avec `\n`.

`stamp.configuration` répond à la syntaxe de pdfcpu, disponible à l'adresse suivante: https://pdfcpu.io/core/watermark

---

Adds sequential stamps to the first page of all PDF files in the folder provided as an argument or, if none is provided, in the current folder.

The sequence is defined by the name of each file, which must start with a number followed by a space.

The original files are preserved, and stamped copies are created.

The text of the stamp and its configuration (position, dimensions, colors, etc.) are defined in a YAML file located in the default configuration folder:

    Linux: `~/.config/tamponey/config.yaml`
    MacOS: `~/Library/Application Support/tamponey/config.yaml`
    Windows: `%APPDATA%\tamponey\config.yaml`

This configuration file defines stamp.template and stamp.configuration.

`stamp.template` is a text template with a placeholder for the number to add, identified by `%s`. Line breaks can be added with `\n`.

`stamp.configuration` follows the syntax of pdfcpu, available at the following URL: https://pdfcpu.io/core/watermark

---
``` yaml
# Exemple
stamp:
    template: "Cabinet Trucmuche\nAvocats\Pièce n°%s\n"
    configuration: "pos: tr, aligntext: c, scalef: 0.5 abs, ma:5, bo:3 round #b30000, c:#b30000, bgc:#ffffff, rot:0"
```

