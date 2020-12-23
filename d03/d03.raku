# : can be used to call methods, and can switch order of method and invocant
my @input = map lines: *.comb;

sub count-trees($dr, $dc) {
    # -> pointy blocks are lambdas
    return +grep -> ($r, $c) {
        @input[$r][$c % @input[$r]] eq '#'
    # ... is sequence operator, here infinite sequences
    # Z is infix zip
    }, ((0, * + 1 ...^ * >= @input) Z (0, * + 3 ... *));
}

say count-trees 1, 3;

# reduction metaoperator [*]
# TODO nicer way to splat tuples?
say [*] [(1, 1), (1, 3), (1, 5), (1, 7), (2, 1)].map({ count-trees $_[0], $_[1] });
