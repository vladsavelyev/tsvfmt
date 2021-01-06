# tsvfmt

Tab-separated file viewer for command line. Nicely aligns colums in the output.


## Installation

The easiest way is to download the latest binary from the [releases](https://github.com/vladsaveliev/tsvfmt/releases) and make sure to chmod +x the resulting binary.

If you are using go, you can build the binary from source with:

```
go build tsv.go

```

tsvfmt is also available on conda:

```
conda install -c vladsaveliev tsvfmt
``` 

## Usage

```
tsv file.tsv
tsv file.tsv.gz      # comressed files supported
tsv file.csv         # csv files recognised
cat file.tsv | tsv   # piping works
```

## Examples

```
$ cat test/test.tsv
#sample	n	ann	variants	j
NA12878	1		395	12
NA12878	nan	some longer line	4021
NA12878	46	some longer line	11344	497

$ tsv test/test.tsv
#sample    n  ann               variants    j
NA12878    1                         395   12
NA12878  nan  some longer line      4021
NA12878   46  some longer line     11344  497
```

## Advanced options

Custom delimiter

```
tsv -d '|' test/pipe_del_file.txt
```

Scan up to 1000 lines to determine column widths. Default is 100

```
tsv test/test.csv -l 1000
```

Set the maximum width of a column to 20 characters. If a column is wider, it will be cropped

```
tsv -max 10 test/test.vcf.gz

##fileformat=VCFv4.2
#CHROM      POS  ID  REF  ALT         QUAL    FILTER  INFO
chr1     946247  .   G    A,ACGCCGG$  101.77  .       DP=21;AF=$
chr1    1031232  .   T    C           98.17   .       DP=21;AF=$
```

