tsv test/test.tsv | grep -q "NA12891    -0.623         41  some random lines"
cat test/test.tsv | tsv | grep -q "NA12891    -0.623         41  some random lines"
tsv test/test.tsv.gz | grep -q "NA12891    -0.623         41  some random lines"
tsv test/test.csv | grep -q "NA12891    -0.623         41  some random lines"
