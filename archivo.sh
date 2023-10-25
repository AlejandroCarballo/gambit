git add .
git commit -m "Ultimo Commit"
git push
go build main
rm main.zip
zip "$main.zip" "$main"