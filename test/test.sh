tsv test/test.tsv | grep -q "NA12878   46  some longer line     11344  497"
cat test/test.tsv | tsv | grep -q "NA12878   46  some longer line     11344  497"
tsv -max 5 test/test.vcf.gz | grep -q 'chr1   9462$  .   G    A,AC\$  101.\$  .      DP=2\$'
tsv test/test.csv | grep -q "NA12878    -0.865         36  some"
tsv -d '|' test/pipe_del_file.txt | grep -q "NA12891 A  -0.623   41  some annotation"