# rsa-private-key-generator

A cli tool that helps you generate a private RSA key. Choose the amount of bits with the first argument so to generate with 8192 bits, run:

    $ rsa-private-key-generator 8192

Remember to always keep your private keys private and hidden, if saving the key to a file, make sure that the file has the permissions 0600.

## Try it out

Make sure that you have go (https://golang.org/doc/install) installed and the go bin/ directory on your PATH variable, then: 

Install with

    $ go get github.com/mojlighetsministeriet/rsa-private-key-generator

Run with

    $ rsa-private-key-generator 8192

## License

All the code in this project goes under the license GPLv3 (https://www.gnu.org/licenses/gpl-3.0.en.html).
