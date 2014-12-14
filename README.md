seqLoGo
=======

seqLoGo is able to extract these information of sequence logos shown below, from quite a few of sequences.

* Base counts
* GC contents (-gc option)
* Compressed sequences (-str num option)

About sequence logo in the field of biology, please refer http://weblogo.berkeley.edu.

This package is useful when you would like to make sequence logos during genome-wide analysis.
Previous sequence logo software products cannot deal with too many sequences.
And seqLoGo provides reduced sequences which produce the same logo in the accuracy of (-str num).

Please type
```
	% go get github.com/carushi/seqLoGo
```
and you can get a seqLoGo binary file in $GOPATH/bin.

### Future Plan

* Write test code.
* Produce an image of sequence logo directly.

