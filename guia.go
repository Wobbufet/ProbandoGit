package guia

/*
pwd         - para saber donde se esta
cd "        - para navegar (..)"arriba" (name)"escritorio al cual se va"
ls          - para ver lo que contiene la carpeta
git init    - inicia el repositorio en la carpeta seleccionada
            crea 2 carpetas invisibles en dicho repositorio
git status  - para ver los archivos que se tienen
            rojo significa que esta en el aire, no se agrego al entorno de trabajo
git add ""  - Agregas el archivo a tu entorno de trabajo. (git add guia.go)
            aparecerá en azul pues ya se incluyó

git commit  - crea un punto de control del código, pues puede tener cambios
        así que primero se crea unarchivo para configurar

git c --global user.email caceresgabriel2001@gmail.com     -si se escribe mal solo se vuelve a poner, ya que se sobreescribe
git config --global user.name Gabriel

git commit  - abre el editor de código, si no deja editar hay que presionar la tecla 'I'
            cuando se termine se presiona "Esc" y se pone ":wq" (write-quit)
            al acabar debe haber creado IDs para diferenciar las versiones que se guardan

git log     - para ver los commits que se han creado (el de arriba es el mas reciente)
            si ejecuto alguna intrucción dirá que no hay cambios y no hay nada que hacer,
            pero si hacemos cambios en el código, al poner 'status' veremos que estará en rojo

git checkout -- guia.go     - esto retornará a su estado anterior el archivo, al punto en el que se guardó con "add"

git diff guia.go        - para ver las diferencias hechas en los archivos
                    rojo: lo que se quitó       verde: lo que se agregó
                    para salir de ese menú se preiona la tecla "Q" (quit)
*/
