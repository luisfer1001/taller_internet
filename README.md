## Create project with go - TALLER INTERNET

```
go mod init github.com/test
go run ./cmd/basic/main.go
go run ./cmd/server/main.go

go get -u github.com/go-chi/chi/v5
go get -u github.com/mattn/go-sqlite3
```

## Git

```
git remote add origin https://github.com/cgalvisleon/test.git
git branch -M main

git remote add origin
git push -u origin --all

git push --set-upstream origin develop
git push --set-upstream origin juan
git push --set-upstream origin joan &&
git push --set-upstream origin fernando &&
git push --set-upstream origin nicolas &&
git push --set-upstream origin bernardo &&
git push --set-upstream origin cesar
```

## Git Config

```
git config --global alias.s "status -s -b"
git config --global alias.get "! git pull origin $1"
git config --global alias.new "! git checkout -b $1"
git config --global alias.set "! git checkout $1"
git config --global alias.prepare "! git add .; git commit -m Prepare"
git config --global alias.initial "! git add .; git commit -m 'Initial commit'; git push --set-upstream origin main"
git config --global alias.backup "! git add .; git commit -m 'Backup'; git push -u origin $1"
git config --global alias.update "! git add .; git commit -m 'Update'; git push -u origin $1"
git config --global alias.review "! git add .; git commit -m 'Review'; git push -u origin $1"
git config --global alias.tags "! git push -u origin --tags"
git config --global alias.rhs 'reset HEAD^ --soft'
git config --global alias.rhh 'reset HEAD^ --hard'
git config --global alias.clean 'git rm -r --cached .'

git remote -v
```

## Models

List models to created

```
Juan: formulario
Atributs:
-id int
-name text
-descripcion text (Listo)

Luis Fernando: imagenes
Atributs:
-id int
-name: text
-file: text
-id_formulario: int

Joan: Faulers
Atributs:
-id int
-code int
-faules text

Nicolas: FauresAreas
Atributs:
-id int
-id_fauler: int
-name_area: text

Luis Bernardo: FauresFeedback
Atributs:
-id int
-id_fauler: int
-feedback text

Cesar: FaulersClients
Atributs:
-id int
-id_fauler: int
-client_name: text
```

## Git

```
git push --set-upstream origin

git add .;
git commit -m 'Review';
git push -u origin $1;

git push -u ofigin --force;
```

## Merge

```
git rhs #Reset commit and preserve changes
git rhh #Reset commit and discard changes

git reset HEAD^ --hard
git reset HEAD^ --soft

git reset <id> --hard
git reset <id> --soft

git branch -v #Listar ramas

git merge <rama>
git merge -X theirs <rama>
git rebase origin/<rama>
git checkout origin/<rama> -- <filename>

git checkout origin/nicolas -- pkg/fauresAreas/controlador.go
```
