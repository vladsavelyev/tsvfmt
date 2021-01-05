go build tsv.go
chmod a+x tsv
mkdir -p $PREFIX/bin
cp tsv $PREFIX/bin/tsv
