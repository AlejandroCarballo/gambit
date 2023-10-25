git add .
git commit -m "Ultimo Commit"
git push
go build main.go
rm -f main.zip
zip main.zip main