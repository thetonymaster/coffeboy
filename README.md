# Coffeboy

##Setup
First you should have a file structure like this: ````$GOPATH/src/github.com/crowdint```` then clone this repo inside the ````crowdint```` folder.

Then get the godeps package ````go get github.com/tools/godep````.

Once downloaded, you should add the the bin folder to your PATH in your .bashrc file or equivalent:
````export PATH=$PATH:$GOPATH/bin````, and then reload or open a new terminal.

Now you will be able to run ````godep restore```` to download all the depndencies.
