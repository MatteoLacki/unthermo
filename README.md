## Basic tools used to get raw thermo files.

This code is by no means mine. It's made by guys building proteininspector.
Here you will find simple wrappers and, above all, installation instructions.
This is intended for people who, like me, never touched golang and want to keep it that way.

# Installation:
0. Get golang (like they describe here: https://golang.org/doc/install)
1. Get this github repo.
2. Navigate to where the *.go files are.
3. To compile of whatever golang does with the source files, run
``bash
find . -type f -name "*.go" -exec go build {} \;
``
4. copy the output to some bin and chmod +x them. Damn, maybe even add them to PATH.

Voila!
Look ma', almost no golang exposure!
