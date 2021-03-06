seqLoGo
=======

seqLoGo is able to extract these information for sequence motifs shown below

* Base counts
* GC contents (-gc option)
* Amino acid or any characters counts (-any option)
* Compressed sequences (-str num option)

from **quite a few of sequences**.

About sequence logo in the field of biology, please refer http://weblogo.berkeley.edu.

This package is useful when you would like to make sequence logos during genome-wide analysis.
Previous sequence logo software products cannot deal with too many sequences.
And seqLoGo provides reduced sequences which produce the same logo in the accuracy of (-str num).

Please type
```
	% go get github.com/carushi/seqLoGo
```
and you can get a seqLoGo binary file in $GOPATH/bin.


Here is an example.
```
	% seqLoGo -input_file temp.txt -gc
	% seqLoGo -input_file temp.fa -any -fasta
	% seqLoGo -any -str 20  (-> Stdin)
```

### Future Plan

* Write test code.
* Produce an image of sequence logo directly.

