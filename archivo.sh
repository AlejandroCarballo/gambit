git add .
git commit -m "Ultimo Commit"
git push
# shellcheck disable=SC2121
set GOOS=linux
# shellcheck disable=SC2121
set GOARCH=amd64
go build main
rm main.zip
zip main.zip main
