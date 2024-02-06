## Create project with go

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
