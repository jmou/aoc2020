# : can be used to call methods, and can switch order of method and invocant
my @input = map lines: *.comb;

sub count-trees($dr, $dc) {
    # -> pointy blocks are lambdas
    # () around parameters destructures a list (from zipped lists)
    return +grep -> ($r, $c) {
        @input[$r][$c % @input[$r]] eq '#'
    # ... is sequence operator; * RHS is infinite, else conditional bound
    # Z is infix zip
    }, ((0, * + $dr ...^ * >= @input) Z (0, * + $dc ... *));
}

say count-trees 1, 3;

# reduction metaoperator [*]
# | slip operator in call makes parameters "slippy" (like splatting)
# , construct the list; () are just used for precedence
# [] circumfix coerces to Array (intention here was just to disambiguate parentheses)
say [*] [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)].map({ count-trees |$_ });
